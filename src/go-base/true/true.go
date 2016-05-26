package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: true [OPTION]...")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(1)
	}

	os.Exit(0)
}
