package main

import (
	"github.com/DataDrake/cli-ng/cmd"
	_ "github.com/DataDrake/go-base/exec"
)

type GlobalFlags struct{}

func main() {
	r := cmd.Root {
		Name: "go-base",
		Short: "Alternative to GNU Coreutils written in Go",
		Flags: &GlobalFlags{},
        Single: true,
	}
	r.Run()
}
