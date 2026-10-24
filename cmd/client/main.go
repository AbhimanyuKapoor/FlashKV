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
	fmt.Println("Type EXIT to quit.")
	fmt.Printf("\n")

	// Continuously read everything coming from server
	go func() {
		reader := bufio.NewReader(conn)

		for {
			fmt.Printf("flashkv> ")
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nDisconnected from server.")
				os.Exit(0)
			}

			fmt.Print("\r\033[2K") // Clear current line (ANSI terminals)
			fmt.Println(msg)
		}
	}()

	input := bufio.NewReader(os.Stdin)

	for {
		command, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		command = strings.TrimSpace(command)

		if strings.EqualFold(command, "EXIT") {
			return
		}

		_, err = conn.Write([]byte(command + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
