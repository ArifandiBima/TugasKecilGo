package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot connect to server!")
		os.Exit(1)
	} else {
		fmt.Printf("Connected to server!\n")
	}

	connReader := bufio.NewReader(conn)
	localReader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("Type your message: ")
		msg, err := localReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot read the message!")
		} else {
			fmt.Printf("The msg is send!\n")
		}

		conn.Write([]byte(msg)) // fmt.Fprintf(conn, msg)

		echo, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read!\n")
			os.Exit(1)
		} else {
			fmt.Println("The echo has been received!\n")
		}

		fmt.Println(echo)
	}
}
