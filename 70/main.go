package main

import (
	"fmt"
	"runtime/debug"
)

var commitSHA string

func main() {
	buildInfo, _ := debug.ReadBuildInfo()
	fmt.Println("runtime/debug")
	fmt.Printf("%+v\n", buildInfo)
	fmt.Printf("Main: %+v\n", buildInfo.Main)
	fmt.Printf("Settings: %+v\n\n", buildInfo.Settings)

	fmt.Println("manual")
	fmt.Println(commitSHA)
}
