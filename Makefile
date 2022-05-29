DOCKER := docker

BUF_IMAGE := bufbuild/buf:1.4.0

buf-build:
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) build

buf-ls-files:
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) ls-files

buf-lint: 
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) lint

buf-generate: 
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) generate