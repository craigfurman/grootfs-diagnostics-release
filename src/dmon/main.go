package main

import (
	"fmt"
	"os"
	"time"

	"dmon/dmon"

	"code.cloudfoundry.org/lager"
)

func main() {
	logger := lager.NewLogger("dmon")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	eventEmitter := &dmon.LoggingEventEmitter{Logger: logger}
	processManager := &dmon.LinuxProcessManager{}
	d := &dmon.Dmon{EventEmitter: eventEmitter, ProcessManager: processManager}

	if len(os.Args) != 2 {
		fmt.Println("usage: dmon <directory to check>")
		os.Exit(1)
	}

	if err := d.CheckFilesystemAvailability(logger, os.Args[1], time.Second*10); err != nil {
		os.Exit(1)
	}
}
