package main

import (
	"fmt"
	"io"
	"log"
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
	port, err := strconv.ParseInt(args[2], 10, 0)
	if err != nil {
		fmt.Println("Unable to parse port.")
		fmt.Println("Usage: proxy <destination> <port>")
		os.Exit(0)
	}

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

		go spawn(conn, fmt.Sprintf("%s:%d", destination, port))
	}
}

func spawn(src net.Conn, dstHost string) {
	fmt.Printf("Connected in %s\n", src.RemoteAddr().String())

	dst, err := net.Dial("tcp", dstHost)
	if err != nil {
		fmt.Printf("Unable to connect to %s\n", dstHost)
		os.Exit(0)
	}

	fmt.Printf("Connected to %s\n", dstHost)

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
