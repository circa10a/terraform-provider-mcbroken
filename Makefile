NAMESPACE=circa10a
PROVIDER_NAME=mcbroken
PROVIDER_FULL_PATH=$(REGISTRY)/$(NAMESPACE)/$(PROVIDER_NAME)
PROVIDER_FULL_NAME=terraform-provider-$(PROVIDER_NAME)
PROJECT=$(NAMESPACE)/$(PROVIDER_FULL_NAME)
VERSION=0.1.3

build-mac: PLUGIN_DIR = ~/.terraform.d/plugins/local/provider/$(PROVIDER_NAME)/$(VERSION)/darwin_amd64
build-mac:
	test -d $(PLUGIN_DIR) || mkdir -p $(PLUGIN_DIR)
	go build -o $(PLUGIN_DIR)/$(PROVIDER_FULL_NAME)

build-linux: PLUGIN_DIR = ~/.terraform.d/plugins/local/provider/$(PROVIDER_NAME)/$(VERSION)/linux_amd64
build-linux:
	test -d $(PLUGIN_DIR) || mkdir -p $(PLUGIN_DIR)
	go build -o $(PLUGIN_DIR)/$(PROVIDER_FULL_NAME)

lint:
	@if ! command -v golangci-lint 1>/dev/null; then\
		echo "Need to install golangci-lint";\
		exit 1;\
	fi;\
	golangci-lint run