package main

import (
	"bufio"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	dir = strings.ReplaceAll(dir, "$(pwd)/", "")
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}
		info, err := dirEntry.Info()
		if err != nil {
			return nil, err
		}
		envVar := strings.ReplaceAll(dirEntry.Name(), "=", "")

		env[envVar] = EnvValue{
			Value:      "",
			NeedRemove: true,
		}
		if info.Size() > 0 {
			file, err := os.Open(dir + "/" + dirEntry.Name())
			if err != nil {
				return nil, err
			}
			defer file.Close()

			content := ""
			scanner := bufio.NewScanner(file)
			if scanner.Scan() {
				content = strings.ReplaceAll(scanner.Text(), "\x00", "\n")
			}
			if content != "" {
				env[envVar] = EnvValue{
					Value:      content,
					NeedRemove: false,
				}
			}
		}
	}
	return env, nil
}
