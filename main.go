package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/chnzjm/sample/matchers"
	"github.com/chnzjm/sample/search"
)

// init is called prior to main
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	if len(os.Args) > 1 {
		// Perform the search for the specified term.
		search.Run(os.Args[1])
	} else {
		fmt.Printf("Usage: Program searchTerm\n")
	}
}
