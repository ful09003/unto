package internal

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

type ExecutorWithOutErrPipes interface {
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Start() error
	Wait() error
	GetProcessState() *os.ProcessState
	GetArgs() []string
}

type CommandExecutor struct {
	c *exec.Cmd
}

func NewCommandExecutor(cmd *exec.Cmd) CommandExecutor {
	return CommandExecutor{c: cmd}
}

func (c CommandExecutor) StdoutPipe() (io.ReadCloser, error) {
	return c.c.StdoutPipe()
}

func (c CommandExecutor) StderrPipe() (io.ReadCloser, error) {
	return c.c.StderrPipe()
}

func (c CommandExecutor) Start() error {
	return c.c.Start()
}

func (c CommandExecutor) Wait() error {
	return c.c.Wait()
}

func (c CommandExecutor) GetProcessState() *os.ProcessState {
	return c.c.ProcessState
}

func (c CommandExecutor) GetArgs() []string {
	return c.c.Args
}

type MockExecutor struct{}

func (m MockExecutor) StdoutPipe() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader("stdout")), nil
}

func (m MockExecutor) StderrPipe() (io.ReadCloser, error) {
	return io.NopCloser(strings.NewReader("stderr")), nil
}

func (m MockExecutor) Start() error {
	return nil
}

func (m MockExecutor) Wait() error {
	return nil
}

func (m MockExecutor) GetProcessState() *os.ProcessState {
	return &os.ProcessState{}
}

func (m MockExecutor) GetArgs() []string {
	return []string{"arg1", "arg2"}
}
