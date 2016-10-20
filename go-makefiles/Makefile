BINARIES_DIR ?= cmd
BUILD_DIR ?= bin
PKG_PATH ?= github.com/agalitsyn/go-example
BINARIES_PATH := $(PKG_PATH)/$(BINARIES_DIR)
PROJECT_PKGS := $$(go list ./... | grep -v /vendor/)
PROJECT_BINS := $$(go list $(BINARIES_PATH)/...)


install:
	go install $(PKG_PATH)/...

build:
	for bin in $(PROJECT_BINS); do \
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -x -v -o $(BINARIES_DIR)/$$(basename $$bin)/$(BUILD_DIR)/$$(basename $$bin) $$bin ; \
	done

test:
	for pkg in $(PROJECT_PKGS); do \
        golint $$pkg ; \
        go vet $$pkg ; \
        go test -cover -v $$pkg ; \
    done

run: install
	goreman start

sloccount:
	find . -path ./vendor -prune -o -name "*.go" -print0 | xargs -0 wc -l