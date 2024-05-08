package main

import "testing"

func TestRunCmd(t *testing.T) {
	type args struct {
		cmd []string
		env Environment
	}
	nonEmptyMap := nonEmptyMap()

	tests := []struct {
		name           string
		args           args
		wantReturnCode int
	}{
		{
			name:           "Empty cmd",
			args:           args{cmd: make([]string, 0), env: nonEmptyMap},
			wantReturnCode: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReturnCode := RunCmd(tt.args.cmd, tt.args.env); gotReturnCode != tt.wantReturnCode {
				t.Errorf("RunCmd() = %v, want %v", gotReturnCode, tt.wantReturnCode)
			}
		})
	}
}
