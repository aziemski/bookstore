## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

PATH  := $(PATH):$(LOCALBIN)

SHELL=/bin/sh
.SHELLFLAGS = -e -c

GOBIN	:= $(shell pwd)/bin
GO 		?= $(shell which go)
.DEFAULT_GOAL := help

LDFLAGS 	:= ""
TESTS 		:= .
TESTFLAGS	:= -race -cover -coverprofile=coverage.out

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## Tools Binaries
OAPI_CODEGEN ?= $(LOCALBIN)/oapi-codegen
MOQ ?= $(LOCALBIN)/moq
GOLANGCI_LINT ?= $(LOCALBIN)/golangci-lint

.PHONY: moq
moq: $(MOQ) ## Download moq locally if necessary.
$(MOQ): $(LOCALBIN)
	test -s $(LOCALBIN)/moq || GOBIN=$(LOCALBIN) go install github.com/matryer/moq@v0.3.4

.PHONY: oapi_codegen
oapi_codegen: $(OAPI_CODEGEN) ## Download oapi_codegen locally if necessary.
$(OAPI_CODEGEN): $(LOCALBIN)
	test -s $(LOCALBIN)/oapi_codegen || GOBIN=$(LOCALBIN) go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT) ## Download golangci-lint locally if necessary.
$(GOLANGCI_LINT): $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.1

.PHONY: download-tools
download-tools: ## Download all go tools
download-tools: moq oapi_codegen golangci-lint

.PHONY: generate
generate: oapi_codegen moq ## Generate boilerplate
	$(GO) generate ./...

.PHONY: build
build: ## Build app
	GOBIN=$(LOCALBIN) $(GO) install -ldflags $(LDFLAGS)  ./...

.PHONY: run
run:
	$(GO) run .

.PHONY: lint
lint: golangci-lint ## Run lint.
	$(GOLANGCI_LINT) run -v ./...

.PHONY: fmt
fmt: golangci-lint ## Fix lint errors.
	$(GOLANGCI_LINT) run -v --fix ./...

.PHONY: test
test:  ## Run tests
	$(GO) test -race -cover -coverprofile=coverage.out ./...

.PHONY: coverage
coverage: ## Show coverage summary (CI contract)
	$(GO) tool cover -func=coverage.out

.PHONY: untracked
untracked: generate ## Check for no untracked files
	git status
	git --no-pager diff
	git diff-index --quiet HEAD --
