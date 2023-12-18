package utils

import (
	"os/exec"
	"strings"
	"syscall"
)

// func ExecuteCommands(command string) (string, error) {
// 	cmd := exec.Command("/bin/sh", "-c", command)

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(output), nil
// }

func ExecuteCommands(command string) (string, error) {
	cmdParts := strings.Fields(command)
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)

	// Setup a pipe to pass the sudo password
	password := "awet_awet"
	cmd.Stdin = strings.NewReader(password + "\n")

	// For handling interactive sudo password prompts
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
