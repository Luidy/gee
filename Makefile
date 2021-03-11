BUILDPATH ?= $(CURDIR)
BASE = $(BUILDPATH)
# GOENV   ?= CGO_ENABLED=0 GOOS=darwin
GOENV   ?= CGO_ENABLED=0 GOOS=linux
GOBUILD = ${GOENV} go

.PHONY: build
build:
	cd $(BUILDPATH) && $(GOBUILD) build -o $(BUILDPATH)/bin/gee

.PHONY: pkg.list
pkg.list:
ifeq ($(wildcard $(CURDIR)/pkg.list),)
	cd $(BASE) && GO111MODULE=on go list -f {{.Dir}} ./... > pkg.list
endif
	