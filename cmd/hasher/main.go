package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/f1zm0/hades/pkg/hashing"
)

func main() {
	flag.Usage = func() {
		helpMsg := []string{
			"Usage:",
			"hasher '<string> [<string>...]'",
			"",
		}
		fmt.Fprintf(os.Stderr, strings.Join(helpMsg, "\n"))
	}

	flag.Parse()

	// Check if 1+ cli args has been specified
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("\n")
	for _, s := range flag.Args() {
		fmt.Printf("%s => %d\n", s, hashing.XORHash([]byte(s)))
	}
}
