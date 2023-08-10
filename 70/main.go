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
	fmt.Printf("Settings: %+v\n", buildInfo.Settings)
	fmt.Printf("Commit: %+v\n\n", Commit)

	fmt.Println("manual")
	fmt.Println(commitSHA)
}

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}()
