BUILDPATH ?= $(CURDIR)
GOENV   ?= CGO_ENABLED=0 GOOS=windows
GOBUILD = ${GOENV} go

.PHONY: build
build:
	cd $(BUILDPATH) && $(GOBUILD) build -i -o outfile
	