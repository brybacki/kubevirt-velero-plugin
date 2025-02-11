#!/usr/bin/env bash

#Copyright 2021 The CDI Authors.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

set -e

if [ -z "$KUBEVIRTCI_PATH" ]; then
    KUBEVIRTCI_PATH="$(
        cd "$(dirname "$BASH_SOURCE[0]")/"
        echo "$(pwd)/"
    )"../../cluster-up/
fi

script_dir="$(cd "$(dirname "$0")" && pwd -P)"
velero_dir=${script_dir}/../velero
source "${script_dir}"/../config.sh

${velero_dir}/velero  \
  --kubeconfig $(pwd)/_ci-configs/${KUBEVIRT_PROVIDER}/.kubeconfig \
  plugin add ${IMAGE}:${VERSION}

sleep 15
${velero_dir}/velero  \
  --kubeconfig $(pwd)/_ci-configs/${KUBEVIRT_PROVIDER}/.kubeconfig \
  plugin get