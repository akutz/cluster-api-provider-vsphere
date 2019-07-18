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

################################################################################
##                               BUILD ARGS                                   ##
################################################################################
# The CentOS image on which this image is based.
ARG CENTOS_IMAGE=centos:7.6.1810

################################################################################
##                                  MAIN                                      ##
################################################################################
FROM ${CENTOS_IMAGE}
LABEL "maintainer" "Andrew Kutz <akutz@vmware.com>"

# Install the EPEL repo, development tools, curl, git, and pip.
RUN yum -y install curl git && \
    yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm && \
    yum -y groupinstall 'Development Tools' && \
    yum -y install python2-pip

# Install the packages needed to build the cloud-init RPM.
RUN yum -y install \
  python-configobj \
  python-oauthlib \
  python-six \
  PyYAML \
  python-jsonpatch \
  python-jinja2 \
  python-jsonschema \
  python-requests \
  e2fsprogs \
  iproute \
  net-tools \
  rsyslog \
  sudo \
  python-devel

# Clone the cloud-init repo and install its Python requirements with pip.
# This will override some of the above packages, but that's required because
# the newer Python libs are needed while the RPM build process also requires
# the presence of the above packages.
WORKDIR /cloud-init
RUN git clone https://github.com/number5/cloud-init.git . && \
    pip install --force --upgrade --ignore-installed -r requirements.txt
ENV SKIP_UNCOMITTED_CHANGES_CHECK=1

# Checkout the cloud-init ref to build.
ARG CLOUD_INIT_REF=19.1
RUN git checkout -b build/${CLOUD_INIT_REF} ${CLOUD_INIT_REF}

# Build the binary and source RPMs.
RUN make rpm
