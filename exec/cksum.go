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

// +build !nocksum

package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"hash/crc32"
	"io"
	"os"
)

func init() {
	cmd.Register(&CkSum)
}

// CkSum implements the "cksum" subcommand
var CkSum = cmd.Sub{
	Name:  "cksum",
	Short: "checksum and count the bytes in a file (IEEE CRC-32, not POSIX)",
	Args:  &CkSumArgs{},
	Run:   CkSumRun,
}

// CkSumArgs are args unique to the "cksum" subcommand
type CkSumArgs struct {
	Files []string `zero:"yes" desc:"File to encode or decode"`
}

// CkSumRun carries out the "cksum" subcommand
func CkSumRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	args := c.Args.(*CkSumArgs)
	if len(args.Files) == 0 {
		ckSum(os.Stdin, "")
		return
	}
	for _, f := range args.Files {
		input, err := os.Open(f)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		ckSum(input, f)
		input.Close()
	}
}

func ckSum(in io.Reader, name string) {
	h := crc32.NewIEEE()
	size, err := io.Copy(h, in)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%d %d %s\n", h.Sum32(), size, name)
}
