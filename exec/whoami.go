// +build !nowhoami
package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"os/user"
)

func init() {
    cmd.Register(&WhoAmI)
}

var WhoAmI = cmd.Sub {
    Name: "whoami",
    Short: "print effective userid",
    Flags: &WhoAmIFlags{},
    Args: &WhoAmIArgs{},
    Run: WhoAmIRun,
}

type WhoAmIFlags struct {}
type WhoAmIArgs struct {}

func WhoAmIRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    // flags := c.Flags.(*WhoAmIFlags)
    // args := c.Args.(*WhoAmIArgs)
	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(u.Username)
}
