// +build !nobase64
package exec

import (
	"encoding/base64"
	"github.com/DataDrake/cli-ng/cmd"
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	cmd.Register(&Base64)
}

var Base64 = cmd.Sub {
	Name: "base64",
	Short: "base64 encode/decode data and print to standard output",
	Flags: &Base64Flags{},
	Args: &Base64Args{},
	Run: Base64Run,
}

type Base64Flags struct {
	Decode bool `short:"d" long:"decode" desc:"decode instead of encode (default)"`
}

type Base64Args struct {
	File string `desc:"File to encode or decode"`
}

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
