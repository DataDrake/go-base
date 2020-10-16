// +build !nomd5sum
package exec

import (
	"crypto"
	"github.com/DataDrake/go-base/cmd"
)

func init() {
    cmd.Register(&MD5Sum)
}

var MD5Sum = cmd.CMD {
    Name: "md5sum",
    Short: "compute and check MD5 sums",
    Flags: &SumFlags{
		Text: true,
	},
    Args: &SumArgs{},
    Run: MD5SumRun,
}

func MD5SumRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*SumFlags)
    args := c.Args.(*SumArgs)

    sumRun(flags, args, crypto.MD5, crypto.MD5.Size()*8)
}
