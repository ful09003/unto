package internal

import (
	"bytes"
	"io"
	"time"
)

type CommandWrapper struct {
	trackedEnvVars           map[string]string
	invokedTime              time.Time
	cEx                      ExecutorWithOutErrPipes
	stdErrBytes, stdOutBytes *bytes.Buffer
}

func NewCommandWrapper(c CommandExecutor) *CommandWrapper {
	return &CommandWrapper{
		cEx:            c,
		invokedTime:    time.Now(),
		trackedEnvVars: map[string]string{},
		stdErrBytes:    new(bytes.Buffer),
		stdOutBytes:    new(bytes.Buffer),
	}
}

func (c *CommandWrapper) WithTrackedEnvVars(m map[string]string) *CommandWrapper {
	c.trackedEnvVars = m
	return c
}

func (c *CommandWrapper) WithExecTime(t time.Time) *CommandWrapper {
	c.invokedTime = t
	return c
}

func (c *CommandWrapper) Exec() error {
	so, soErr := c.cEx.StdoutPipe()
	se, seErr := c.cEx.StderrPipe()
	if soErr != nil {
		return soErr
	}
	if seErr != nil {
		return seErr
	}

	startErr := c.cEx.Start()

	_, _ = io.Copy(c.stdErrBytes, se)
	_, _ = io.Copy(c.stdOutBytes, so)
	if startErr != nil {
		return startErr
	}
	if err := c.cEx.Wait(); err != nil {
		return err
	}

	return nil
}

type Marshalable struct {
	Command   string            `yaml:"command"`
	Args      []string          `yaml:"args"`
	Env       map[string]string `yaml:"env_vars"`
	InvokedAt time.Time         `yaml:"invoked_at"`
	Stdout    string            `yaml:"stdout"`
	Stderr    string            `yaml:"stderr"`
	ExitCode  int               `yaml:"rc"`
}

func (c *CommandWrapper) ToMarshalable() Marshalable {
	return Marshalable{
		Command:   c.cEx.GetArgs()[0],
		Args:      c.cEx.GetArgs()[1:],
		Env:       c.trackedEnvVars,
		InvokedAt: c.invokedTime,
		Stdout:    c.stdOutBytes.String(),
		Stderr:    c.stdErrBytes.String(),
		ExitCode:  c.cEx.GetProcessState().ExitCode(),
	}
}
