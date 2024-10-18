package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(host string, portschan, reschan chan int) {
	for port := range portschan {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			reschan <- 0
			continue
		}
		conn.Close()
		reschan <- port
	}
}

func scanBallanced(host string, ports []int, width int) {
	fmt.Printf("Scanning %s, ports:\n", host)

	portschan := make(chan int, width)
	reschan := make(chan int)
	var openports []int

	for i := 0; i < cap(portschan); i++ {
		go worker(host, portschan, reschan)
	}

	go func() {
		for _, port := range ports {
			portschan <- port
		}
	}()

	for i := 0; i < len(ports); i++ {
		port := <-reschan
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(portschan)
	close(reschan)

	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("  %d\topen\n", port)
	}
}
