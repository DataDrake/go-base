// +build !nofalse
package exec

import (
	"github.com/DataDrake/cli-ng/cmd"
	"os"
)

func init() {
    cmd.Register(&False)
}

var False = cmd.Sub {
    Name: "false",
    Short: "do nothing, unsuccessfully",
    Flags: &FalseFlags{},
    Args: &FalseArgs{},
    Run: FalseRun,
}

type FalseFlags struct {}

type FalseArgs struct {}

func FalseRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    //flags := c.Flags.(*FalseFlags)
    //args := c.Args.(*FalseArgs)
	os.Exit(1)
}
