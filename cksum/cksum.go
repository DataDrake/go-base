package main

import (
	"flag"
	"fmt"
	"hash/adler32"
	"io"
	"os"
)

func usage() {
	fmt.Println("Usage: cksum [FILE]")
	flag.PrintDefaults()
}

func main() {
	var err error
	filename := ""
	flag.Usage = func() { usage() }
	flag.Parse()

	args := flag.Args()

	input := os.Stdin
	switch len(args) {
	case 0:
	case 1:
		if args[0] != "-" {
			input, err = os.Open(args[0])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			filename = args[0]
		}
	default:
		usage()
		os.Exit(1)
	}

	h := adler32.New()
	size, err := io.Copy(h, input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%d %d %s\n", h.Sum32(), size, filename)
	os.Exit(0)
}
