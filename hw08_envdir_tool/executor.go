package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 0
	}

	envVariables := []string{}
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
			continue
		}
		envVariables = append(envVariables, fmt.Sprintf("%s=%s ", k, v.Value))
	}

	//command := exec.Command("/bin/bash", "testdata/echo.sh", "arg1=1", "arg2=2")
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = envVariables
	command.Env = append(command.Env, os.Environ()...)

	output, err := command.CombinedOutput()
	if err != nil {
		return exec.ExitError{}.ExitCode()
	}
	fmt.Println(string(output))

	return 0
}
