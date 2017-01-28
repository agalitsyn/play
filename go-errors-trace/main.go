package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/trace"
)

func IRaiseAnErrorToo() error {
	file, err := os.Open("/root/.bashrc")
	if err != nil {
		return trace.Wrap(err)
	}
	defer file.Close()

	return nil
}

func IRaiseAnError() error {
	if err := IRaiseAnErrorToo(); err != nil {
		return trace.Wrap(err, "I failed, but here is text which masked original error text with custom message")
	}
	return nil
}

func InitLoggerDebug() {
	// clear existing hooks:
	log.StandardLogger().Hooks = make(log.LevelHooks)
	log.SetFormatter(&trace.TextFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func main() {
	fmt.Println("compare output:")

	InitLoggerDebug()
	if err := IRaiseAnError(); err != nil {
		fmt.Println("======= error with trace")
		log.Error(err)

		fmt.Println("======= debug report")
		log.Error(trace.DebugReport(err))

		os.Exit(1)
	}
}
