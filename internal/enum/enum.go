package enum

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// Peaksin kasutama LookupHost
func BasicLookup(domain string, wordlist io.Reader) {

	buf := bufio.NewScanner(wordlist)

	for buf.Scan() {
		go func() {
			prefix := buf.Text()
			subdomain := prefix + "." + domain
			address, err := net.LookupHost(subdomain)
			if err != nil {
				return
			} else {
				fmt.Println(subdomain)
				fmt.Println(address)
				fmt.Println()
			}
		}()
	}
}
