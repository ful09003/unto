package internal

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCommandWrapper_Exec(t *testing.T) {
	type fields struct {
		trackedEnvVars map[string]string
		invokedTime    time.Time
		cEx            ExecutorWithOutErrPipes
		stdErrBytes    *bytes.Buffer
		stdOutBytes    *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				trackedEnvVars: nil,
				invokedTime:    time.Time{},
				cEx:            MockExecutor{},
				stdErrBytes:    new(bytes.Buffer),
				stdOutBytes:    new(bytes.Buffer),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandWrapper{
				trackedEnvVars: tt.fields.trackedEnvVars,
				invokedTime:    tt.fields.invokedTime,
				cEx:            tt.fields.cEx,
				stdErrBytes:    tt.fields.stdErrBytes,
				stdOutBytes:    tt.fields.stdOutBytes,
			}
			if err := c.Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommandWrapper_ToMarshalable(t *testing.T) {
	type fields struct {
		trackedEnvVars map[string]string
		invokedTime    time.Time
		cEx            ExecutorWithOutErrPipes
		stdErrBytes    *bytes.Buffer
		stdOutBytes    *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		want   Marshalable
	}{
		{
			name: "happy path",
			fields: fields{
				trackedEnvVars: map[string]string{"SHELL": "/bin/bash"},
				invokedTime:    time.Time{},
				cEx:            MockExecutor{},
				stdErrBytes:    new(bytes.Buffer),
				stdOutBytes:    new(bytes.Buffer),
			},
			want: Marshalable{
				Command:   "arg1",
				Args:      []string{"arg2"},
				Env:       map[string]string{"SHELL": "/bin/bash"},
				InvokedAt: time.Time{},
				Stdout:    "stdout",
				Stderr:    "stderr",
				ExitCode:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CommandWrapper{
				trackedEnvVars: tt.fields.trackedEnvVars,
				invokedTime:    tt.fields.invokedTime,
				cEx:            tt.fields.cEx,
				stdErrBytes:    tt.fields.stdErrBytes,
				stdOutBytes:    tt.fields.stdOutBytes,
			}
			_ = c.Exec()
			if got := c.ToMarshalable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMarshalable() = %v, want %v", got, tt.want)
			}
		})
	}
}
