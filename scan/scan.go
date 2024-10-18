package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: scan <ip address>")
		os.Exit(0)
	}
	host := args[1]
	ports := []int{22, 80, 445, 3128}

	fast := false

	if fast {
		scanFast(host, ports)
	} else {
		scanBallanced(host, ports, 10)
	}
}
