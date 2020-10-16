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

// +build !noyes

package exec

import (
	"bufio"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"strings"
)

func init() {
	cmd.Register(&Yes)
}

// Yes implements the "yes" subcommand
var Yes = cmd.Sub{
	Name:  "yes",
	Short: "output a string repeatedly until killed",
	Flags: &YesFlags{},
	Args:  &YesArgs{},
	Run:   YesRun,
}

// YesFlags are flags unique to the "yes" subcommand
type YesFlags struct{}

// YesArgs are args unique to the "yes" subcommand
type YesArgs struct {
	Strings []string `desc:"String(s) to print out"`
}

// YesRun carries out the "yes" subcommand
func YesRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	// flags := c.Flags.(*YesFlags)
	args := c.Args.(*YesArgs)
	s := "y"
	if len(args.Strings) > 0 {
		s = strings.Join(args.Strings, " ")
	}
	s += "\n"
	var b []byte
	for i := 0; i < 1000; i++ {
		b = append(b, []byte(s)...)
	}
	buff := bufio.NewWriterSize(os.Stdout, 1000*len(s))
	for {
		buff.Write(b)
	}
}
