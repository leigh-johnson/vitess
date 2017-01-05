#!/usr/bin/env python
"""Google Compute Engine sandbox components."""

import logging
import os
import subprocess

import sandlet


class Cluster(object):
  """Google Container Engine Cluster."""

  _DEFAULT_ZONE = 'us-central1-b'
  _DEFAULT_MACHINE_TYPE = 'n1-standard-4'
  _DEFAULT_NODE_COUNT = 5

  def __init__(self, params):
    self.params = params

  def start(self):
    zone = self.params.get('gke_zone', self._DEFAULT_ZONE)
    machine_type = self.params.get('machine_type', self._DEFAULT_MACHINE_TYPE)
    node_count = str(self.params.get('node_count', self._DEFAULT_NODE_COUNT))
    subprocess.call(['gcloud', 'config', 'set', 'compute/zone', zone])
    subprocess.call(
        ['gcloud', 'container', 'clusters', 'create', self.params['name'],
         '--machine-type', machine_type, '--num-nodes', node_count, '--scopes',
         'storage-rw'])

  def stop(self):
    zone = self.params.get('gke_zone', self._DEFAULT_ZONE)
    subprocess.call(['gcloud', 'container', 'clusters', 'delete',
                     self.params['name'], '-z', zone, '-q'])


class Port(sandlet.Sandlet):
  """Used for forwarding ports in Google Container Engine."""

  def __init__(self, name, port):
    self.name = name
    self.port = port
    super(Port, self).__init__(name)

  def start(self):
    # Check for existence first.
    with open(os.devnull, 'w') as dn:
      # Suppress output for the existence check to prevent unnecessary output.
      firewall_rules = subprocess.check_output(
          ['gcloud', 'compute', 'firewall-rules', 'list', self.name],
          stderr=dn)
      if self.name in firewall_rules:
        logging.info('Firewall rule %s already exists, skipping creation.',
                     self.name)
        return
    subprocess.call(['gcloud', 'compute', 'firewall-rules', 'create',
                     self.name, '--allow', 'tcp:%s' % str(self.port)])

  def stop(self):
    try:
      subprocess.check_call(
          ['gcloud', 'compute', 'firewall-rules', 'delete', self.name, '-q'])
    except subprocess.CalledProcessError:
      logging.warn('Failed to delete firewall rule %s', self.name)
