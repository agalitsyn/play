see [article about go 1.18 build features](https://shibumi.dev/posts/go-18-feature/)

It should look like this
```
❯ go version -m ./test
./test: go1.18
	path	github.com/shibumi/test
	mod	github.com/shibumi/test	(devel)
	build	-compiler=gc
	build	CGO_ENABLED=1
	build	CGO_CFLAGS=
	build	CGO_CPPFLAGS=
	build	CGO_CXXFLAGS=
	build	CGO_LDFLAGS=
	build	GOARCH=amd64
	build	GOOS=linux
	build	GOAMD64=v1
	build	vcs=git
	build	vcs.revision=7e22e19e829d84170072d2459e5870876df495ed
	build	vcs.time=2022-04-03T16:59:50Z
	build	vcs.modified=false
```

But it doesn't
