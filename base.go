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

package main

import (
	"github.com/DataDrake/cli-ng/cmd"
	_ "github.com/DataDrake/go-base/exec"
)

func main() {
	r := cmd.Root{
		Name:   "go-base",
		Short:  "Alternative to GNU Coreutils written in Go",
		Single: true,
	}
    cmd.Register(&cmd.GenManPages)
    cmd.Register(&cmd.GenSingleLinks)
	r.Run()
}
