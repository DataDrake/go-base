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

// +build !noecho

package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"strings"
)

func init() {
	cmd.Register(&Echo)
}

// Echo implements the "echo" subcommand
var Echo = cmd.Sub{
	Name:  "echo",
	Short: "display a line of text",
	Flags: &EchoFlags{
		NoEscapes: true,
	},
	Args: &EchoArgs{},
	Run:  EchoRun,
}

// EchoFlags are flags unique to the "echo" subcommand
type EchoFlags struct {
	NoNewlines bool `short:"n" long:"no-newlines" desc:"do not output the trailing newline"`
	Escapes    bool `short:"e" long:"escapes"     desc:"enable interpretation of backslash escapes"`
	NoEscapes  bool `short:"E" long:"no-escapes"  desc:"disable interpretation of backslash escapes (default)"`
}

// EchoArgs are args unique to the "echo" subcommand
type EchoArgs struct {
	Strings []string `desc:"strings to print"`
}

// EchoRun carries out the "echo" subcommand
func EchoRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*EchoFlags)
	args := c.Args.(*EchoArgs)

	if flags.Escapes {
		fmt.Println("Escapes are not currently supported")
		os.Exit(1)
	}

	for _, v := range args.Strings {
		if flags.Escapes && !flags.NoEscapes {
			v = strings.Replace(v, "\\n", "\n", -1)
			fmt.Printf("%s", v)
		} else {
			fmt.Printf("%s", v)
		}
		if !flags.NoNewlines {
			println()
		}
	}
}
