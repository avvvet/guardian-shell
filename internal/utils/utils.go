package utils

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
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

func PrintWelcomeMessage() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintln(`
**************************

Welcome to Guardian Shell!

Your friendly tool for remote command execution.
Effortlessly manage and troubleshoot systems from remote with this reverse shell tool.
Use it for ethical tasks and Embrace hassle-free remote control.
Happy exploring!

Author Awet Tsegazeab

Current server date and time:`, currentTime, `

Enter your command to run it on remote shell

**************************`)
}
