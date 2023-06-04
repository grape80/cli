MAKEFLAGS += --no-builtin-rules
.DEFAULT_GOAL := help

include .build/env
logDir := $(LOG_DIR)
logTestDir := $(LOG_TEST_DIR)

now = $(shell date '+%Y%m%d-%H%M%S')

.PHONY: help ## Show help.
help:
	@cat $(MAKEFILE_LIST) | grep '##' | grep -v 'MAKEFILE_LIST' | sed s/^.PHONY:// | awk -F \#\# '{ printf "%-20s%s\n", $$1, $$2 }'

.PHONY: gotest ## Run go test.
gotest:
	mkdir -p $(logTestDir)
	go test -v ./... -coverpkg=./... -count=1 -coverprofile=$(logTestDir)/gocover-$(now).out | tee $(logTestDir)/gotest-$(now).log
	go tool cover -html=$(logTestDir)/gocover-$(now).out -o $(logTestDir)/gocover-$(now).html

.PHONY: gobench ## Run go benchmark.
gobench:
	mkdir -p $(logTestDir)
	go test -bench . ./example/... -benchmem | tee $(logTestDir)/gobench-$(now).log

.PHONY: test ## Run test.
test: gotest gobench

.PHONY: cicd ## Run CI/CD.
cicd: test

.PHONY: cleanlog ## Clean log.
cleanlog:
	rm -rfv $(logDir)
