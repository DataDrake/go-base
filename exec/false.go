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

// +build !nofalse

package exec

import (
	"github.com/DataDrake/cli-ng/cmd"
	"os"
)

func init() {
	cmd.Register(&False)
}

// False implements the "false" subcommand
var False = cmd.Sub{
	Name:  "false",
	Short: "do nothing, unsuccessfully",
	Flags: &FalseFlags{},
	Args:  &FalseArgs{},
	Run:   FalseRun,
}

// FalseFlags are flags unique to the "false" subcommand
type FalseFlags struct{}

// FalseArgs are args unique to the "false" subcommand
type FalseArgs struct{}

// FalseRun carries out the "false" subcommand
func FalseRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	//flags := c.Flags.(*FalseFlags)
	//args := c.Args.(*FalseArgs)
	os.Exit(1)
}
