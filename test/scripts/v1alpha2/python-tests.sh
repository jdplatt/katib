#!/bin/bash

# Copyright 2018 The Kubernetes Authors.
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

# This shell script is used to run the python tests in the argo workflow

exit 0

pip install -r cmd/suggestion/v1alpha2/katib-suggestion/requirements.txt
pip install -r pkg/suggestion/test_requirements.txt
pushd pkg/suggestion
python setup.py develop
popd
pushd pkg/api/python
python setup.py develop
popd
pylint pkg/suggestion/v1alpha2/katib_suggestion --disable=fixme --exit-zero --reports=y
pytest pkg/suggestion/v1alpha2/tests --verbose --cov=pkg/suggestion/v1alpha2/katib_suggestion --cov-report term-missing
