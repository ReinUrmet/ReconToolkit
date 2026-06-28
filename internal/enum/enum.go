package enum

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

func BasicLookup(domain string, wordlist io.Reader) {

	buf := bufio.NewScanner(wordlist)
	wg := &sync.WaitGroup{}

	for buf.Scan() {
		prefix := buf.Text()
		wg.Add(1)
		go func(prefix string) {
			defer wg.Done()
			subdomain := prefix + "." + domain
			address, err := net.LookupHost(subdomain)
			if err != nil {
				return
			} else {
				fmt.Println(subdomain)
				fmt.Println(address)
				fmt.Println()
			}
		}(prefix)
	}
	wg.Wait()
}
