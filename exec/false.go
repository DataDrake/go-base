// +build !nofalse
package exec

import (
	"github.com/DataDrake/go-base/cmd"
	"os"
)

func init() {
    cmd.Register(&False)
}

var False = cmd.CMD {
    Name: "false",
    Short: "do nothing, unsuccessfully",
    Flags: &FalseFlags{},
    Args: &FalseArgs{},
    Run: FalseRun,
}

type FalseFlags struct {}

type FalseArgs struct {}

func FalseRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    //flags := c.Flags.(*FalseFlags)
    //args := c.Args.(*FalseArgs)
	os.Exit(1)
}
