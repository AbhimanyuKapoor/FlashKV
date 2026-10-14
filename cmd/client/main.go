package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to FlashKV 🚀")
	fmt.Println("Type EXIT to quit.\n")

	serverReader := bufio.NewReader(conn)
	inputReader := bufio.NewReader(os.Stdin)

	for {

		// Read prompt
		prompt, err := serverReader.ReadString(' ')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(prompt)

		// Read user input
		command, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSpace(command)

		if strings.EqualFold(command, "EXIT") {
			fmt.Println("Goodbye 👋")
			return
		}

		_, err = conn.Write([]byte(command + "\n"))
		if err != nil {
			log.Fatal(err)
		}

		// Read server response
		response, err := serverReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(response)
	}
}
