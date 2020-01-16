SHELL := /bin/bash
DEP_PRESENT := $(shell command -v dep 2> /dev/null)
GOIMPORTS_PRESENT := $(shell command -v goimports 2> /dev/null)
GOMETALINT_PRESENT := $(shell command -v golangci-lint 2> /dev/null)
export VERSION ?= 0.0.1
TEST_NAME ?= all
RELEASE_NOTES ?= notes/$(VERSION).md
TEST_UNIT_FLAGS ?= -timeout 10s -p 4 -race -cover -coverprofile=reports/c.out
TEST_UNIT_PACKAGE ?= ./...
LINT_FOLDERS ?= $(shell go list ./... | sed 's|github.com/$(OWNER)/$(REPO)/||' | grep -v github.com/$(OWNER)/$(REPO))
LINT_CHANGES = $(shell gofmt -d -e -s $(LINT_FOLDERS))
REPORT_PATH := reports

define HELP

$(BINARY) v$(VERSION) Makefile
=================================

## Development targets

- format:                 Formats the codebase according to gofmt and goimports
- test:                   Runs unit tests.
- lint:                   Runs golanccli-lint against all packages

endef
export HELP

.DEFAULT: help
.PHONY: help
help:
	@ echo "$$HELP"

.PHONY: deps
deps:
ifndef DEP_PRESENT
	@ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
ifndef GOIMPORTS_PRESENT
	@ go get -u golang.org/x/tools/cmd/goimports
endif
ifndef GOMETALINT_PRESENT
	@ go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif

.PHONY: vendor
vendor: deps
	@ echo "-> Installing $(BINARY) dependencies..."
	@ dep ensure -v

.PHONY: test
test: _report_path
	@ go test $(TEST_UNIT_FLAGS) $(TEST_UNIT_PACKAGE)

.PHONY: format
format: deps fix-imports
	@ gofmt -e -w -s $(LINT_FOLDERS)

.PHONY: fix-imports
fix-imports:
	@ goimports -w $(LINT_FOLDERS)

.PHONY: clean
clean:
	@ rm -rf vendor reports

.PHONY: _report_path
_report_path:
	@ mkdir -p $(REPORT_PATH)

.PHONY: update-vendor
update-vendor:
	@ echo "-> Upgrading vendors that don't have a fixed version..."
	@ dep ensure -update

.PHONY: unit-coverage
unit-coverage: _report_path
	@ echo "-> Generating coverage report..."
	@ go tool cover -html=$(REPORT_PATH)/c.out -o $(REPORT_PATH)/coverage.html

.PHONY: lint
lint: deps
	@ golangci-lint run --enable-all -D errcheck,dupl,lll,gochecknoglobals,maligned,gochecknoinits,gosec,scopelint,goconst --max-issues-per-linter=0 --max-same-issues=0 --out-format=tab
