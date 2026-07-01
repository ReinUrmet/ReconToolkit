package enum

import (
	"bufio"
	"io"
	"net"
	"sync"

	"github.com/Reinurmet/recontoolkit/internal/core"
)

func BasicLookup(domain string, wordlist io.Reader) []core.Host {
	buf := bufio.NewScanner(wordlist)
	hosts := []core.Host{}

	wg := &sync.WaitGroup{}
	var mu sync.Mutex

	for buf.Scan() {
		prefix := buf.Text()
		wg.Add(1)
		go func(prefix string) {
			defer wg.Done()
			subdomain := prefix + "." + domain
			address, err := net.LookupHost(subdomain)
			if err != nil {
				return
			}
			result := core.Host{Name: subdomain, IPs: address}
			mu.Lock()
			hosts = append(hosts, result)
			mu.Unlock()
		}(prefix)
	}

	wg.Wait()
	return hosts
}
