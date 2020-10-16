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

// +build !nobase32

package exec

import (
	"encoding/base32"
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"io"
	"log"
	"os"
)

func init() {
	cmd.Register(&Base32)
}

// Base32 implements the "base32" subcommand
var Base32 = cmd.Sub{
	Name:  "base32",
	Short: "base32 encode/decode data and print to standard output",
	Flags: &Base32Flags{},
	Args:  &Base32Args{},
	Run:   Base32Run,
}

// Base32Flags are flags unique to the "base32" subcommand
type Base32Flags struct {
	Decode bool `short:"d" long:"decode" desc:"decode instead of encode (default)"`
}

// Base32Args are args unique to the "base32" subcommand
type Base32Args struct {
	File string `desc:"File to encode or decode"`
}

// Base32Run carries out the "base32" subcommand
func Base32Run(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*Base32Flags)
	args := c.Args.(*Base32Args)
	input, err := os.Open(args.File)
	if err != nil {
		log.Fatalf("Failed to open file '%s', reason: '%s'\n", args.File, err)
	}
	defer input.Close()
	if flags.Decode {
		dec := base32.NewDecoder(base32.StdEncoding, input)
		_, err = io.Copy(os.Stdout, dec)
	} else {
		enc := base32.NewEncoder(base32.StdEncoding, os.Stdout)
		_, err = io.Copy(enc, input)
		fmt.Println()
	}
	if err != nil {
		log.Fatalf("Operation failed, reason: %s\n", err.Error())
	}
	os.Exit(0)
}
