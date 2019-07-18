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

all: build

IMAGE_VERSION ?= 7.6.1810
IMAGE_NAME ?= gcr.io/cluster-api-provider-vsphere/extra/cloud-init/centos
IMAGE_TAG ?= $(IMAGE_NAME):$(IMAGE_VERSION)
CLOUD_INIT_REF ?= 19.1

build:
	docker build \
	  --build-arg CENTOS_IMAGE=centos:$(IMAGE_VERSION) \
	  --build-arg CLOUD_INIT_REF=$(CLOUD_INIT_REF) \
	  -t $(IMAGE_TAG) .
.PHONY: build

push:
	docker push $(IMAGE_TAG)
.PHONY: push
