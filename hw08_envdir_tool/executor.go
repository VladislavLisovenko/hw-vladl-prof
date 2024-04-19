package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 0
	}

	sb := strings.Builder{}
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
			continue
		}
		sb.WriteString(fmt.Sprintf("%s=%s ", k, v.Value))
	}

	commandText := fmt.Sprint(sb.String(), cmd[0])
	command := exec.Command(commandText, cmd[1:]...)

	if err := command.Start(); err != nil {
		return exec.ExitError{}.ExitCode()
	}
	if err := command.Wait(); err != nil {
		return exec.ExitError{}.ExitCode()
	}
	return 0
}
