package core

import (
	"encoding/json"
	"fmt"
	"io"
)

// This is where the function is that turns a struct into output
func PlainWriter(writer io.Writer, host Host) {
	fmt.Fprintln(writer, host.Name)
}

// for front end use later
func NdjsonWriter(writer io.Writer, host Host) {
	encoder := json.NewEncoder(writer)
	jsonResult := encoder.Encode(host)
	fmt.Fprintln(writer, jsonResult)
}
