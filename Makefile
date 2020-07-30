BUILDPATH ?= $(CURDIR)
GOENV   ?= CGO_ENABLED=0 GOOS=darwin
# GOENV   ?= CGO_ENABLED=0 GOOS=linux
GOBUILD = ${GOENV} go

.PHONY: build
build:
	cd $(BUILDPATH) && $(GOBUILD) build -i -o $(BUILDPATH)/bin/gee
	