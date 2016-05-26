package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func usage() {
	fmt.Println("Usage: tty [OPTION]...")
	flag.PrintDefaults()
}

func findRdev(dev_id uint64, path string) (string, error) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	files, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var dev_stat syscall.Stat_t
	for _, f := range files {
		err := syscall.Stat(filepath.Join(path, f), &dev_stat)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		if dev_id == dev_stat.Dev {
			return filepath.Join(path, f), nil
		}
	}
	return "", nil
}

func main() {
	flag.Usage = func() { usage() }
	silent := flag.Bool("s", false, "print nothing, only return an exit status")
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
		os.Exit(1)
	}
	var term_stat syscall.Stat_t
	err := syscall.Fstat(int(os.Stdin.Fd()), &term_stat)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dev, err := findRdev(term_stat.Dev, "/dev/pts")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(dev) == 0 {
		dev, err = findRdev(term_stat.Dev, "/dev")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if len(dev) == 0 {
			os.Exit(1)
		}
	}
	if !*silent {
		fmt.Println(dev)
	}
	os.Exit(0)
}
