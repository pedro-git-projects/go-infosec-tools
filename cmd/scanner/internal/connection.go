package internal

import (
	"fmt"
	"net"
	"sort"
)

func Scanner(protocol, networkAddress string) {
	ports := make(chan int, 100)
	results := make(chan int)
	openPorts := make([]int, 0)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, protocol, networkAddress)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1204; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports, results chan int, protocol, networkAddress string) {
	for p := range ports {
		addr := fmt.Sprintf("%s:%d", networkAddress, p)
		fmt.Println(addr)
		conn, err := net.Dial(protocol, addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
