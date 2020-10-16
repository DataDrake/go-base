// +build !noyes
package exec

import (
    "bufio"
	"github.com/DataDrake/cli-ng/cmd"
	"strings"
    "os"
)

func init() {
    cmd.Register(&Yes)
}

var Yes = cmd.Sub {
    Name: "yes",
    Short: "output a string repeatedly until killed",
    Flags: &YesFlags{},
    Args: &YesArgs{},
    Run: YesRun,
}

type YesFlags struct {}

type YesArgs struct {
    Strings []string `desc:"String(s) to print out"`
}

func YesRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    // flags := c.Flags.(*YesFlags)
    args := c.Args.(*YesArgs)
	s := "y"
	if len(args.Strings) > 0 {
		s = strings.Join(args.Strings, " ")
	}
    s += "\n"
    var b []byte
    for i := 0; i < 1000; i++ {
        b = append(b, []byte(s)...)
    }
    buff := bufio.NewWriterSize(os.Stdout,1000*len(s))
	for {
		buff.Write(b)
	}
}
