package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func ListenUser(conn *net.Conn) {
	reader := bufio.NewReader(*conn)

	for {
		// Get message
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Something wrong when accepting message!")
			break
		} else {
			fmt.Printf("Get Message: " + msg)
		}

		// Push message to client
		fmt.Fprintf(*conn, "Your Message: %s", msg)
	}
}

func Listening(ln *net.Listener, wg *sync.WaitGroup) {
	for {
		conn, err := (*ln).Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection!\n")
			break
		} else {
			fmt.Printf("Someone join the server!\n")
		}

		go ListenUser(&conn)
	}
	wg.Done()
}

func main() {
	// fmt.Println("I'm Server")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen!")
		os.Exit(1)
	}
	fmt.Printf("### Server Started ###\n")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go Listening(&ln, &wg)
	wg.Wait()
}
