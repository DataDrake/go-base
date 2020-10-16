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

// +build !notrue

package exec

import (
	"github.com/DataDrake/cli-ng/cmd"
	"os"
)

func init() {
	cmd.Register(&True)
}

// True implements the "true" subcommand
var True = cmd.Sub{
	Name:  "true",
	Short: "do nothing, successfully",
	Flags: &TrueFlags{},
	Args:  &TrueArgs{},
	Run:   TrueRun,
}

// TrueFlags are flags unique to the "true" subcommand
type TrueFlags struct{}

// TrueArgs are args unique to the "true" subcommand
type TrueArgs struct{}

// TrueRun carries out the "true" subcommand
func TrueRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	// flags := c.Flags.(*TrueFlags)
	// args := c.Args.(*TrueArgs)
	os.Exit(0)
}
