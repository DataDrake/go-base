package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: false [OPTION]...")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()
	os.Exit(1)
}
