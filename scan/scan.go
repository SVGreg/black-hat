package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fast := contains("-f", args)

	if !(len(args) == 2 || len(args) == 3 && fast) {
		fmt.Println("Usage: scan <ip address> or scan -f <ip address>")
		os.Exit(0)
	}

	host := func(fast bool, args []string) string {
		if fast {
			return args[2]
		} else {
			return args[1]
		}
	}(fast, args)

	ports := []int{22, 80, 445, 3128}

	if fast {
		scanFast(host, ports)
	} else {
		scanBallanced(host, ports, 10)
	}
}

func contains(value string, slice []string) bool {
	for _, element := range slice {
		if value == element {
			return true
		} else {
			continue
		}
	}
	return false
}
