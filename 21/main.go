package main

import "fmt"

type MyStruct struct {
	Foo, Bar string
}

func NewMyStruct() *MyStruct {
	return &MyStruct{Foo: "foo", Bar: "bar"}
}

func main() {
	mystr := NewMyStruct()
	fmt.Printf("mystr = %+v\n", mystr)
	mystr.Foo = "Not Foo"
	fmt.Printf("mystr = %+v\n", mystr)
}
