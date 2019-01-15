package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	net.DefaultResolver.PreferGo = false
	targetHostName := os.Args[1]

	r := net.Resolver{
		PreferGo: true,
		Dial:     Route53DNSDialer,
	}
	ctx := context.Background()

	// fmt.Println("DNS Result", ipaddr)
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
		addr, err := r.LookupHost(ctx, targetHostName)
		if err != nil {
			panic(err)
		}
		for _, v := range addr {
			fmt.Println("Resovle addr is ", v)
			if _, exits := result[v]; exits {
				result[v] = result[v] + 1
			} else {
				result[v] = 1
			}
		}
		fmt.Println("---")
		time.Sleep(250 * time.Millisecond)

	}
}
func Route53DNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "205.251.196.7:53")
}
