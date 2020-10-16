// +build !nocat
package exec

import (
	"bufio"
	"fmt"
	"github.com/DataDrake/go-base/cmd"
	"os"
	"strconv"
	"strings"
)

func init() {
    cmd.Register(&Cat)
}

var Cat = cmd.CMD {
    Name: "cat",
    Short: "concatenate files and print on the standard output",
    Flags: &CatFlags{},
    Args: &CatArgs{},
    Run: CatRun,
}

type CatFlags struct {
    All             bool `short:"A" long:"all"         desc:"equivalent to -v -E -T"`
    NumberSome      bool `short:"b"                    desc:"number nonempty output lines, overrides -n"`
	NonprintingEnds bool `short:"e"                    desc:"equivalent to -v -E"`
	Ends            bool `short:"E" long:"ends"        desc:"display $ at the end of each line"`
	Number          bool `short:"n" long:"number"      desc:"number all output lines"`
	Squeeze         bool `short:"s" long:"squeeze"     desc:"suppress repeated empty output lines"`
	NonprintingTabs bool `short:"t"                    desc:"equivalent to -v -T"`
	Tabs            bool `short:"T" long:"tabs"        desc:"display TAB characters as ^I"`
	Nonprinting     bool `short:"v" long:"nonprinting" desc:"use ^ and M- notation, except for LFD and TAB"`
}

type CatArgs struct {
    Files []string `desc:"Files to concatenate"`
}

var catReplaces = []string {
	"\000", "^@",
	"\001", "^A",
	"\002", "^B",
	"\003", "^C",
	"\004", "^D",
	"\005", "^E",
	"\006", "^F",
	"\007", "^G",
	"\010", "^H",
	"\013", "^K",
	"\014", "^L",
	"\015", "^M",
	"\016", "^N",
	"\017", "^O",
	"\020", "^P",
	"\021", "^Q",
	"\022", "^R",
	"\023", "^S",
	"\024", "^T",
	"\025", "^U",
	"\026", "^V",
	"\027", "^W",
	"\030", "^X",
	"\031", "^Y",
	"\032", "^Z",
	"\033", "^[",
	"\034", "^\\",
	"\035", "^]",
	"\036", "^^",
	"\027", "^_",
	"\177", "^?",
}

func CatRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*CatFlags)
    args := c.Args.(*CatArgs)
	if flags.All {
		flags.Nonprinting = true
		flags.Ends = true
		flags.Tabs = true
	}
	if flags.NonprintingEnds {
		flags.Nonprinting = true
		flags.Ends = true
	}
	if flags.NonprintingTabs {
		flags.Nonprinting = true
		flags.Tabs = true
	}
	if flags.NumberSome {
		flags.Number = false
	}
	prev_empty := false
	lineno := 1
	replacer := strings.NewReplacer(catReplaces...)
	for _, file := range args.Files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			var formatted string
			//add line number to line
			if flags.Number || (flags.NumberSome && (len(line) > 0)) {
				formatted = strconv.Itoa(lineno) + " "
			}
			//show invisible tab characters
			if flags.Tabs {
				line = strings.Replace(line, "\t", "^I", -1)
			}
			//resolve other non-printing characters
			if flags.Nonprinting {
				line = replacer.Replace(line)
			}
			//ignore squeeze lines
			if (flags.Squeeze && prev_empty && (len(line) == 0)) {
				lineno++
				continue
			}
			//add line to params
			formatted += line
			//add $ for line endings
			if flags.Ends {
				formatted += "$"
			}
			fmt.Println(formatted)
			// keep track of empty lines and line number
			prev_empty = len(line) == 0
			lineno++
		}
		f.Close()
	}
	os.Exit(0)
}
