package foo

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	expected := Do(1)
	fmt.Println(expected)
	if "run 1" != expected {
		t.Fail()
	}
}
