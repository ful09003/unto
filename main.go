package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/ful09003/unto/internal"
)

var (
	flagExportableEnvVars = flag.String("exportEnvVars", "SHELL,PWD,USER,TERM", "comma-separated list of env vars to export in this recording")
	flagCommandTimeout    = flag.Duration("timeout", 1*time.Minute, "duration allowable for command to complete")
)

func main() {
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
		fmt.Println(err)
	}

	out, err := yaml.Marshal(cmdStruct.ToMarshalable())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
