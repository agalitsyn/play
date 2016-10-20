build-verbose:
	go build -x

build-show-asm:
	go build -gcflags="-S"

disasm:
	go tool objdump -s main.main go-debugging-code-generation

names-table:
	go tool nm go-debugging-code-generation

build-show-ssa:
	GOSSAFUNC=main go build && xdg-open ssa.html

build-show-optimizations:
	go build -gcflags="-m=2"

build-no-optimizations:
	go build -gcflags="-l -N"

build-show-lexer:
	go build -gcflags="-x"
