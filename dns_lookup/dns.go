// Read a file containing host names, one per line
// Do a DNS lookup for each host
// Output a file per host, with their IP addresses
// The output file name should be the hostname
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
	//"errors"
)

func main() {
	fmt.Println("dns lookup")

	if len(os.Args) < 2 {
		log.Fatal("Specify file as argument")
	}

	hostFile := os.Args[1]
	hosts := getHosts(hostFile)

	startTime := time.Now()

	for _, h := range hosts {
		fmt.Printf("Resolving host %s...\n", h)

		err := ResolveHost(h)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
	}

	dur := time.Since(startTime)
	fmt.Printf("duration: %v\n", dur)
}

func getHosts(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(file)
	ret := make([]string, 0)

	for input.Scan() {
		line := input.Text()

		// should not be an empty line
		if len(line) > 0 {
			ret = append(ret, line)
		}
	}

	return ret
}

func ResolveHost(h string) error {
	ret := make([]string, 0)

	addrs, err := net.LookupHost(h)
	if err != nil {
		return err
	}

	for i := 0; i < len(addrs); i++ {
		segs := strings.SplitAfter(addrs[i], " ")
		ret = append(ret, segs...)
	}

	// write to file
	f, err := os.Create(h)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for _, ip := range addrs {
		f.WriteString(fmt.Sprintf("%s\n", ip))
	}

	f.Sync()
	return nil
}
