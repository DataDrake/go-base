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

// +build !nosha384sum

package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
	cmd.Register(&SHA384Sum)
}

// SHA384Sum implements the "sha384sum" subcommand
var SHA384Sum = cmd.Sub{
	Name:  "sha384sum",
	Short: "compute and check SHA384 sums",
	Flags: &SumFlags{
		Text: true,
	},
	Args: &SumArgs{},
	Run:  SHA384SumRun,
}

// SHA384SumRun carries out the "sha384sum" subcommand
func SHA384SumRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*SumFlags)
	args := c.Args.(*SumArgs)

	sumRun(flags, args, crypto.SHA384, crypto.SHA384.Size()*8)
}
