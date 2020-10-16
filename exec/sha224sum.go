// +build !nosha224sum
package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
    cmd.Register(&SHA224Sum)
}

var SHA224Sum = cmd.Sub {
    Name: "sha224sum",
    Short: "compute and check SHA224 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA224SumRun,
}

func SHA224SumRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA224, crypto.SHA224.Size()*8)
}
