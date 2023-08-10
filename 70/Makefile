GIT_HASH := $(shell git log --format="%h" -n 1 2> /dev/null)

.PHONY: test
test:
	go build -ldflags "-X main.commitSHA=$(GIT_HASH)" -buildvcs -o 70 main.go
	./70
