// +build !nosha256sum
package exec

import (
	"crypto"
	"github.com/DataDrake/go-base/cmd"
)

func init() {
    cmd.Register(&SHA256Sum)
}

var SHA256Sum = cmd.CMD {
    Name: "sha256sum",
    Short: "compute and check SHA256 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA256SumRun,
}

func SHA256SumRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA256, crypto.SHA256.Size()*8)
}
