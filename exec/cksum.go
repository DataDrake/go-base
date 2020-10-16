// +build !nocksum
package exec

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"hash/adler32"
	"io"
	"os"
)

func init() {
    cmd.Register(&CkSum)
}

var CkSum = cmd.Sub {
    Name: "cksum",
    Short: "checksum and count the bytes in a file",
    Flags: &CkSumFlags{},
    Args: &CkSumArgs{},
    Run: CkSumRun,
}

type CkSumFlags struct {}

type CkSumArgs struct {
    File string `desc:"File to encode or decode"`
}

func CkSumRun(r *cmd.Root, c *cmd.Sub) {
    // gFlags := r.Flags.(*GlobalFlags)
    // flags := c.Flags.(*CkSumFlags)
    args := c.Args.(*CkSumArgs)

	input, err := os.Open(args.File)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	h := adler32.New()
	size, err := io.Copy(h, input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%d %d %s\n", h.Sum32(), size, args.File)
}
