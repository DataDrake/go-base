// +build !noecho
package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"strings"
)

func init() {
    cmd.Register(&Echo)
}

var Echo = cmd.Sub {
    Name: "echo",
    Short: "display a line of text",
    Flags: &EchoFlags{
		NoEscapes: true,
	},
    Args: &EchoArgs{},
    Run: EchoRun,
}

type EchoFlags struct {
    NoNewlines bool `short:"n" long:"no-newlines" desc:"do not output the trailing newline"`
    Escapes    bool `short:"e" long:"escapes"     desc:"enable interpretation of backslash escapes"`
    NoEscapes  bool `short:"E" long:"no-escapes"  desc:"disable interpretation of backslash escapes (default)"`
}

type EchoArgs struct {
    Strings []string `desc:"strings to print"`
}

func EchoRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*EchoFlags)
    args := c.Args.(*EchoArgs)

	if flags.Escapes {
		fmt.Println("Escapes are not currently supported")
		os.Exit(1)
	}

	for _, v := range args.Strings {
		if flags.Escapes && !flags.NoEscapes {
			v = strings.Replace(v, "\\n", "\n", -1)
			fmt.Printf("%s", v)
		} else {
			fmt.Printf("%s", v)
		}
		if !flags.NoNewlines {
			println()
		}
	}
}
