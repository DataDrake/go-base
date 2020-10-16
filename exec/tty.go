// +build !notty
package exec

import (
	"fmt"
	"github.com/DataDrake/go-base/cmd"
	"os"
	"path/filepath"
	"syscall"
)

func init() {
    cmd.Register(&TTY)
}

var TTY = cmd.CMD {
    Name: "tty",
    Short: "print the file name of the terminal connected to standard input",
    Flags: &TTYFlags{},
    Args: &TTYArgs{},
    Run: TTYRun,
}

type TTYFlags struct {
    Silent bool `short:"s" long:"silent" desc:"print nothing, only return an exit status"`
    Quiet bool `short:"q" long:"quiet" desc:"alias for (s) silent"`
}

type TTYArgs struct {}

func findRdev(dev_id uint64, path string) (dev string, err error) {
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	files, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	var dev_stat syscall.Stat_t
	for _, f := range files {
		err = syscall.Stat(filepath.Join(path, f), &dev_stat)
		if err != nil {
			return
		}
		if dev_id == dev_stat.Dev {
			dev = filepath.Join(path, f)
			return
		}
	}
	return "", fmt.Errorf("device not found")
}

func TTYRun(r *cmd.Root, c *cmd.CMD) {
    // gFlags := r.Flags.(*GlobalFlags)
    flags := c.Flags.(*TTYFlags)
    // args := c.Args.(*TTYArgs)

	var term_stat syscall.Stat_t
	err := syscall.Fstat(int(os.Stdin.Fd()), &term_stat)
	if err != nil {
		if !flags.Silent && !flags.Quiet {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
	dev, err := findRdev(term_stat.Dev, "/dev/pts")
	if err != nil {
		if !flags.Silent && !flags.Quiet {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
	if len(dev) == 0 {
		dev, err = findRdev(term_stat.Dev, "/dev")
		if err != nil {
			if !flags.Silent && !flags.Quiet {
				fmt.Println(err.Error())
			}
			os.Exit(1)
		}
	}
	if !flags.Silent && !flags.Quiet {
		fmt.Println(dev)
	}
}
