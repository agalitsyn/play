.PHONY: all
all: build-stdout build-file

.PHONY: build-stdout
build-stdout:
	docker build --rm --tag "test-stdout:latest" --file stdout.Dockerfile .


.PHONY: build-file
build-file:
	docker build --rm --tag "test-file:latest" --file file.Dockerfile .
