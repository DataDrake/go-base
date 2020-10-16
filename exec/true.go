// +build !notrue
package exec

import (
	"github.com/DataDrake/go-base/cmd"
	"os"
)

func init() {
    cmd.Register(&True)
}

var True = cmd.CMD {
    Name: "true",
    Short: "do nothing, successfully",
    Flags: &TrueFlags{},
    Args: &TrueArgs{},
    Run: TrueRun,
}

type TrueFlags struct {}
type TrueArgs struct {}

func TrueRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    // flags := c.Flags.(*TrueFlags)
    // args := c.Args.(*TrueArgs)
	os.Exit(0)
}
