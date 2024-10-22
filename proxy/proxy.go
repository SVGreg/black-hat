package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: proxy <destination> <port>")
		os.Exit(0)
	}

	source := "localhost"
	destination := args[1]
	port, _ := strconv.ParseInt(args[2], 10, 0)

	fmt.Printf("Proxying traffic from %s to %s on port %d\n", source, destination, port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Unable to bind port")
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Unable to accept connection")
		}

		go spawn(conn)
	}
}

func spawn(src net.Conn) {
	fmt.Printf("Connected in %s\n", src.RemoteAddr().String())
}
