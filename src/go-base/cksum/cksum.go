package main
import (
	"os"
	"io"
	"fmt"
	"flag"
	"hash/adler32"
)

func usage(){
	fmt.Println("Usage: cksum [FILE]")
	flag.PrintDefaults()
}

func main() {
	var err error
	filename := ""
	flag.Usage = func() {usage()}
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
		return
	}

	h := adler32.New()
	size, err := io.Copy(h, input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%d %d %s\n",h.Sum32(),size,filename)
}