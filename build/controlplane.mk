# Copyright 2022 YANDEX LLC
# This is modified version of the software, made by the Crossplane Authors
# and available at: https://github.com/crossplane-contrib/provider-jet-template

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

KIND_CLUSTER_NAME ?= local-dev
CROSSPLANE_NAMESPACE ?= crossplane-system
# CROSSPLANE_VERSION ?= 2.0.2
CROSSPLANE_CHART_REPO ?= https://charts.crossplane.io/stable
CROSSPLANE_CHART_NAME ?= crossplane

CONTROLPLANE_DUMP_DIRECTORY ?= $(OUTPUT_DIR)/controlplane-dump

controlplane.up: $(HELM) $(KUBECTL) $(KIND)
	@$(INFO) setting up controlplane
	@$(KIND) get kubeconfig --name $(KIND_CLUSTER_NAME) >/dev/null 2>&1 || $(KIND) create cluster --name=$(KIND_CLUSTER_NAME)
	@$(INFO) "setting kubectl context to kind-$(KIND_CLUSTER_NAME)"
	@$(KUBECTL) config use-context "kind-$(KIND_CLUSTER_NAME)"
	@$(HELM) repo add crossplane-build-module $(CROSSPLANE_CHART_REPO) --force-update
	@$(HELM) repo update
ifndef CROSSPLANE_ARGS
	@$(INFO) setting up crossplane core without args
	@$(HELM) get notes -n $(CROSSPLANE_NAMESPACE) crossplane >/dev/null 2>&1 || $(HELM) install crossplane --create-namespace --namespace=$(CROSSPLANE_NAMESPACE) crossplane-build-module/$(CROSSPLANE_CHART_NAME)
else
	@$(INFO) setting up crossplane core with args $(CROSSPLANE_ARGS)
	@$(HELM) get notes -n $(CROSSPLANE_NAMESPACE) crossplane >/dev/null 2>&1 || $(HELM) install crossplane --create-namespace --namespace=$(CROSSPLANE_NAMESPACE) --set "args={${CROSSPLANE_ARGS}}" crossplane-build-module/$(CROSSPLANE_CHART_NAME)
endif

controlplane.down: $(KIND)
	@$(INFO) deleting controlplane
	@$(KIND) delete cluster --name=$(KIND_CLUSTER_NAME)
	@$(OK) deleting controlplane

controlplane.dump: $(KUBECTL)
	mkdir -p $(CONTROLPLANE_DUMP_DIRECTORY)
	@$(KUBECTL) cluster-info dump --output-directory $(CONTROLPLANE_DUMP_DIRECTORY) --all-namespaces || true
	@$(KUBECTL) get crossplane --all-namespaces > $(CONTROLPLANE_DUMP_DIRECTORY)/all-crossplane.txt || true
	@$(KUBECTL) get crossplane --all-namespaces -o yaml > $(CONTROLPLANE_DUMP_DIRECTORY)/all-crossplane.yaml || true
