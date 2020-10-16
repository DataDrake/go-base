// +build !nobase32
package exec

import (
	"encoding/base32"
	"github.com/DataDrake/go-base/cmd"
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	cmd.Register(&Base32)
}

var Base32 = cmd.CMD {
	Name: "base32",
	Short: "base32 encode/decode data and print to standard output",
	Flags: &Base32Flags{},
	Args: &Base32Args{},
	Run: Base32Run,
}

type Base32Flags struct {
	Decode bool `short:"d" long:"decode" desc:"decode instead of encode (default)"`
}

type Base32Args struct {
	File string `desc:"File to encode or decode"`
}

func Base32Run(r *cmd.Root, c *cmd.CMD) {
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
