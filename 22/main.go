package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type LongStruct struct {
	Foo string
	Bar string
	Baz string
}

type BarBaz struct {
	Msg        string
	secretData string
}

func (l *LongStruct) Serialize() ([]byte, error) {
	b := BarBaz{Msg: fmt.Sprintf("%s - %s", l.Bar, l.Baz), secretData: "hide from serialization"}

	res, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	l := LongStruct{Foo: "foo", Bar: "bar", Baz: "baz"}

	res, err := l.Serialize()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(res))
}
