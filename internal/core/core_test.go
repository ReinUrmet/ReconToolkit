package core

import (
	"bytes"
	"strings"
	"testing"
)

// Use this for testing: go test ./internal/core -v
type scopeTestCase struct {
	name  string
	host  string
	scope []string
	want  bool
}

type IpScopeTestCase struct {
	name   string
	hostIp string
	scope  []string
	want   bool
}

func TestNdjsonWriter(t *testing.T) {

	bb := &bytes.Buffer{}

	var host1 Host
	host1.Name = "Testdomain"
	host1.IPs = []string{"172.198.123.255"}

	err := NdjsonWriter(bb, host1)
	if err != nil {
		t.Fatalf("NdjsonWriter returned error: %v", err)
	}

	output := bb.String()
	if !strings.Contains(output, "Testdomain") {
		t.Errorf("Doesn't contain Testdomain")
	}
	if strings.Contains(output, "<nil>") {
		t.Errorf("Contains <nil>")
	}
}

func TestPlainWriter(t *testing.T) {
	bb := &bytes.Buffer{}
	var host1 Host
	host1.Name = "Testdomain"
	host1.IPs = []string{"172.198.123.255"}
	PlainWriter(bb, host1)
	output := bb.String()
	if !strings.Contains(output, "Testdomain") {
		t.Errorf("Doesn't contain Testdomain")
	}

}

func TestInScope(t *testing.T) {
	tests := []scopeTestCase{
		{name: "exact match", host: "example.com", scope: []string{"example.com"}, want: true},
		{name: "no match", host: "example.com", scope: []string{"other.com"}, want: false},
		{name: "wildcard match", host: "api.example.com", scope: []string{"*.example.com"}, want: true},
		{name: "wildcard substring trap", host: "evilexample.com", scope: []string{"*.example.com"}, want: false},
		{name: "trailing dot", host: "example.com.", scope: []string{"example.com"}, want: true},
		{name: "case insensitivity", host: "EXAMPLE.com.", scope: []string{"example.com"}, want: true},
	}

	for _, tt := range tests {
		result := InScope(tt.host, tt.scope)
		if result != tt.want {
			t.Errorf("InScope(%q, %v) = %v, want %v", tt.host, tt.scope, result, tt.want)
		}

	}
}

func TestIPInScope(t *testing.T) {
	tests := []IpScopeTestCase{
		{name: "single IP exact match", hostIp: "255.255.255.255", scope: []string{"255.255.255.255/32"}, want: true},
		{name: "single IP no match", hostIp: "196.255.255.255", scope: []string{"255.255.255.255/32"}, want: false},
		{name: "in CIDR range", hostIp: "10.0.0.15", scope: []string{"10.0.0.0/24"}, want: true},
		{name: "outside CIDR range", hostIp: "10.0.1.15", scope: []string{"10.0.0.0/24"}, want: false},
		{name: "invalid IP input", hostIp: "not-an-ip", scope: []string{"10.0.0.0/24"}, want: false},
		{name: "malformed CIDR entry skipped", hostIp: "10.0.0.15", scope: []string{"garbage-cidr", "10.0.0.0/24"}, want: true},
	}
	for _, tt := range tests {
		result := IPInScope(tt.hostIp, tt.scope)
		if result != tt.want {
			t.Errorf("IPInScope(%q, %v) = %v, want %v", tt.hostIp, tt.scope, result, tt.want)
		}
	}
}
