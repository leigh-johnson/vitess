// Code generated by MockGen. DO NOT EDIT.
// Source: resharding_wrangler.go

// Package resharding is a generated GoMock package.
package resharding

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	topodata "vitess.io/vitess/go/vt/proto/topodata"
)

// MockReshardingWrangler is a mock of ReshardingWrangler interface
type MockReshardingWrangler struct {
	ctrl     *gomock.Controller
	recorder *MockReshardingWranglerMockRecorder
}

// MockReshardingWranglerMockRecorder is the mock recorder for MockReshardingWrangler
type MockReshardingWranglerMockRecorder struct {
	mock *MockReshardingWrangler
}

// NewMockReshardingWrangler creates a new mock instance
func NewMockReshardingWrangler(ctrl *gomock.Controller) *MockReshardingWrangler {
	mock := &MockReshardingWrangler{ctrl: ctrl}
	mock.recorder = &MockReshardingWranglerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReshardingWrangler) EXPECT() *MockReshardingWranglerMockRecorder {
	return m.recorder
}

// CopySchemaShardFromShard mocks base method
func (m *MockReshardingWrangler) CopySchemaShardFromShard(ctx context.Context, tables, excludeTables []string, includeViews bool, sourceKeyspace, sourceShard, destKeyspace, destShard string, waitSlaveTimeout time.Duration) error {
	ret := m.ctrl.Call(m, "CopySchemaShardFromShard", ctx, tables, excludeTables, includeViews, sourceKeyspace, sourceShard, destKeyspace, destShard, waitSlaveTimeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopySchemaShardFromShard indicates an expected call of CopySchemaShardFromShard
func (mr *MockReshardingWranglerMockRecorder) CopySchemaShardFromShard(ctx, tables, excludeTables, includeViews, sourceKeyspace, sourceShard, destKeyspace, destShard, waitSlaveTimeout interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopySchemaShardFromShard", reflect.TypeOf((*MockReshardingWrangler)(nil).CopySchemaShardFromShard), ctx, tables, excludeTables, includeViews, sourceKeyspace, sourceShard, destKeyspace, destShard, waitSlaveTimeout)
}

// WaitForFilteredReplication mocks base method
func (m *MockReshardingWrangler) WaitForFilteredReplication(ctx context.Context, keyspace, shard string, maxDelay time.Duration) error {
	ret := m.ctrl.Call(m, "WaitForFilteredReplication", ctx, keyspace, shard, maxDelay)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForFilteredReplication indicates an expected call of WaitForFilteredReplication
func (mr *MockReshardingWranglerMockRecorder) WaitForFilteredReplication(ctx, keyspace, shard, maxDelay interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForFilteredReplication", reflect.TypeOf((*MockReshardingWrangler)(nil).WaitForFilteredReplication), ctx, keyspace, shard, maxDelay)
}

// MigrateServedTypes mocks base method
func (m *MockReshardingWrangler) MigrateServedTypes(ctx context.Context, keyspace, shard string, servedType topodata.TabletType, reverse, skipReFreshState bool, filteredReplicationWaitTime time.Duration, reverseReplication bool) error {
	ret := m.ctrl.Call(m, "MigrateServedTypes", ctx, keyspace, shard, servedType, reverse, skipReFreshState, filteredReplicationWaitTime, reverseReplication)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateServedTypes indicates an expected call of MigrateServedTypes
func (mr *MockReshardingWranglerMockRecorder) MigrateServedTypes(ctx, keyspace, shard, servedType, reverse, skipReFreshState, filteredReplicationWaitTime, reverseReplication interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateServedTypes", reflect.TypeOf((*MockReshardingWrangler)(nil).MigrateServedTypes), ctx, keyspace, shard, servedType, reverse, skipReFreshState, filteredReplicationWaitTime, reverseReplication)
}
