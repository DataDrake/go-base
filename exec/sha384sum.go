// +build !nosha384sum
package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
    cmd.Register(&SHA384Sum)
}

var SHA384Sum = cmd.Sub {
    Name: "sha384sum",
    Short: "compute and check SHA384 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA384SumRun,
}

func SHA384SumRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA384, crypto.SHA384.Size()*8)
}
