//
// Copyright 2016-2020 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// +build !notty

package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"os"
	"path/filepath"
	"syscall"
)

func init() {
	cmd.Register(&TTY)
}

// TTY implements the "tty" subcommand
var TTY = cmd.Sub{
	Name:  "tty",
	Short: "print the file name of the terminal connected to standard input",
	Flags: &TTYFlags{},
	Run:   TTYRun,
}

// TTYFlags are flags unique to the "tty" subcommand
type TTYFlags struct {
	Silent bool `short:"s" long:"silent" desc:"print nothing, only return an exit status"`
	Quiet  bool `short:"q" long:"quiet" desc:"alias for (s) silent"`
}

func findRdev(devID uint64, path string) (dev string, err error) {
	dir, err := os.Open(path)
	if err != nil {
		return
	}
	files, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	var devStat syscall.Stat_t
	for _, f := range files {
		err = syscall.Stat(filepath.Join(path, f), &devStat)
		if err != nil {
			return
		}
		if devID == devStat.Dev {
			dev = filepath.Join(path, f)
			return
		}
	}
	return "", fmt.Errorf("device not found")
}

// TTYRun carries out the "tty" subcommand
func TTYRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	flags := c.Flags.(*TTYFlags)
	var termStat syscall.Stat_t
	err := syscall.Fstat(int(os.Stdin.Fd()), &termStat)
	if err != nil {
		if !flags.Silent && !flags.Quiet {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
	dev, err := findRdev(termStat.Dev, "/dev/pts")
	if err != nil {
		if !flags.Silent && !flags.Quiet {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
	if len(dev) == 0 {
		dev, err = findRdev(termStat.Dev, "/dev")
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
