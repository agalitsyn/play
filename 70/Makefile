GIT_HASH := $(shell git log --format="%h" -n 1 2> /dev/null)
GIT_TAG := $(shell git describe --tags --always --abbrev=0)

BUILD_ARGS ?= -ldflags \
	"-X gitlab.ptsecurity.com/cybsi/box/internal/version.Tag=$(GIT_TAG)"



.PHONY: fail
fail:
	go build -ldflags "-X main.commitSHA=$(GIT_HASH) -X main.Tag=$(GIT_TAG)" -o 70 main.go
	./70

.PHONY: success
success:
	go build -ldflags "-X main.commitSHA=$(GIT_HASH) -X main.Tag=$(GIT_TAG)" -o 70 ./...
	./70
