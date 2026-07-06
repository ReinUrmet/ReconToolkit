package core

import "strings"

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

// This function is for when
func IPInScope(hostIp string, scope []string) bool {
	return true
}
