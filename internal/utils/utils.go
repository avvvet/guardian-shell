package utils

import "os/exec"

func ExecuteCommands(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
