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

// +build !nosha224sum

package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
	cmd.Register(&SHA224Sum)
}

// SHA224Sum implements the "sha224sum" subcommand
var SHA224Sum = cmd.Sub{
	Name:  "sha224sum",
	Short: "compute and check SHA224 sums",
	Flags: &SumFlags{
		Text: true,
	},
	Args: &SumArgs{},
	Run:  SHA224SumRun,
}

// SHA224SumRun carries out the "sha224sum" subcommand
func SHA224SumRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*SumFlags)
	args := c.Args.(*SumArgs)
	sumRun(flags, args, crypto.SHA224, crypto.SHA224.Size()*8)
}
