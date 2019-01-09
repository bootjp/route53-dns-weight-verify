package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	targetHostName := os.Args[1]

	result := map[string]int{}
	defer func() {
		total := 0
		for _, v := range result {
			total += v
		}
		println(total)

		for k, v := range result {
			fmt.Println(k, v)
			fmt.Printf("%f\n", float64(v)/float64(total))
		}

	}()

	for index := 0; index < 1000; index++ {
		addr, err := net.LookupHost(targetHostName)
		if err != nil {
			fmt.Println("Resolve error ", err)
			os.Exit(1)
		}
		for _, v := range addr {
			fmt.Println("Resovle addr is ", v)
			if _, exits := result[v]; exits {
				result[v] = result[v] + 1
			} else {
				result[v] = 1
			}

		}
		time.Sleep(500 * time.Millisecond)

	}
}
