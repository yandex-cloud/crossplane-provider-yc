# ====================================================================================
# Setup Project

PROJECT_NAME := crossplane-provider-yc
PROJECT_REPO := github.com/yandex-cloud/$(PROJECT_NAME)

INTERNAL_CR ?= crpih38ka022n1ng31n0

export TERRAFORM_VERSION := 1.6.1

export TERRAFORM_PROVIDER_HOST ?= terraform-mirror.yandexcloud.net/registry.terraform.io
export TERRAFORM_MIRROR_HOST ?= hashicorp-releases.yandexcloud.net
export TERRAFORM_PROVIDER_SOURCE := yandex-cloud/yandex
export TERRAFORM_PROVIDER_REPO ?= https://github.com/yandex-cloud/terraform-provider-yandex
export TERRAFORM_PROVIDER_VERSION := 0.117.0
export TERRAFORM_PROVIDER_DOWNLOAD_NAME := terraform-provider-yandex
export TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX := https://$(TERRAFORM_PROVIDER_HOST)/$(TERRAFORM_PROVIDER_SOURCE)
export TERRAFORM_DOCS_PATH ?= website/docs/r
export TERRAFORM_NATIVE_PROVIDER_BINARY ?= terraform-provider-yandex_$(TERRAFORM_PROVIDER_VERSION)_x5
# the version of chainsaw to use
export CHAINSAW_VERSION ?= 0.2.0
export CROSSPLANE_CLI_VERSION ?= current
export CROSSPLANE_CLI_CHANNEL ?= stable



PLATFORMS ?= linux_amd64

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# ====================================================================================
# Setup Output

-include build/makelib/output.mk

# ====================================================================================
# Setup Go

# Set a sane default so that the nprocs calculation below is less noisy on the initial
# loading of this file
NPROCS ?= 1

# each of our test suites starts a kube-apiserver and running many test suites in
# parallel can lead to high CPU utilization. by default we reduce the parallelism
# to half the number of CPU cores.
GO_TEST_PARALLEL := $(shell echo $$(( $(NPROCS) / 2 )))

GO_REQUIRED_VERSION ?= 1.21
GOLANGCILINT_VERSION ?= 1.55.1
GO_STATIC_PACKAGES = $(GO_PROJECT)/cmd/provider $(GO_PROJECT)/cmd/generator
GO_LDFLAGS += -X $(GO_PROJECT)/internal/version.Version=$(VERSION)
GO_SUBDIRS += cmd internal apis
GO111MODULE = on
-include build/makelib/golang.mk

# ====================================================================================
# Setup Kubernetes tools

KIND_VERSION = v0.19.0
UP_VERSION = v0.28.0
UP_CHANNEL = stable
UPTEST_VERSION = v1.1.2
-include build/makelib/k8s_tools.mk

# ====================================================================================
# Setup Images

REGISTRY_ORGS ?= cr.yandex/$(INTERNAL_CR)/yandex-cloud xpkg.upbound.io/yandexcloud
IMAGES = $(PROJECT_NAME)
-include build/makelib/imagelight.mk

# ====================================================================================
# Setup XPKG

XPKG_REG_ORGS ?= cr.yandex/$(INTERNAL_CR)/yandex-cloud/crossplane xpkg.upbound.io/yandexcloud
# NOTE(hasheddan): skip promoting on xpkg.upbound.io as channel tags are
# inferred.
XPKG_REG_ORGS_NO_PROMOTE ?= xpkg.upbound.io/upbound
XPKGS = $(PROJECT_NAME)
-include build/makelib/xpkg.mk

# ====================================================================================
# Fallthrough

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# NOTE(hasheddan): we force image building to happen prior to xpkg build so that
# we ensure image is present in daemon.
xpkg.build.crossplane-provider-yc: do.build.images

# NOTE(hasheddan): we ensure up is installed prior to running platform-specific
# build steps in parallel to avoid encountering an installation race condition.
build.init: $(UP)

# ====================================================================================
# Setup Terraform for fetching provider schema
TERRAFORM := $(TOOLS_HOST_DIR)/terraform-$(TERRAFORM_VERSION)
TERRAFORM_WORKDIR := $(WORK_DIR)/terraform
TERRAFORM_PROVIDER_SCHEMA := config/schema.json
CHAINSAW := $(TOOLS_HOST_DIR)/chainsaw-$(CHAINSAW_VERSION)
CROSSPLANE_CLI := $(TOOLS_HOST_DIR)/crossplane-cli-$(CROSSPLANE_CLI_VERSION)

CROSSPLANE_UPTEST := $(TOOLS_HOST_DIR)/crossplane-uptest-$(UPTEST_VERSION)

