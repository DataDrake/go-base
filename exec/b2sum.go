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

// +build !nob2sum

package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
	cmd.Register(&B2Sum)
}

// B2Sum implements the "b2sum" subcommand
var B2Sum = cmd.Sub{
	Name:  "b2sum",
	Short: "compute and check BLAKE2 message digest",
	Flags: &B2SumFlags{
		Text: true,
	},
	Args: &SumArgs{},
	Run:  B2SumRun,
}

// B2SumFlags are flags unique to the "b2sum" subcommand
type B2SumFlags struct {
	Binary bool  `short:"b" long:"binary" desc:"read in binary mode (default)"`
	Check  bool  `short:"c" long:"check"  desc:"read sums from the FILEs and check them"`
	Tag    bool  `short:"T" long:"tag"    desc:"create a BSD-style checksum"`
	Text   bool  `short:"t" long:"text"   desc:"read in text mode (ignored)"`
	Zero   bool  `short:"z" long:"zero"   desc:"end each output line with NUL, not newline"`
	Length int64 `short:"l" long:"length" desc:"digest length in bits; must not exceed the maximum for the blake2 algorithm and must be a multiple of 8"`
}

func (fs *B2SumFlags) convert() *SumFlags {
	return &SumFlags{
		Binary: fs.Binary,
		Check:  fs.Check,
		Tag:    fs.Tag,
		Text:   fs.Text,
		Zero:   fs.Zero,
	}
}

// B2SumRun carries out the "b2sum" subcommand
func B2SumRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*B2SumFlags)
	args := c.Args.(*SumArgs)
	length := crypto.BLAKE2b_512.Size() * 8
	if flags.Length > 0 {
		length = int(flags.Length)
	}
	sumRun(flags.convert(), args, crypto.BLAKE2b_512, length)
}
