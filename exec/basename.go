// +build !nobasename
package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"path/filepath"
	"strings"
)

func init() {
    cmd.Register(&Basename)
}

var Basename = cmd.Sub {
    Name: "basename",
    Short: "strip directory and suffix from filenames",
    Flags: &BasenameFlags{},
    Args: &BasenameArgs{},
    Run: BasenameRun,
}

type BasenameFlags struct {
    Multiple bool   `short:"a" long:"multiple"    desc:"allow multiple paths to be specified"`
    Suffix   string `short:"s" long:"trim-suffix" desc:"remove a trailing SUFFIX; implies -a"`
    Zero     bool   `short:"z" long:"null"        desc:"end each output line with NUL, not newline"`
}

type BasenameArgs struct {
    Paths []string `desc:"Path(s) to compute basename"`
}

func getBase(path string, suffix string, zero bool) {
	base := filepath.Base(path)
	if len(suffix) > 0 {
		base = strings.TrimSuffix(base, suffix)
	}
	if zero {
		fmt.Printf("%s\000", base)
	} else {
		fmt.Println(base)
	}
}

func BasenameRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*BasenameFlags)
    args := c.Args.(*BasenameArgs)

	if len(args.Paths) == 0 {
		r.SubUsage(c)
		os.Exit(1)
	}

	if flags.Multiple || len(flags.Suffix) > 0 {
		for _, path := range args.Paths {
			getBase(path, flags.Suffix, flags.Zero)
		}
	} else {
		if len(args.Paths) > 0 {
			r.SubUsage(c)
			os.Exit(1)
		}
		getBase(args.Paths[0], "", flags.Zero)
	}
}
