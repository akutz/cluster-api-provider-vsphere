#!/bin/bash

# Copyright 2019 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

CENTOS_VERSIONS="7.3.1611 7.4.1708 7.5.1804 7.6.1810 latest"
for v in ${CENTOS_VERSIONS}; do \
  echo "building CentOS ${v}"
  IMAGE_VERSION="${v}" make
  docker push "gcr.io/cluster-api-provider-vsphere/extra/cloud-init/centos:${v}"
done
