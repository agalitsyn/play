bench:
	go test -v -bench=.

trace:
	go test -v -trace trace.out .
	go tool trace trace.out

pprof:
	go test -v -bench=. -benchmem -cpuprofile profile.out
	go tool pprof -http :8080 profile.out
