package main

import (
	"flag"
	"log"
	"os"

	"github.com/Reinurmet/recontoolkit/internal/core"
	"github.com/Reinurmet/recontoolkit/internal/enum"
)

// If i wanna run use: go run ./cmd/subdomain -domain google.com -wordlistName testdata/subdomains.txt
func main() {

	// Define the flags
	domain := flag.String("domain", "string", "Name of the domain you want to check")
	wordlistName := flag.String("wordlistName", "string", "Input the wordlist")
	// Parse the command-line input
	flag.Parse()

	wordlist, err := os.Open(*wordlistName)
	if err != nil {
		log.Fatal(err)
	}

	hosts := enum.BasicLookup(*domain, wordlist)
	// _ skipps the index but keeps the host value
	for _, host := range hosts {
		core.PlainWriter(os.Stdout, host)
	}
	defer wordlist.Close()
}
