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

package exec

import (
	"crypto"
	_ "crypto/md5"    // Needed
	_ "crypto/sha1"   // Needed
	_ "crypto/sha256" // Needed
	_ "crypto/sha512" // Needed
	"fmt"
	_ "golang.org/x/crypto/blake2b" // Needed
	"io"
	"os"
)

// SumFlags are flags shared by all of the "*sum" subcommands
type SumFlags struct {
	Binary bool `short:"b" long:"binary" desc:"read in binary mode (default)"`
	Check  bool `short:"c" long:"check"  desc:"read sums from the FILEs and check them"`
	Tag    bool `short:"T" long:"tag"    desc:"create a BSD-style checksum"`
	Text   bool `short:"t" long:"text"   desc:"read in text mode (ignored)"`
	Zero   bool `short:"z" long:"zero"   desc:"end each output line with NUL, not newline"`
}

// SumArgs are args shares by all of the "*sum" subcommands
type SumArgs struct {
	Files []string `desc:"File(s) to encode or decode"`
}

type hashSum struct {
	Type string
	Mode rune
	File string
	Sum  []byte
}

// Print a hash line
func (s hashSum) Print(term rune, length int) {
	fmt.Printf("%x %c%s%c", s.Sum[:length], s.Mode, s.File, term)
}

// PrintTag prints a tag formatted hash line
func (s hashSum) PrintTag(term rune, length int) {
	fmt.Printf("%s (%s) = %x%c", s.Type, s.File, s.Sum[:length], term)
}

// hashnames maps a cypher to its string name
var hashNames = map[crypto.Hash]string{
	crypto.BLAKE2b_512: "BLAKE2b",
	crypto.MD5:         "MD5",
	crypto.SHA1:        "SHA1",
	crypto.SHA224:      "SHA224",
	crypto.SHA256:      "SHA256",
	crypto.SHA384:      "SHA384",
	crypto.SHA512:      "SHA512",
}

// sumFile calculates the hahs for a single file
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
	sum = hashSum{
		Type: hashNames[hash],
		Mode: mode,
		File: file,
		Sum:  h.Sum(nil),
	}
	return
}

// sumAll gagses all of the files in a directory
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

// sumRun implements the basis of the "*sum" subcommands
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
			sum.PrintTag(term, length/8)
		} else {
			sum.Print(term, length/8)
		}
	}
}
