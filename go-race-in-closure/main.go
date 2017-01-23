package main

import (
	"errors"
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string
}

func main() {
	var items []*MyStruct
	fooFunc := func(items []*MyStruct) {
		for _, item := range items {
			fmt.Println(item.Name)
		}
	}
	funcs := map[string]interface{}{"foo": fooFunc}
	Call(funcs, "foo", items)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
