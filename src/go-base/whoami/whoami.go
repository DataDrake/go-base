package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

func usage() {
	fmt.Println("Usage: whoami [OPTION]...")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() { usage() }
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(1)
	}

	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(u.Username)
	os.Exit(0)
}
