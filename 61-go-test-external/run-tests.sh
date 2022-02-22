#!/bin/bash

set -ex

rm coverage.out || true

curl -X POST localhost:1234/test
curl -X POST localhost:1234/deathblow

go tool cover -func=coverage.out

# Output:
# github.com/agalitsyn/play/61-go-test-external/lib.go:5:		Echo		100.0%
# github.com/agalitsyn/play/61-go-test-external/lib/lib.go:5:	Echo		100.0%
# github.com/agalitsyn/play/61-go-test-external/main.go:36:	runMain		100.0%
# github.com/agalitsyn/play/61-go-test-external/main.go:50:	main		93.3%
# total:								(statements)	95.5%
