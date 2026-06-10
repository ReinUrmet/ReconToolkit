package main

import (
	"flag"
	"log"
	"os"

	"github.com/Reinurmet/recontoolkit/internal/enum"
)

// If i wanna run use: go run ./cmd/subdomain -domain google.com -wordlist_name testdata/subdomains.txt
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

	enum.BasicLookup(*domain, wordlist)

	defer wordlist.Close()
}
