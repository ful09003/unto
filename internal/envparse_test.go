package internal

import (
	"reflect"
	"testing"
)

func Test_gatherVarVal(t *testing.T) {
	type args struct {
		inEnv []string
		k     string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "happy path",
			args: args{
				inEnv: []string{"SHELL=/bin/bash"},
				k:     "SHELL",
			},
			want:  "SHELL",
			want1: "/bin/bash",
		},
		{
			name: "happy path not found",
			args: args{
				inEnv: []string{"SHELL=/bin/bash"},
				k:     "USER",
			},
			want:  "USER",
			want1: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := gatherVarVal(tt.args.inEnv, tt.args.k)
			if got != tt.want {
				t.Errorf("gatherVarVal() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("gatherVarVal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGatherEnvVarVals(t *testing.T) {
	type args struct {
		inEnv []string
		k     []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "happy path",
			args: args{
				inEnv: []string{"SHELL=/bin/bash", "USER=me"},
				k:     []string{"SHELL", "USER"},
			},
			want: map[string]string{"SHELL": "/bin/bash", "USER": "me"},
		},
		{
			name: "happy path nothing found",
			args: args{
				inEnv: []string{"SHELL=/bin/bash", "USER=me"},
				k:     []string{"TERM"},
			},
			want: map[string]string{"TERM": ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GatherEnvVarVals(tt.args.inEnv, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GatherEnvVarVals() = %v, want %v", got, tt.want)
			}
		})
	}
}
