// +build !notrue
package exec

import (
	"github.com/DataDrake/cli-ng/cmd"
	"os"
)

func init() {
    cmd.Register(&True)
}

var True = cmd.Sub {
    Name: "true",
    Short: "do nothing, successfully",
    Flags: &TrueFlags{},
    Args: &TrueArgs{},
    Run: TrueRun,
}

type TrueFlags struct {}
type TrueArgs struct {}

func TrueRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    // flags := c.Flags.(*TrueFlags)
    // args := c.Args.(*TrueArgs)
	os.Exit(0)
}
