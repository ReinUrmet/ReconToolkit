package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

// This is where the function is that turns a struct into output
func PlainWriter(writer io.Writer, host Host) {
	fmt.Fprintln(writer, host.Name)
}

//TODO
//NDJSON writer buggy (prints stray <nil>)

// for front end use later
func NdjsonWriter(writer io.Writer, host Host) {
	encoder := json.NewEncoder(writer)
	jsonResult := encoder.Encode(host)
	fmt.Fprintln(writer, jsonResult)
}

//TODO
//NDJSON parsing + stdin/-f wiring

// For reading lists (reqired for tools like: probe, dns, takeover)
func TargetReader(input io.Reader) []Host {
	buf := bufio.NewScanner(input)
	lines := []Host{}

	for buf.Scan() {
		line := buf.Text()
		lines = append(lines, Host{Name: line})
	}
	return lines
}
