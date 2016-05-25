package main
import (
	"fmt"
	"flag"
)

func usage(){
	fmt.Println("Usage: yes [String]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = func() {usage()}
	flag.Parse()

	args := flag.Args()

	s := "y"
	switch len(args) {
	case 0:
	case 1:
		s = args[0]
	default:
		usage()
	}

	for {
		fmt.Println(s)
	}
}