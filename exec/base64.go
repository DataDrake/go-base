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

// +build !nobase64

package exec

import (
	"encoding/base64"
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"io"
	"log"
	"os"
)

func init() {
	cmd.Register(&Base64)
}

// Base64 implements the "base64" subcommand
var Base64 = cmd.Sub{
	Name:  "base64",
	Short: "base64 encode/decode data and print to standard output",
	Flags: &Base64Flags{},
	Args:  &Base64Args{},
	Run:   Base64Run,
}

// Base64Flags are flags unique to the "base64" subcommand
type Base64Flags struct {
	Decode bool `short:"d" long:"decode" desc:"decode instead of encode (default)"`
}

// Base64Args are args unique to the "base32" subcommand
type Base64Args struct {
	File string `desc:"File to encode or decode"`
}

// Base64Run carries out the "base64" subcommand
func Base64Run(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*Base64Flags)
	args := c.Args.(*Base64Args)
	input, err := os.Open(args.File)
	if err != nil {
		log.Fatalf("Failed to open file '%s', reason: '%s'\n", args.File, err)
	}
	defer input.Close()
	if flags.Decode {
		dec := base64.NewDecoder(base64.StdEncoding, input)
		_, err = io.Copy(os.Stdout, dec)
	} else {
		enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
		_, err = io.Copy(enc, input)
		fmt.Println()
	}
	if err != nil {
		log.Fatalf("Operation failed, reason: %s\n", err.Error())
	}
	os.Exit(0)
}
