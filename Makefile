SHELL := /usr/bin/env bash

.DEFAULT_GOAL := help

DOCKERCMD=$(shell which docker)
SWAGGER_VERSION=v0.30.3
SWAGGER := $(DOCKERCMD) run --rm -t -u "$(shell id -u):$(shell id -g)" -v $(shell pwd):/src -w /src quay.io/goswagger/swagger:$(SWAGGER_VERSION)

ifeq ($(VERSION),)
VERSION := v2.13.1
endif

HARBOR_2.0_SPEC=api/v2.0/swagger.yaml
HARBOR_CLIENT_2.0_DIR=pkg/sdk/v2.0
HARBOR_2.0_SPEC_URL=https://raw.githubusercontent.com/goharbor/harbor/$(VERSION)/api/v2.0/swagger.yaml

## --------------------------------------
## Help
## --------------------------------------

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Client
## --------------------------------------

.PHONY: update-spec
update-spec: ## update all swagger spec files
	@echo "Downloading spec from Harbor with version: $(VERSION)"
	@wget ${HARBOR_2.0_SPEC_URL} -O ${HARBOR_2.0_SPEC}

.PHONY: gen-harbor-api
gen-harbor-api: update-spec ## generate goswagger client for harbor
	@$(SWAGGER) generate client -f ${HARBOR_2.0_SPEC} --target=$(HARBOR_CLIENT_2.0_DIR) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC

.PHONY: cleanup
cleanup:
	rm --recursive --force pkg/sdk/v2.0/models pkg/sdk/v2.0/client

.PHONY: test
test: ## run the test
	go test ./...

get-go-version:
	@echo "$(VERSION)" | sed -E "s/v([0-9]+)\.([0-9]+)\.([0-9]+)/v0.\1\2.\3/g" -

all: cleanup gen-harbor-api
