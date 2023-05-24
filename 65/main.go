package foo

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Do(i))
	}
}

func Do(i int) string {
	r := fmt.Sprintf("run %d", i)
	return r
}

type Foo struct {
	Name string
}
