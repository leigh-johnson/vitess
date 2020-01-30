#!/usr/bin/env python
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
"""This is the setup script for the submodules in the Vitess python client.
"""

__version__ = '4.0.1'
from setuptools import setup

setup(
      name="vitess",
      packages=["vtctl", "vtdb", "vtproto", "vttest", "util"],
      platforms='Linux',
      setup_requires=['wheel'],
      install_requires=[
            'protobuf',
            'grpcio'
      ],
      version=__version__
     )
