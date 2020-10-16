// +build !nosha512sum
package exec

import (
	"crypto"
	"github.com/DataDrake/go-base/cmd"
)

func init() {
    cmd.Register(&SHA512Sum)
}

var SHA512Sum = cmd.CMD {
    Name: "sha512sum",
    Short: "compute and check SHA512 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA512SumRun,
}

func SHA512SumRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA512, crypto.SHA512.Size()*8)
}
