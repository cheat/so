# paths
makefile := $(realpath $(lastword $(MAKEFILE_LIST)))
cmd_dir  := ./cmd/so
dist_dir := ./dist

# executables
CAT    := cat
COLUMN := column
CTAGS  := ctags
GO     := go
GREP   := grep
GZIP   := gzip --best
LINT   := revive
MKDIR  := mkdir -p
RM     := rm
SCC    := scc
SED    := sed
SORT   := sort
ZIP    := zip -m

# build flags
BUILD_FLAGS  := -ldflags="-s -w" -mod vendor -trimpath
GOBIN        :=
TMPDIR       := /tmp

# release binaries
releases :=                     \
	$(dist_dir)/so-darwin-amd64 \
	$(dist_dir)/so-linux-amd64  \
	$(dist_dir)/so-linux-arm5   \
	$(dist_dir)/so-linux-arm6   \
	$(dist_dir)/so-linux-arm7   \
	$(dist_dir)/so-windows-amd64.exe

## build: builds an executable for your architecture
.PHONY: build
build: $(dist_dir) clean vendor generate
	$(GO) build $(BUILD_FLAGS) -o $(dist_dir)/so $(cmd_dir)

## build-release: builds release executables
.PHONY: build-release
build-release: $(releases)

## ci: builds a "release" executable for the current architecture (used in ci)
.PHONY: ci
ci: | setup prepare build

# so-darwin-amd64
$(dist_dir)/so-darwin-amd64: prepare
	GOARCH=amd64 GOOS=darwin \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# so-linux-amd64
$(dist_dir)/so-linux-amd64: prepare
	GOARCH=amd64 GOOS=linux \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# so-linux-arm5
$(dist_dir)/so-linux-arm5: prepare
	GOARCH=arm GOOS=linux GOARM=5 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# so-linux-arm6
$(dist_dir)/so-linux-arm6: prepare
	GOARCH=arm GOOS=linux GOARM=6 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# so-linux-arm7
$(dist_dir)/so-linux-arm7: prepare
	GOARCH=arm GOOS=linux GOARM=7 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# so-windows-amd64
$(dist_dir)/so-windows-amd64.exe: prepare
	GOARCH=amd64 GOOS=windows \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(ZIP) $@.zip $@

# ./dist
$(dist_dir):
	$(MKDIR) $(dist_dir)

.PHONY: generate
generate:
	$(GO) generate $(cmd_dir)

## install: builds and installs so on your PATH
.PHONY: install
install: build
	$(GO) install $(BUILD_FLAGS) $(GOBIN) $(cmd_dir) 

## clean: removes compiled executables
.PHONY: clean
clean: $(dist_dir)
	$(RM) -f $(dist_dir)/*

## distclean: removes the tags file
.PHONY: distclean
distclean:
	$(RM) -f tags

## setup: installs revive (linter) and scc (sloc tool)
.PHONY: setup
setup:
	GO111MODULE=off $(GO) get -u github.com/boyter/scc github.com/mgechev/revive

## sloc: counts "semantic lines of code"
.PHONY: sloc
sloc:
	$(SCC) --exclude-dir=vendor

## tags: builds a tags file
.PHONY: tags
tags:
	$(CTAGS) -R --exclude=vendor --languages=go

## vendor: downloads, tidies, and verifies dependencies
.PHONY: vendor
vendor:
	$(GO) mod vendor && $(GO) mod tidy && $(GO) mod verify

## fmt: runs go fmt
.PHONY: fmt
fmt:
	$(GO) fmt ./...

## lint: lints go source files
.PHONY: lint
lint: vendor
	$(LINT) -exclude vendor/... ./...

## vet: vets go source files
.PHONY: vet
vet:
	$(GO) vet ./...

## test: runs unit-tests
.PHONY: test
test: 
	$(GO) test ./...

## coverage: generates a test coverage report
.PHONY: coverage
coverage:
	$(GO) test ./... -coverprofile=$(TMPDIR)/so-coverage.out && \
	$(GO) tool cover -html=$(TMPDIR)/so-coverage.out

## check: formats, lints, vets, vendors, and run unit-tests
.PHONY: check
check: | vendor fmt lint vet test

.PHONY: prepare
prepare: | $(dist_dir) clean generate vendor fmt lint vet test

## help: displays this help text
.PHONY: help
help:
	@$(CAT) $(makefile) | \
	$(SORT)             | \
	$(GREP) "^##"       | \
	$(SED) 's/## //g'   | \
	$(COLUMN) -t -s ':'
