package internal

import "time"

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

type Marshalable struct {
	Command   string            `yaml:"command"`
	Args      []string          `yaml:"args"`
	Env       map[string]string `yaml:"env_vars"`
	InvokedAt time.Time         `yaml:"invoked_at"`
	Stdout    string            `yaml:"stdout"`
	Stderr    string            `yaml:"stderr"`
	ExitCode  int               `yaml:"rc"`
}
