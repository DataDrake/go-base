package exec

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	_ "golang.org/x/crypto/blake2b"
	"fmt"
	"io"
	"os"
)

type SumFlags struct {
    Binary bool `short:"b" long:"binary" desc:"read in binary mode (default)"`
    Check  bool `short:"c" long:"check"  desc:"read sums from the FILEs and check them"`
	Tag    bool `short:"T" long:"tag"    desc:"create a BSD-style checksum"`
    Text   bool `short:"t" long:"text"   desc:"read in text mode (ignored)"`
	Zero   bool `short:"z" long:"zero"   desc:"end each output line with NUL, not newline"`
}

type SumArgs struct {
    Files []string `desc:"File(s) to encode or decode"`
}

type hashSum struct {
    Type string
    Mode rune
    File string
    Sum  []byte
}

func (s hashSum) Print(term rune, length int) {
    fmt.Printf("%x %c%s%c", s.Sum[:length], s.Mode, s.File, term)
}

func (s hashSum) PrintTag(term rune, length int) {
    fmt.Printf("%s (%s) = %x%c", s.Type, s.File, s.Sum[:length], term)
}

var hashNames = map[crypto.Hash]string {
    crypto.BLAKE2b_512: "BLAKE2b",
    crypto.MD5:         "MD5",
    crypto.SHA1:        "SHA1",
    crypto.SHA224:	    "SHA224",
    crypto.SHA256:	    "SHA256",
    crypto.SHA384:	    "SHA384",
    crypto.SHA512:	    "SHA512",
}

func sumFile(file string, mode rune, hash crypto.Hash) (sum hashSum, err error) {
    input, err := os.Open(file)
    if err != nil {
        return
    }
    defer input.Close()
    h := hash.New()
    if _, err = io.Copy(h, input); err != nil {
        return
    }
    sum = hashSum {
        Type: hashNames[hash],
        Mode: mode,
        File: file,
        Sum:  h.Sum(nil),
    }
    return
}

func sumAll(files []string, mode rune, hash crypto.Hash) (sums []hashSum, err error) {
    var sum hashSum
    for _, file := range files {
        if sum, err = sumFile(file, mode, hash); err != nil {
            return
        }
        sums = append(sums, sum)
    }
    return
}

func sumRun(flags *SumFlags, args *SumArgs, hash crypto.Hash, length int) {
	mode := ' '
	if flags.Binary {
        flags.Text = false
        mode = '*'
    }
    term := '\n'
    if flags.Zero {
        term = '\000'
    }
	if (length / 8) > hash.Size() {
		fmt.Printf("Digest length '%d' is greater than hash size '%d'\n", length, hash.Size()*8)
		os.Exit(1)
	}
	if (length % 8) != 0 {
		fmt.Printf("Digest length '%d' is not a multiple of 8\n", length)
		os.Exit(1)
	}
    sums, err := sumAll(args.Files, mode, hash)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    for _, sum := range sums {
        if flags.Tag {
            sum.PrintTag(term, length / 8)
        } else {
            sum.Print(term, length / 8)
        }
    }
}
