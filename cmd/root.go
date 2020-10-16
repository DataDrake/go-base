//
// Copyright 2017-2020 Bryan T. Meyers <root@datadrake.com>
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

package cmd

import (
	"fmt"
	"github.com/DataDrake/go-base/options"
	"os"
	"path/filepath"
	"sort"
)

var subcommands map[string]*CMD

func Register(c *CMD) {
	if subcommands == nil {
		subcommands = make(map[string]*CMD)
	}
	subcommands[c.Name] = c
}

// Root is the main command that runs everything
type Root struct {
	Name        string
	Short       string
	Flags       interface{}
}

// Usage prints the usage for this program
func (r *Root) Usage() {
	// Print Usage
	fmt.Printf("USAGE: %s [OPTIONS] CMD\n\n", r.Name)
	// Print Description
	if len(r.Short) > 0 {
		fmt.Printf("DESCRIPTION: %s\n\n", r.Short)
	}
	// Print sub-commands
	fmt.Printf("COMMANDS:\n\n")
	// Key the names of the sub-commands and find the longest command and alias
	var keys []string
	maxKey := 0
	for key, cmd := range subcommands {
		if cmd.Hidden {
			continue
		}
		keys = append(keys, key)
		if len(key) > maxKey {
			maxKey = len(key)
		}
	}
	sort.Strings(keys)
	// Add spacing for ()
	format := fmt.Sprintf("    %%%ds : %%s\n", maxKey)
	for _, k := range keys {
		fmt.Printf(format, k, subcommands[k].Short)
	}
	print("\n")
	// Print the global flags
	if r.Flags != nil {
		fmt.Printf("GLOBAL FLAGS:\n\n")
		PrintFlags(r.Flags)
	}
	os.Exit(1)
}

// Run finds the appropriate CMD and executes it, or prints the global Usage
func (r *Root) Run() {
	p, sub := options.NewParser(os.Args)
	if sub == "" {
		r.Usage()
	}
	sub = filepath.Base(sub)
	// Get the subcommand if it exists
	c := subcommands[sub]
	if c == nil {
		r.Usage()
	}
	// Handle any flags for the RootCMD
	var extras int
	if r.Flags != nil {
		extras = p.SetFlags(r.Flags)
	}
	// Not yet supported
	if c.Flags != nil {
		extras += p.SetFlags(c.Flags)
	}
	// Check for unknown flags
	p.UnknownFlags()
	// Handle the arguments for the subcommand
	if !p.SetArgs(c.Args, extras) {
		Usage(r, c)
		os.Exit(1)
	}
	c.Run(r, c)
}
