package utils

import (
	"flag"
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

func GetPortFromArgs() int {
	portPtr := flag.Int("port", 20090, "-port The port number to connect to guardian shell")
	flag.Parse()
	return *portPtr
}

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
