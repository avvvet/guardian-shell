package utils

import (
	"os/exec"
	"strings"
)

func ExecuteCommands(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// func ExecuteCommands(command string) (string, error) {
// 	// Check if the command is interactive
// 	interactive := isInteractive(command)

// 	// Split the command into parts
// 	cmdParts := strings.Fields(command)

// 	// Use sudo to execute the command if interactive
// 	if interactive {
// 		cmd := exec.Command("sudo", append([]string{"-S"}, cmdParts...)...)
// 		// Setup a pipe to pass the sudo password
// 		password := "your_password"
// 		cmd.Stdin = strings.NewReader(password + "\n")

// 		// Run the command
// 		output, err := cmd.CombinedOutput()
// 		if err != nil {
// 			return "", err
// 		}
// 		return string(output), nil
// 	}

// 	// For non-interactive commands, run them directly
// 	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(output), nil
// }

func isInteractive(command string) bool {
	// Check for keywords or characters that may suggest interactivity
	interactiveKeywords := []string{"sudo", "su", "passwd", "ssh", "psql", "mysql", "bash"}
	for _, keyword := range interactiveKeywords {
		if strings.Contains(command, keyword) {
			return true
		}
	}
	return false
}
