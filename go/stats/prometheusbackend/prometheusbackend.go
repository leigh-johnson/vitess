package prometheusbackend

import (
	"expvar"
	"net/http"
	"strings"
	"unicode"

	log "github.com/golang/glog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"vitess.io/vitess/go/stats"
)

// PromBackend implements PullBackend using Prometheus as the backing metrics storage.
type PromBackend struct {
	namespace string
}

var (
	be *PromBackend
)

// Init initializes the Prometheus be with the given namespace.
func Init(namespace string) {
	http.Handle("/metrics", promhttp.Handler())
	be := &PromBackend{namespace: namespace}
	stats.Register(be.publishPrometheusMetric)
}

// ValueType specifies whether the value of a metric goes up monotonically
// or if it goes up and down. This is useful for exporting to backends that
// differentiate between counters and gauges.
type ValueType int

const (
	// CounterValue is used to specify a value that only goes up (but can be reset to 0).
	CounterValue = iota
	// GaugeValue is used to specify a value that goes both up and down.
	GaugeValue
)

// PublishPromMetric is used to publish the metric to Prometheus.
func (be *PromBackend) publishPrometheusMetric(name string, v expvar.Var) {
	switch st := v.(type) {
	case *stats.Counter:
		be.newMetric(st, name, CounterValue)
	case *stats.Gauge:
		be.newMetric(&st.Counter, name, GaugeValue)
	case *stats.CounterFunc:
		be.newMetricFunc(st, name, CounterValue)
	case *stats.GaugeFunc:
		be.newMetricFunc(&st.CounterFunc, name, GaugeValue)
	case *stats.CountersWithLabels:
		be.newMetricWithLabels(&st.Counters, name, st.LabelName(), CounterValue)
	case *stats.CountersWithMultiLabels:
		be.newCountersWithMultiLabels(st, name)
	case *stats.CountersFuncWithMultiLabels:
		be.newMetricsFuncWithMultiLabels(st, name, CounterValue)
	case *stats.GaugesFuncWithMultiLabels:
		be.newMetricsFuncWithMultiLabels(&st.CountersFuncWithMultiLabels, name, GaugeValue)
	case *stats.GaugesWithLabels:
		be.newMetricWithLabels(&st.Counters, name, st.LabelName(), GaugeValue)
	case *stats.GaugesWithMultiLabels:
		be.newGaugesWithMultiLabels(st, name)
	case *stats.Timings:
		be.newTiming(st, name)
	case *stats.MultiTimings:
		be.newMultiTiming(st, name)
	default:
		log.Warningf("Unsupported type for %s: %T", name, st)
	}
}

func (be *PromBackend) newMetricWithLabels(c *stats.Counters, name string, labelName string, vt ValueType) {
	collector := &metricWithLabelsCollector{
		counter: c,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			c.Help(),
			[]string{labelName},
			nil),
		vt: vt}

	prometheus.MustRegister(collector)
}

func (be *PromBackend) newCountersWithMultiLabels(cml *stats.CountersWithMultiLabels, name string) {
	c := &metricWithMultiLabelsCollector{
		cml: cml,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			cml.Counters.Help(),
			labelsToSnake(cml.Labels()),
			nil),
	}

	prometheus.MustRegister(c)
}

func (be *PromBackend) newGaugesWithMultiLabels(gml *stats.GaugesWithMultiLabels, name string) {
	c := &multiGaugesCollector{
		gml: gml,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			gml.Help(),
			labelsToSnake(gml.Labels()),
			nil),
	}

	prometheus.MustRegister(c)
}

func (be *PromBackend) newMetricsFuncWithMultiLabels(cfml *stats.CountersFuncWithMultiLabels, name string, vt ValueType) {
	collector := &metricsFuncWithMultiLabelsCollector{
		cfml: cfml,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			cfml.Help(),
			labelsToSnake(cfml.Labels()),
			nil),
		vt: vt,
	}

	prometheus.MustRegister(collector)
}

func (be *PromBackend) newTiming(t *stats.Timings, name string) {
	collector := &timingsCollector{
		t: t,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			t.Help(),
			[]string{"Histograms"}, // hard coded label key
			nil),
	}

	prometheus.MustRegister(collector)
}

func (be *PromBackend) newMultiTiming(mt *stats.MultiTimings, name string) {
	collector := &multiTimingsCollector{
		mt: mt,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			mt.Help(),
			labelsToSnake(mt.Labels()),
			nil),
	}

	prometheus.MustRegister(collector)
}

func (be *PromBackend) newMetric(c *stats.Counter, name string, vt ValueType) {
	collector := &metricCollector{
		counter: c,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			c.Help(),
			nil,
			nil),
		vt: vt}

	prometheus.MustRegister(collector)
}

func (be *PromBackend) newMetricFunc(cf *stats.CounterFunc, name string, vt ValueType) {
	collector := &metricFuncCollector{
		cf: cf,
		desc: prometheus.NewDesc(
			be.buildPromName(name),
			cf.Help(),
			nil,
			nil),
		vt: vt}

	prometheus.MustRegister(collector)
}

func (be *PromBackend) buildPromName(name string) string {
	s := strings.TrimPrefix(toSnake(name), be.namespace)
	return prometheus.BuildFQName("", be.namespace, s)
}

func labelsToSnake(labels []string) []string {
	output := make([]string, len(labels))
	for i, l := range labels {
		output[i] = toSnake(l)
	}
	return output
}

func toSnake(name string) string {
	// Special cases
	r := strings.NewReplacer("VSchema", "Vschema", "VtGate", "Vtgate")
	name = r.Replace(name)
	// Camel case => snake case
	runes := []rune(name)

	var out []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < len(runes) && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}