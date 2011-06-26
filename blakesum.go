// blakesum command calculates BLAKE-256 and BLAKE-224 checksum of files.
package main

import (
	"flag"
	"fmt"
	"github.com/dchest/blake256"
	"hash"
	"io"
	"os"
)

var is224 = flag.Bool("224", false, "Use BLAKE-224")

func calcSum(f *os.File) (sum []byte, err os.Error) {
	var h hash.Hash
	if *is224 {
		h = blake256.New224()
	} else {
		h = blake256.New()
	}
	_, err = io.Copy(h, f)
	sum = h.Sum()
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		// Read from stdin.
		sum, err := calcSum(os.Stdin)
		if err != nil {
			fmt.Println(os.Stderr, "*** error reading from stdin")
			os.Exit(1)
		}
		fmt.Printf("%x\n", sum)
		os.Exit(0)
	}
	var hashname string
	if *is224 {
		hashname = "BLAKE-224"
	} else {
		hashname = "BLAKE-256"
	}
	exitNo := 0
	for i := 0; i < flag.NArg(); i++ {
		filename := flag.Arg(i)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "*** error opening %q\n", filename)
			exitNo = 1
			continue
		}
		sum, err := calcSum(f)
		f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "*** error reading %q\n", filename)
			exitNo = 1
			continue
		}
		fmt.Printf("%s (%s) = %x\n", hashname, filename, sum)
	}
	os.Exit(exitNo)
}
