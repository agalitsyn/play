#!/bin/bash

set -ex

rm testbin || true

# Selecting concrete files from one package is working
# go test -coverprofile=coverage.out -c main.go main_test.go lib.go -o testbin

# Selecting concrete files from different packages is not working
# ERR: named files must all be in one directory; have ./ and lib/
# go test -coverprofile=coverage.out -c main.go main_test.go lib/lib.go -o testbin

# Selecting by packages works
# go test -coverpkg=.,github.com/agalitsyn/play/61-go-test-external/lib -c -o testbin

# If main in another folder, not works
# go test -coverpkg=./... -c -o ../testbin

packages="$(go list ./... | grep -v cmd | tr '\n' ',' | sed 's/,$/\n/')"

cd cmd

# Selecting by packages works
go test -coverpkg=.,"$packages" -c -o ../testbin
cd ../

./testbin -test.coverprofile=coverage.out -test.v -test.run=TestMain
