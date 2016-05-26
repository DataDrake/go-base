package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: yes [String]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	args := flag.Args()

	s := "y"
	switch len(args) {
	case 0:
	case 1:
		s = args[0]
	default:
		usage()
		os.Exit(1)
	}

	for {
		fmt.Println(s)
	}
}
