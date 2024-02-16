package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ful09003/unto/internal"
)

var (
	flagExportableEnvVars = flag.String("exportEnvVars", "SHELL,PWD,USER,TERM", "comma-separated list of env vars to export in this recording")
	flagCommandTimeout    = flag.Duration("timeout", 1*time.Minute, "duration allowable for command to complete")
)

func main() {
	cwd, _ := os.Executable()
	flagSaveToDir := flag.String("savedir", cwd, "directory to save unto outputs into")

	origEnviron := os.Environ()
	flag.Parse()
	inFlags := flag.Args()

	ctx, cancel := context.WithTimeout(context.Background(), *flagCommandTimeout)
	defer cancel()

	cmdExecutor := internal.NewCommandExecutor(exec.CommandContext(ctx, inFlags[0], inFlags[1:]...))
	cmdStruct := internal.NewCommandWrapper(cmdExecutor)
	trackedEnvVars := internal.GatherEnvVarVals(origEnviron, strings.Split(*flagExportableEnvVars, ","))
	cmdStruct = cmdStruct.WithTrackedEnvVars(trackedEnvVars).
		WithExecTime(time.Now())

	if err := cmdStruct.Exec(); err != nil {
		log.Println(err)
	}

	if err := internal.PersistOutput(*flagSaveToDir, cmdStruct.ToMarshalable()); err != nil {
		log.Println(err)
	}
}
