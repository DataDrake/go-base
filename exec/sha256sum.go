// +build !nosha256sum
package exec

import (
	"crypto"
	"github.com/DataDrake/cli-ng/cmd"
)

func init() {
    cmd.Register(&SHA256Sum)
}

var SHA256Sum = cmd.Sub {
    Name: "sha256sum",
    Short: "compute and check SHA256 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA256SumRun,
}

func SHA256SumRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA256, crypto.SHA256.Size()*8)
}
