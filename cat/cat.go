package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Println("Usage: cat [OPTION]... [FILE]...")
	flag.PrintDefaults()
}

func main() {

	flag.Usage = func() { usage() }
	show_all := flag.Bool("A", false, "equivalent to -v -E -T")
	number_non_empty := flag.Bool("b", false, "number nonempty output lines, overrides -n")
	show_nonprinting_ends := flag.Bool("e", false, "equivalent to -v -E")
	show_ends := flag.Bool("E", false, "display $ at the end of each line")
	number_all := flag.Bool("n", false, "number all output lines")
	squeeze := flag.Bool("s", false, "suppress repeated empty output lines")
	show_nonprinting_tabs := flag.Bool("t", false, "equivalent to -v -T")
	show_tabs := flag.Bool("T", false, "display TAB characters as ^I")
	show_nonprinting := flag.Bool("v", false, "use ^ and M- notation, except for LFD and TAB")
	flag.Parse()

	args := flag.Args()

	prev_empty := false
	lineno := 1
	for _, v := range args {
		f, err := os.Open(v)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()

			formatted := ""
			//add line number to line
			if *number_all || (*number_non_empty && (len(line) > 0)) {
				formatted = formatted + strconv.Itoa(lineno) + " "
			}

			//show invisible tab characters
			if *show_all || *show_tabs || *show_nonprinting_tabs {
				line = strings.Replace(line, "\t", "^I", -1)
			}

			//resolve other non-printing characters
			if *show_all || *show_nonprinting || *show_nonprinting_ends {
				//line = line
			}

			//ignore squeeze lines
			if !(*squeeze && prev_empty && (len(line) == 0)) {

				//add line to params
				formatted += line

				//add $ for line endings
				if *show_all || *show_ends || *show_nonprinting_ends {
					formatted += "$"
				}

				fmt.Println(formatted)

				prev_empty = len(line) == 0
				lineno++
			}

		}

		f.Close()
	}
	os.Exit(0)
}
