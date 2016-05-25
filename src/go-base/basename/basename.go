package main
import (
	"fmt"
	"flag"
	"path/filepath"
	"strings"
)

func usage(){
	fmt.Println("Usage: basename NAME [SUFFIX]")
	fmt.Println("Usage: basename [OPTIONS]... NAME...")
	flag.PrintDefaults()
}

func getBase(path string,suffix string, zero bool) {
	base := filepath.Base(path)
	if len(suffix) > 0 {
		base = strings.TrimSuffix(base,suffix)
	}
	if zero {
		fmt.Printf("%s\000",base)
	} else {
		fmt.Println(base)
	}
}

func main() {

	flag.Usage = func() {usage()}
	multiple := flag.Bool("a",false,"support multiple arguments and treat each as a NAME")
	suffix := flag.String("s","","remove a trailing SUFFIX; implies -a")
	zero := flag.Bool("z",false,"end each output line with NUL, not newline")
	flag.Parse()

	args :=flag.Args()

	if len(args) == 0 {
		usage()
		return
	}

	if *multiple || len(*suffix) > 0 {
		for _,v := range args {
			getBase(v,*suffix,*zero)
		}
	} else {
		getBase(args[0],*suffix,*zero)
	}
}