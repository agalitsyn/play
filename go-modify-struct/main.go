package main

import "fmt"

type Test struct {
	Foo, Bar string
}

func (t *Test) SetFoo(value string) {
	t.Foo = value
}

func main() {
	t := Test{Foo: "foo", Bar: "bar"}
	fmt.Printf("%v", t)
	t.SetFoo("asdf")
	fmt.Printf("%v", t)
}
