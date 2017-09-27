NOROOT := -u $$(id -u):$$(id -g)
SRCDIR := /go/src/github.com/agalitsyn/38-context-propagation
DOCKERFLAGS := --rm=true $(NOROOT) -v $(CURDIR):$(SRCDIR) -w $(SRCDIR)
BUILD_IMAGE := golang
BUILD_IMAGE_TAG ?= 1.9

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo

.PHONY: docker-build
docker-build:
	docker run $(DOCKERFLAGS) $(BUILD_IMAGE):$(BUILD_IMAGE_TAG) make
