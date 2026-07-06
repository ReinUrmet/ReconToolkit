package core

import (
	"net"
	"strings"
)

/*
This function checks two things:
1. Exact match - Plain string equality(It's case insensitive)
2. Wildcard match - Starts with "*."
*/
func InScope(host string, scope []string) bool {

	for _, item := range scope {
		if strings.ToLower(item) == strings.ToLower(host) {
			return true
		} else if strings.HasPrefix(strings.ToLower(item), "*.") {
			baseDomain := strings.TrimPrefix(item, "*.")
			//Checks if host is same as the ending and id host's length - basedomain's length is .
			if strings.HasSuffix(strings.ToLower(host), strings.ToLower(baseDomain)) && len(host) > len(baseDomain) && host[len(host)-len(baseDomain)-1] == '.' {
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
