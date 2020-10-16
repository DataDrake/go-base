// +build !nosha1sum
package exec

import (
	"crypto"
	"github.com/DataDrake/go-base/cmd"
)

func init() {
    cmd.Register(&SHA1Sum)
}

var SHA1Sum = cmd.CMD {
    Name: "sha1sum",
    Short: "compute and check SHA1 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: SHA1SumRun,
}

func SHA1SumRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.SHA1, crypto.SHA1.Size()*8)
}
