package core

import (
	"net"
	"strings"

	"golang.org/x/net/idna"
)

/*
This function checks two things:
1. Exact match - Plain string equality(It's case insensitive)
2. Wildcard match - Starts with "*."
*/
func InScope(host string, scope []string) bool {

	host = strings.ToLower(host)
	//Trailing dot filtering
	host = strings.TrimSuffix(host, ".")
	//Punycode fix (go get golang.org/x/net/idna)
	host, err := idna.ToASCII(host)
	if err != nil {
		return false
	}

	for _, item := range scope {
		item = strings.ToLower(item)
		//Trailing dot filtering
		item = strings.TrimSuffix(item, ".")
		//Punycode fix for item aswell
		item, err := idna.ToASCII(item)
		if err != nil {
			continue
		}

		if item == host {
			return true
		} else if strings.HasPrefix(item, "*.") {
			baseDomain := strings.TrimPrefix(item, "*.")
			baseDomain = strings.ToLower(baseDomain)
			//Checks if host is same as the ending and id host's length - basedomain's length is .
			if strings.HasSuffix(host, baseDomain) && len(host) > len(baseDomain) && host[len(host)-len(baseDomain)-1] == '.' {
				return true
			}
		}
	}
	return false
}

// This function is for checking if a single IP address fall within a range of IPs (in this case, 10.0.0.0 through 10.0.0.255)?
func IPInScope(hostIp string, scope []string) bool {
	ip := net.ParseIP(hostIp)
	if ip == nil {
		return false
	}
	for _, entry := range scope {
		_, ipnet, err := net.ParseCIDR(entry)
		if err != nil {
			continue
		} else if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}
