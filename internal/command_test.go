package internal

import (
	"bytes"
	"os/exec"
	"testing"
	"time"
)

func TestCommandWrapper_Exec(t *testing.T) {
	type fields struct {
		trackedEnvVars map[string]string
		invokedTime    time.Time
		command        *exec.Cmd
		stdErrBytes    *bytes.Buffer
		stdOutBytes    *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandWrapper{
				trackedEnvVars: tt.fields.trackedEnvVars,
				invokedTime:    tt.fields.invokedTime,
				command:        tt.fields.command,
				stdErrBytes:    tt.fields.stdErrBytes,
				stdOutBytes:    tt.fields.stdOutBytes,
			}
			if err := c.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
