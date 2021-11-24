DOCKERCMD=$(shell which docker)
SWAGGER := $(DOCKERCMD) run --rm -it -v $(HOME):$(HOME) -w $(shell pwd) quay.io/goswagger/swagger

ifeq ($(VERSION),)
VERSION := v2.4.0
endif

HARBOR_ASSIST_SPEC=api/swagger.yaml
HARBOR_CLIENT_ASSIST_DIR=pkg/sdk/assist
HARBOR_ASSIST_SPEC_URL=https://raw.githubusercontent.com/goharbor/harbor/$(VERSION)/api/swagger.yaml

HARBOR_2.0_SPEC=api/v2.0/swagger.yaml
HARBOR_CLIENT_2.0_DIR=pkg/sdk/v2.0
HARBOR_2.0_SPEC_URL=https://raw.githubusercontent.com/goharbor/harbor/$(VERSION)/api/v2.0/swagger.yaml

HARBOR_2.0_LEGACY_SPEC=api/v2.0/legacy_swagger.yaml
HARBOR_CLIENT_2.0_LEGACY_DIR=pkg/sdk/v2.0/legacy
HARBOR_2.0_LEGACY_SPEC_URL=https://raw.githubusercontent.com/goharbor/harbor/$(VERSION)/api/v2.0/legacy_swagger.yaml


update-swagger-spec:
	echo "Version: $(VERSION)"
	wget ${HARBOR_ASSIST_SPEC_URL} -O ${HARBOR_ASSIST_SPEC}
	wget ${HARBOR_2.0_SPEC_URL} -O ${HARBOR_2.0_SPEC}
	wget ${HARBOR_2.0_LEGACY_SPEC_URL} -O ${HARBOR_2.0_LEGACY_SPEC}

gen-harbor-api:
	@$(SWAGGER) generate client -f ${HARBOR_ASSIST_SPEC} --target=$(HARBOR_CLIENT_ASSIST_DIR) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC
	@$(SWAGGER) generate client -f ${HARBOR_2.0_SPEC} --target=$(HARBOR_CLIENT_2.0_DIR) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC
	@$(SWAGGER) generate client -f ${HARBOR_2.0_LEGACY_SPEC} --target=$(HARBOR_CLIENT_2.0_LEGACY_DIR) --template=stratoscale --additional-initialism=CVE --additional-initialism=GC --additional-initialism=OIDC
