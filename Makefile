DOCKER := docker
GO := go

BUF_IMAGE := bufbuild/buf:1.4.0

.PHONY: build
build: cmd/xavier cmd/xavier-etxea

.PHONY: cmd/xavier
cmd/xavier:
	$(GO) build -o bin/xavier ./cmd/xavier

.PHONY: cmd/xavier-etxea
cmd/xavier-etxea:
	$(GO) build -o bin/xavier-etxea ./cmd/xavier-etxea

buf-build:
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) build

buf-ls-files:
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) ls-files

buf-lint: 
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) lint

buf-generate: 
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) generate