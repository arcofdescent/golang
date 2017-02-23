package main

import "fmt"
import "net"
import "strings"

//import "errors"

func main() {
	fmt.Println("dns lookup")

	host := "yahoo.com"
	addrs, err := DNSHost(host)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("Host: %s\n", host)
		for ip, idx := range addrs {
			fmt.Printf("IP %s: %s\n", (idx + 1), ip)
		}
	}
}

func DNSHost(h string) ([]string, error) {
	ret := make([]string, 0)
	addrs, err := net.LookupHost(h)
	if err != nil {
		return ret, err
	} else {
		for i := 0; i < len(addrs); i++ {
			segs := strings.SplitAfter(addrs[i], " ")
			ret = append(ret, segs)
		}
		return ret, nil
	}
}
