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
		wg.Add(1)
		go func() {
			defer wg.Done()
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
	wg.Wait()
}
