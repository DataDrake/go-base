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

// WhoAmI implements the "whoami" subcommand
var WhoAmI = cmd.Sub{
	Name:  "whoami",
	Short: "print effective userid",
	Run:   WhoAmIRun,
}

// WhoAmIRun carries out the "whoami" subcommand
func WhoAmIRun(r *cmd.Root, c *cmd.Sub) {
	// gFlags := r.Flags.(*GlobalFlags)
	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(u.Username)
}
