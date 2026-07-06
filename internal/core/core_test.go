package core

import (
	"bytes"
	"strings"
	"testing"
)

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
