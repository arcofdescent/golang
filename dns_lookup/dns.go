package main

import "fmt"
import "net"
import "strings"

//import "errors"

func main() {
	fmt.Println("dns lookup")

	host := "yahoo.com"
	addrs, err := ResolveHost(host)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Host: %s\n", host)
		for _, ip := range addrs {
			fmt.Printf("IP : %s\n", ip)
		}
	}
}

func ResolveHost(h string) ([]string, error) {
	ret := make([]string, 0)
	addrs, err := net.LookupHost(h)
	if err != nil {
		return ret, err
	} else {
		for i := 0; i < len(addrs); i++ {
			segs := strings.SplitAfter(addrs[i], " ")
			ret = append(ret, segs...)
		}
		return ret, nil
	}
}
