package main

import (
	"flag"
	"fmt"
	"strings"
)

func usage() {
	fmt.Println("Usage: echo [OPTION]... [STRING]...")
	flag.PrintDefaults()
}

func main() {

	flag.Usage = func() { usage() }
	nonewlines := flag.Bool("n", false, "do not output the trailing newline")
	escape := flag.Bool("e", false, "enable interpretation of backslash escapes")
	noescape := flag.Bool("E", false, "disable interpretation of backslash escapes (default)")
	flag.Parse()

	args := flag.Args()

	for _, v := range args {
		if *escape && !*noescape {
			v = strings.Replace(v,"\\n","\n",-1)
			fmt.Printf("%s",v)
		} else {
			fmt.Printf("%s",v)
		}
		if !*nonewlines {
			println()
		}
	}

}
