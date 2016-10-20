package main

import (
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
		return trace.Wrap(err, "custom message")
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
	InitLoggerDebug()

	if err := IRaiseAnError(); err != nil {
		log.Error(trace.DebugReport(err))
		os.Exit(1)
	}
}
