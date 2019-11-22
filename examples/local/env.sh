#!/bin/bash

# Copyright 2019 The Vitess Authors.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

hostname=`hostname -f`
vtctld_web_port=15000

# Set up environment.
export VTTOP=${VTTOP-$VTROOT/src/vitess.io/vitess}

# Try to find mysqld on PATH.
if [ -z "$VT_MYSQL_ROOT" ]; then
  mysql_path=`which mysqld`
  if [ -z "$mysql_path" ]; then
    echo "Can't guess location of mysqld. Please set VT_MYSQL_ROOT manually."
    exit 1
  fi
  export VT_MYSQL_ROOT=$(dirname `dirname $mysql_path`)
fi

# Previously the file specified MYSQL_FLAVOR
# it is now autodetected

if [ "${TOPO}" = "zk2" ]; then
    # Each ZooKeeper server needs a list of all servers in the quorum.
    # Since we're running them all locally, we need to give them unique ports.
    # In a real deployment, these should be on different machines, and their
    # respective hostnames should be given.
    echo "enter zk2 env"
    zkcfg=(\
        "1@$hostname:28881:38881:21811" \
        "2@$hostname:28882:38882:21812" \
        "3@$hostname:28883:38883:21813" \
        )
    printf -v zkcfg ",%s" "${zkcfg[@]}"
    zkcfg=${zkcfg:1}

    zkids='1 2 3'

    # Set topology environment parameters.
    ZK_SERVER="localhost:21811,localhost:21812,localhost:21813"
    # shellcheck disable=SC2034
    TOPOLOGY_FLAGS="-topo_implementation zk2 -topo_global_server_address ${ZK_SERVER} -topo_global_root /vitess/global"

    mkdir -p $VTDATAROOT/tmp
else
    echo "enter etcd2 env"

    case $(uname) in
      Linux)  etcd_platform=linux;;
      Darwin) etcd_platform=darwin;;
    esac

    ETCD_SERVER="localhost:2379"
    ETCD_VERSION=$(cat "${VTROOT}/dist/etcd/.installed_version")
    ETCD_BINDIR="${VTROOT}/dist/etcd/etcd-${ETCD_VERSION}-${etcd_platform}-amd64/"
    TOPOLOGY_FLAGS="-topo_implementation etcd2 -topo_global_server_address $ETCD_SERVER -topo_global_root /vitess/global"

    mkdir -p "${VTDATAROOT}/tmp"
    mkdir -p "${VTDATAROOT}/etcd"
fi
