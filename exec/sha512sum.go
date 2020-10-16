// +build !nosha512sum
package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
    cmd.Register(&SHA512Sum)
}

var SHA512Sum = cmd.Sub {
    Name: "sha512sum",
    Short: "compute and check SHA512 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA512SumRun,
}

func SHA512SumRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA512, crypto.SHA512.Size()*8)
}
