NOROOT := -u $$(id -u):$$(id -g)
SRCDIR := /go/src/github.com/agalitsyn/play/38-context-propagation
DOCKERFLAGS := --rm=true $(NOROOT) -v $(CURDIR):$(SRCDIR) -w $(SRCDIR)
BUILD_IMAGE := golang
BUILD_IMAGE_TAG ?= 1.9

.PHONY: build-osx
build-osx:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo

.PHONY: docker-build-osx
docker-build-osx:
	docker run $(DOCKERFLAGS) $(BUILD_IMAGE):$(BUILD_IMAGE_TAG) make build-osx

.PHONY: docker-build-linux
docker-build-linux:
	docker run $(DOCKERFLAGS) $(BUILD_IMAGE):$(BUILD_IMAGE_TAG) make build-linux

.PHONY: docker-run-linux
docker-run-linux:
	docker run $(DOCKERFLAGS) debian:stable bash run.sh
