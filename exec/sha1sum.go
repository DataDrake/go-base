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

// +build !nosha1sum

package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
	cmd.Register(&SHA1Sum)
}

// SHA1Sum implements the "sha1sum" subcommand
var SHA1Sum = cmd.Sub{
	Name:  "sha1sum",
	Short: "compute and check SHA1 sums",
	Flags: &SumFlags{
		Text: true,
	},
	Args: &SumArgs{},
	Run:  SHA1SumRun,
}

// SHA1SumRun carries out the "sha1sum" subcommand
func SHA1SumRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*SumFlags)
	args := c.Args.(*SumArgs)

	sumRun(flags, args, crypto.SHA1, crypto.SHA1.Size()*8)
}
