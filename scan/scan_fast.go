package main

import (
	"fmt"
	"net"
	"sync"
)

func scanFast(host string, ports []int) {
	var wg sync.WaitGroup

	fmt.Printf("Fast scanning %s, ports:\n", host)
	for _, port := range ports {
		wg.Add(1)

		go func(host string, port int) {
			defer wg.Done()

			_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err == nil {
				fmt.Printf("  - %d\topen\n", port)
			} else {
				fmt.Printf("  - %d\tclosed\n", port)
			}
		}(host, port)
	}
	wg.Wait()
}
