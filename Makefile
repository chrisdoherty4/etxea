DOCKER := docker
GO := go

BUF_IMAGE := bufbuild/buf:1.4.0

.PHONY: build
build: cmd/etxea cmd/etxea-gela

.PHONY: cmd/etxea
cmd/etxea:
	$(GO) build -o bin/etxea ./cmd/etxea

.PHONY: cmd/etxea-gela
cmd/etxea-gela:
	$(GO) build -o bin/etxea-gela ./cmd/etxea-gela

buf-build:
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) build

buf-ls-files:
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) ls-files

buf-lint: 
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) lint

buf-generate: 
	$(DOCKER) run -v$$PWD:$$PWD -w$$PWD $(BUF_IMAGE) generate