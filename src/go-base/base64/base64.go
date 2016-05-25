package main
import (
	"os"
	"encoding/base64"
	"io"
	"log"
	"fmt"
	"flag"
)

func usage(){
	fmt.Println("Usage: base64 [OPTION]... [FILE]")
	flag.PrintDefaults()
}

func main() {
	var err error
	flag.Usage = func() {usage()}
	decode := flag.Bool("d",false,"decode data")
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
		}
	default:
		usage()
	}

	if *decode {
		dec := base64.NewDecoder(base64.StdEncoding, input)
		_, err = io.Copy(os.Stdout, dec)
	} else {
		enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
		_, err = io.Copy(enc, input)
		enc.Close()
		fmt.Println()
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
}