# override target from k8s_tools
$(CROSSPLANE_UPTEST):
	@$(INFO) installing uptest $(CROSSPLANE_UPTEST)
	@curl -fsSLo $(CROSSPLANE_UPTEST) https://github.com/crossplane/uptest/releases/download/$(UPTEST_VERSION)/uptest_$(SAFEHOSTPLATFORM) || $(FAIL)
	@chmod +x $(CROSSPLANE_UPTEST)
	@$(OK) installing uptest $(CROSSPLANE_UPTEST)

$(TERRAFORM):
	@$(INFO) installing terraform $(HOSTOS)-$(HOSTARCH)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-terraform
	@curl -fsSL https://$(TERRAFORM_MIRROR_HOST)/terraform/$(TERRAFORM_VERSION)/terraform_$(TERRAFORM_VERSION)_$(SAFEHOST_PLATFORM).zip -o $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip
	@unzip $(TOOLS_HOST_DIR)/tmp-terraform/terraform.zip -d $(TOOLS_HOST_DIR)/tmp-terraform
	@mv $(TOOLS_HOST_DIR)/tmp-terraform/terraform $(TERRAFORM)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-terraform
	@$(OK) installing terraform $(HOSTOS)-$(HOSTARCH)

# chainsaw download and install
$(CHAINSAW):
	@$(INFO) installing chainsaw $(CHAINSAW_VERSION)
	@mkdir -p $(TOOLS_HOST_DIR)
	@curl -fsSLo $(CHAINSAW).tar.gz --create-dirs https://github.com/kyverno/chainsaw/releases/download/v$(CHAINSAW_VERSION)/chainsaw_$(SAFEHOST_PLATFORM).tar.gz || $(FAIL)
	@tar -xvf $(CHAINSAW).tar.gz chainsaw
	@mv chainsaw $(CHAINSAW)
	@chmod +x $(CHAINSAW)
	@rm $(CHAINSAW).tar.gz
	@$(OK) installing chainsaw $(CHAINSAW_VERSION)

# Crossplane CLI download and install
$(CROSSPLANE_CLI):
	@$(INFO) installing Crossplane CLI $(CROSSPLANE_CLI_VERSION)
	@curl -fsSLo $(CROSSPLANE_CLI) --create-dirs https://releases.crossplane.io/$(CROSSPLANE_CLI_CHANNEL)/$(CROSSPLANE_CLI_VERSION)/bin/$(SAFEHOST_PLATFORM)/crank?source=build || $(FAIL)
	@chmod +x $(CROSSPLANE_CLI)
	@$(OK) installing Crossplane CLI $(CROSSPLANE_CLI_VERSION)


$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)
	@$(INFO) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)
	@mkdir -p $(TERRAFORM_WORKDIR)
	@echo '{"terraform":[{"required_providers":[{"provider":{"source":"'"$(TERRAFORM_PROVIDER_SOURCE)"'","version":"'"$(TERRAFORM_PROVIDER_VERSION)"'"}}],"required_version":"'"$(TERRAFORM_VERSION)"'"}]}' > $(TERRAFORM_WORKDIR)/main.tf.json
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) init -upgrade > $(TERRAFORM_WORKDIR)/terraform-logs.txt 2>&1
	@$(TERRAFORM) -chdir=$(TERRAFORM_WORKDIR) providers schema -json=true > $(TERRAFORM_PROVIDER_SCHEMA) 2>> $(TERRAFORM_WORKDIR)/terraform-logs.txt
	@$(OK) generating provider schema for $(TERRAFORM_PROVIDER_SOURCE) $(TERRAFORM_PROVIDER_VERSION)


pull-docs:
	@if [ ! -d "$(WORK_DIR)/$(notdir $(TERRAFORM_PROVIDER_REPO))" ]; then \
		git clone -c advice.detachedHead=false --depth 1 --filter=blob:none --branch "v$(TERRAFORM_PROVIDER_VERSION)" --sparse "$(TERRAFORM_PROVIDER_REPO)" "$(WORK_DIR)/$(notdir $(TERRAFORM_PROVIDER_REPO))"; \
	fi
	@git -C "$(WORK_DIR)/$(notdir $(TERRAFORM_PROVIDER_REPO))" sparse-checkout set "$(TERRAFORM_DOCS_PATH)"
	@./scripts/add_subcategory_html.sh

generate.init: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs

.PHONY: $(TERRAFORM_PROVIDER_SCHEMA) pull-docs
# ====================================================================================
# Targets


# NOTE: the build submodule currently overrides XDG_CACHE_HOME in order to
# force the Helm 3 to use the .work/helm directory. This causes Go on Linux
# machines to use that directory as the build cache as well. We should adjust
# this behavior in the build submodule because it is also causing Linux users
# to duplicate their build cache, but for now we just make it easier to identify
# its location in CI so that we cache between builds.
go.cachedir:
	@go env GOCACHE

