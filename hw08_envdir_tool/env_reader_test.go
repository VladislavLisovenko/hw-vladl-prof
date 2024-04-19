package main

import (
	"reflect"
	"testing"
)

func nonEmptyMap() map[string]EnvValue {
	nonEmptyMap := make(map[string]EnvValue)
	nonEmptyMap["BAR"] = EnvValue{Value: "bar", NeedRemove: false}
	nonEmptyMap["EMPTY"] = EnvValue{Value: "", NeedRemove: true}
	nonEmptyMap["FOO"] = EnvValue{Value: "foo", NeedRemove: false}
	nonEmptyMap["HELLO"] = EnvValue{Value: `"hello"`, NeedRemove: false}
	nonEmptyMap["UNSET"] = EnvValue{Value: "", NeedRemove: true}
	return nonEmptyMap
}

func TestReadDir(t *testing.T) {
	type args struct {
		dir string
	}
	nonEmptyMap := nonEmptyMap()

	tests := []struct {
		name    string
		args    args
		want    Environment
		wantErr bool
	}{
		{
			name:    "Empty dir, ReadDir has to return Error",
			args:    args{dir: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Wrong dir, ReadDir has to return Error",
			args:    args{dir: "qwqw"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Correct and empty dir, ReadDir has to return empty map without error",
			args:    args{dir: "testemptydir"},
			want:    make(Environment),
			wantErr: false,
		},
		{
			name:    "Correct and non empty dir, ReadDir has to return non empty map without error",
			args:    args{dir: "testdata/env"},
			want:    nonEmptyMap,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
