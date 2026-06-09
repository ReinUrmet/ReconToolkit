package enum

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Peaksin kasutama LookupHost
func BasicLookup(domain string) {

	subdomains, err := os.Open("testdata/subdomains.txt")
	if err != nil {
		log.Fatal(err)
	}
	buf := bufio.NewScanner(subdomains)

	for buf.Scan() {
		prefix := buf.Text()
		subdomain := prefix + "." + domain
		address, err := net.LookupHost(subdomain)
		fmt.Println(address)
		if err != nil {
			fmt.Println(err)
		}
	}
}