# Generate a coverage report for cobertura applying exclusions on
# - generated file
cobertura:
	@cat $(GO_TEST_OUTPUT)/coverage.txt | \
		grep -v zz_ | \
		$(GOCOVER_COBERTURA) > $(GO_TEST_OUTPUT)/cobertura-coverage.xml

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# This is for running out-of-cluster locally, and is for convenience. Running
# this make target will print out the command which was used. For more control,
# try running the binary directly with different arguments.
run: go.build
	@$(INFO) Running Crossplane locally out-of-cluster . . .
	@# To see other arguments that can be provided, run the command with --help instead
	$(GO_OUT_DIR)/provider --debug

# ====================================================================================
# End to End Testing
CROSSPLANE_NAMESPACE = upbound-system
-include build/makelib/local.xpkg.mk
-include build/makelib/controlplane.mk

# This target requires the following environment variables to be set:
UPTEST_EXAMPLE_LIST ?= $(shell ./hack/examples.sh ./examples)
# - UPTEST_EXAMPLE_LIST, a comma-separated list of examples to test
#   To ensure the proper functioning of the end-to-end test resource pre-deletion hook, it is crucial to arrange your resources appropriately. 
#   You can check the basic implementation here: https://github.com/upbound/uptest/blob/main/internal/templates/01-delete.yaml.tmpl.
UPTEST_CLOUD_CREDENTIALS ?= $(shell cat ./key.json)
# - UPTEST_CLOUD_CREDENTIALS (optional), multiple sets of AWS IAM User credentials specified as key=value pairs.
#   The support keys are currently `DEFAULT` and `PEER`. So, an example for the value of this env. variable is:
#   DEFAULT='[default]
#   aws_access_key_id = REDACTED
#   aws_secret_access_key = REDACTED'
#   PEER='[default]
#   aws_access_key_id = REDACTED
#   aws_secret_access_key = REDACTED'
#   The associated `ProviderConfig`s will be named as `default` and `peer`.
UPTEST_DATASOURCE_PATH ?= $(shell ./hack/uptest_data.sh)
# - UPTEST_DATASOURCE_PATH (optional), see https://github.com/upbound/uptest#injecting-dynamic-values-and-datasource
# - CLOUD_ID and FOLDER_ID need to be the IDs of YC cloud and folder, respectively, where tests will be run.

uptest: $(CROSSPLANE_UPTEST) $(KUBECTL) $(CHAINSAW) $(CROSSPLANE_CLI)
	@echo "##teamcity[blockOpened name='uptest' description='run automated e2e tests']"
	@$(INFO) running automated tests
	@KUBECTL=$(KUBECTL) CHAINSAW=$(CHAINSAW) CROSSPLANE_CLI=$(CROSSPLANE_CLI) CROSSPLANE_NAMESPACE=$(CROSSPLANE_NAMESPACE) CREDENTIALS='$(UPTEST_CLOUD_CREDENTIALS)' $(CROSSPLANE_UPTEST) e2e "${UPTEST_EXAMPLE_LIST}" --data-source="${UPTEST_DATASOURCE_PATH}" --setup-script=cluster/test/setup.sh --default-conditions="Test" || $(FAIL)
	@$(OK) running automated tests
	@echo "##teamcity[blockClosed name='uptest']"

controlplane.up-cloud:$(UP) $(KUBECTL)
	@echo "##teamcity[blockOpened name='crossplane' description='set up Crossplane']"
	@$(INFO) setting up controlplane
	@$(KUBECTL) -n upbound-system get cm universal-crossplane-config >/dev/null 2>&1 || $(UP) uxp install
	@$(KUBECTL) -n upbound-system wait deploy crossplane --for condition=Available --timeout=120s
	@$(OK) setting up controlplane
	@echo "##teamcity[blockClosed name='crossplane']"

local-deploy: build controlplane.up local.xpkg.deploy.provider.$(PROJECT_NAME)
	@$(INFO) running locally built provider
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Healthy --timeout 5m
	@$(KUBECTL) -n upbound-system wait --for=condition=Available deployment --all --timeout=5m
	@$(OK) running locally built provider

cloud-reg:
	$(eval REGISTRY:=cr.yandex/$(shell yc container registry get crossplane-e2e-cr --format json | jq -r .id))

cloud.xpkg.deploy.provider: xpkg.push
	@echo "##teamcity[blockOpened name='deploy' description='deploy provider']"
	@$(INFO) deploying provider package $(PROJECT_NAME) $(VERSION)
	@echo '{"apiVersion":"pkg.crossplane.io/v1alpha1","kind":"ControllerConfig","metadata":{"name":"config"},"spec":{"args":["-d"],"image":"$(REGISTRY)/$(PROJECT_NAME)"}}' | $(KUBECTL) apply -f -
	@echo '{"apiVersion":"pkg.crossplane.io/v1","kind":"Provider","metadata":{"name":"$(PROJECT_NAME)"},"spec":{"package":"$(REGISTRY)/$(PROJECT_NAME)","controllerConfigRef":{"name":"config"}}}' | $(KUBECTL) apply -f -
	@$(OK) deploying provider package $(PROJECT_NAME) $(VERSION)

