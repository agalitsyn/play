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

	fmt.Println("manual")
	fmt.Println(commitSHA)

	fmt.Println("manual and runtime")
	fmt.Println(Version())
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

var (
	Tag      string
	Revision string
	BuildAt  string
	Dirty    bool
)

func init() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	for _, setting := range buildInfo.Settings {
		// https://pkg.go.dev/runtime/debug#BuildSetting
		switch setting.Key {
		case "vcs.revision":
			Revision = setting.Value
		case "vcs.time":
			BuildAt = setting.Value
		case "vcs.modified":
			if setting.Value == "true" {
				Dirty = true
			}
		}
	}
}

func Version() string {
	if Revision == "" {
		return "unknown"
	}

	s := fmt.Sprintf("%s %s at %s", Tag, Revision, BuildAt)
	if Dirty {
		s += " dirty"
	}
	return s
}
