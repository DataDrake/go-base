// +build !nob2sum
package exec

import (
	"crypto"
	"github.com/DataDrake/go-base/cmd"
)

func init() {
    cmd.Register(&B2Sum)
}

var B2Sum = cmd.CMD {
    Name: "b2sum",
    Short: "compute and check BLAKE2 message digest",
    Flags: &B2SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: B2SumRun,
}

type B2SumFlags struct {
    Binary bool `short:"b" long:"binary" desc:"read in binary mode (default)"`
    Check  bool `short:"c" long:"check"  desc:"read sums from the FILEs and check them"`
    Tag    bool `short:"T" long:"tag"    desc:"create a BSD-style checksum"`
    Text   bool `short:"t" long:"text"   desc:"read in text mode (ignored)"`
    Zero   bool `short:"z" long:"zero"   desc:"end each output line with NUL, not newline"`

	Length int64 `short:"l" long:"length" desc: "digest length in bits; must not exceed the maximum for the blake2 algorithm and must be a multiple of 8"`
}

func (fs *B2SumFlags) convert() *SumFlags {
	return &SumFlags {
		Binary: fs.Binary,
		Check:  fs.Check,
		Tag:    fs.Tag,
		Text:   fs.Text,
		Zero:   fs.Zero,
	}
}

func B2SumRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*B2SumFlags)
    args := c.Args.(*SumArgs)

	length := crypto.BLAKE2b_512.Size()*8
	if flags.Length > 0 {
		length = int(flags.Length)
	}
	sumRun(flags.convert(), args, crypto.BLAKE2b_512, length)
}
