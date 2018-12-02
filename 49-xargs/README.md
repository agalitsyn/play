# xargs

Run `go run main.go --help` for details. Also you might want to compile it using `go build`.

## Examples

```bash
# classic
$ for i in {1..60}; do echo $i; done | go run main.go echo {}

# with rate limiting (will took ~4s)
$ for i in {1..60} ; do echo $i; done | go run main.go -rate 15 echo {}

# with parallel processes (will took ~1s)
$ for i in {1..60} ; do echo $i; done | go run main.go -rate 15 -inflight 4 echo {}

# will took ~2s
$ (echo 1; echo 2; echo 3; echo 4) | go run main.go -rate 1 -inflight 2 -debug -timer echo {}

# will took ~4s
$ (echo 1 ; sleep 3 ; echo 2 ; echo 3) | go run main.go -rate 1 -inflight 2 echo {}
```
