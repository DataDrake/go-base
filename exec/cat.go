//
// Copyright 2016-2020 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// +build !nocat

package exec

import (
	"bufio"
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"strconv"
	"strings"
)

func init() {
	cmd.Register(&Cat)
}

// Cat implements the "cat" subcommand
var Cat = cmd.Sub{
	Name:  "cat",
	Short: "concatenate files and print on the standard output",
	Flags: &CatFlags{},
	Args:  &CatArgs{},
	Run:   CatRun,
}

// CatFlags are flags unique to the "cat" subcommand
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

// CatArgs are args unique to the "cat" subcommand
type CatArgs struct {
	Files []string `desc:"Files to concatenate"`
}

// catReplaces map non-printable characters to their printable counterparts
var catReplaces = []string{
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

// Clean makes sure the flags are minimally set
func (f *CatFlags) Clean() {
	if f.All {
		f.Nonprinting = true
		f.Ends = true
		f.Tabs = true
	}
	if f.NonprintingEnds {
		f.Nonprinting = true
		f.Ends = true
	}
	if f.NonprintingTabs {
		f.Nonprinting = true
		f.Tabs = true
	}
	if f.NumberSome {
		f.Number = false
	}
}

// CatRun carries out the "cat" subcommand
func CatRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*CatFlags)
	flags.Clean()
	args := c.Args.(*CatArgs)
	prevEmpty := false
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
			//ignore squeeze lines
			if flags.Squeeze && prevEmpty && (line == "") {
				continue
			}
			//add line number to line
			if flags.Number || (flags.NumberSome && line != "") {
				formatted = strconv.Itoa(lineno) + " "
				lineno++
			}
			//show invisible tab characters
			if flags.Tabs {
				line = strings.Replace(line, "\t", "^I", -1)
			}
			//resolve other non-printing characters
			if flags.Nonprinting {
				line = replacer.Replace(line)
			}
			//add line to params
			formatted += line
			//add $ for line endings
			if flags.Ends {
				formatted += "$"
			}
			fmt.Println(formatted)
			// keep track of empty lines and line number
			prevEmpty = len(line) == 0
		}
		f.Close()
	}
	os.Exit(0)
}