xpkg.push: $(UP) 
	@echo "##teamcity[blockOpened name='push' description='push provider image']"
	@$(INFO) pushing provider package $(PROJECT_NAME) $(VERSION)
	@$(UP) xpkg push $(REGISTRY)/$(PROJECT_NAME) -f $(XPKG_OUTPUT_DIR)/$(PLATFORM)/$(PROJECT_NAME)-$(VERSION).xpkg || $(FAIL)
	@echo $(REGISTRY)
	@$(OK) pushing provider package $(PROJECT_NAME) $(VERSION)
	@echo "##teamcity[blockClosed name='push']"

cloud-deploy: tc-build controlplane.up-cloud cloud.xpkg.deploy.provider
	@$(INFO) running locally built provider
	$(eval export PATH=$(PATH):$(TOOLS_HOST_DIR))
	@ln -s $(KUBECTL) $(TOOLS_HOST_DIR)/kubectl
	@$(KUBECTL) wait provider.pkg $(PROJECT_NAME) --for condition=Healthy --timeout 5m
	@$(KUBECTL) -n upbound-system wait --for=condition=Available deployment --all --timeout=5m
	@$(OK) running locally built provider
	@echo "##teamcity[blockClosed name='deploy']"

pre-build:
	@echo "##teamcity[blockOpened name='build' description='build provider image']"

tc-build: pre-build build
	@echo "##teamcity[blockClosed name='build']"


e2e: local-deploy uptest

e2e-cloud: cloud-reg cloud-deploy uptest

crddiff: $(CROSSPLANE_UPTEST)
	@$(INFO) Checking breaking CRD schema changes
	@for crd in $${MODIFIED_CRD_LIST}; do \
		if ! git cat-file -e "$${GITHUB_BASE_REF}:$${crd}" 2>/dev/null; then \
			echo "CRD $${crd} does not exist in the $${GITHUB_BASE_REF} branch. Skipping..." ; \
			continue ; \
		fi ; \
		echo "Checking $${crd} for breaking API changes..." ; \
		changes_detected=$(CROSSPLANE_UPTEST) crddiff revision <(git cat-file -p "$${GITHUB_BASE_REF}:$${crd}") "$${crd}" 2>&1) ; \
		if [[ $$? != 0 ]] ; then \
			printf "\033[31m"; echo "Breaking change detected!"; printf "\033[0m" ; \
			echo "$${changes_detected}" ; \
			echo ; \
		fi ; \
	done
	@$(OK) Checking breaking CRD schema changes

schema-version-diff:
	@$(INFO) Checking for native state schema version changes
	@export PREV_PROVIDER_VERSION=$$(git cat-file -p "${GITHUB_BASE_REF}:Makefile" | sed -nr 's/^export[[:space:]]*TERRAFORM_PROVIDER_VERSION[[:space:]]*:=[[:space:]]*(.+)/\1/p'); \
	echo Detected previous Terraform provider version: $${PREV_PROVIDER_VERSION}; \
	echo Current Terraform provider version: $${TERRAFORM_PROVIDER_VERSION}; \
	mkdir -p $(WORK_DIR); \
	git cat-file -p "$${GITHUB_BASE_REF}:config/schema.json" > "$(WORK_DIR)/schema.json.$${PREV_PROVIDER_VERSION}"; \
	./scripts/version_diff.py config/generated.lst "$(WORK_DIR)/schema.json.$${PREV_PROVIDER_VERSION}" config/schema.json
	@$(OK) Checking for native state schema version changes

.PHONY: cobertura submodules fallthrough run crds.clean

# ====================================================================================
# Special Targets

define CROSSPLANE_MAKE_HELP
Crossplane Targets:
    cobertura             Generate a coverage report for cobertura applying exclusions on generated files.
    submodules            Update the submodules, such as the common build scripts.
    run                   Run crossplane locally, out-of-cluster. Useful for development.

endef
# The reason CROSSPLANE_MAKE_HELP is used instead of CROSSPLANE_HELP is because the crossplane
# binary will try to use CROSSPLANE_HELP if it is set, and this is for something different.
export CROSSPLANE_MAKE_HELP

crossplane.help:
	@echo "$$CROSSPLANE_MAKE_HELP"

help-special: crossplane.help

.PHONY: crossplane.help help-special

.PHONY: test-all

test-all: lint test

up-login: $(UP)
	@$(UP) login