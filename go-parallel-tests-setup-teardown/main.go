package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		name, err := createFile(tmpDir)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
}

func createFile(dir string) (string, error) {
	f, err := ioutil.TempFile(dir, "")
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
