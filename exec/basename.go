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

// +build !nobasename

package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	cmd.Register(&Basename)
}

// Basename implements the "basename" subcommand
var Basename = cmd.Sub{
	Name:  "basename",
	Short: "strip directory and suffix from filenames",
	Flags: &BasenameFlags{},
	Args:  &BasenameArgs{},
	Run:   BasenameRun,
}

// BasenameFlags are flags unique to the "basename" subcommand
type BasenameFlags struct {
	Multiple bool   `short:"a" long:"multiple"    desc:"allow multiple paths to be specified"`
	Suffix   string `short:"s" long:"trim-suffix" desc:"remove a trailing SUFFIX; implies -a"`
	Zero     bool   `short:"z" long:"null"        desc:"end each output line with NUL, not newline"`
}

// BasenameArgs are args unique to the "basename" subcommand
type BasenameArgs struct {
	Paths []string `desc:"Path(s) to compute basename"`
}

func getBase(path string, suffix string, zero bool) {
	base := filepath.Base(path)
	if len(suffix) > 0 {
		base = strings.TrimSuffix(base, suffix)
	}
	if zero {
		fmt.Printf("%s\000", base)
	} else {
		fmt.Println(base)
	}
}

// BasenameRun carries out the "basename" subcommand
func BasenameRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*BasenameFlags)
	args := c.Args.(*BasenameArgs)
	if len(args.Paths) == 0 {
		r.SubUsage(c)
		os.Exit(1)
	}
	if flags.Multiple || len(flags.Suffix) > 0 {
		for _, path := range args.Paths {
			getBase(path, flags.Suffix, flags.Zero)
		}
	} else {
		if len(args.Paths) > 1 {
			r.SubUsage(c)
			os.Exit(1)
		}
		getBase(args.Paths[0], "", flags.Zero)
	}
}
