// blakesum command calculates BLAKE-224, -256, -384, -512 checksums of files.
package main

import (
	"flag"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/dchest/blake256"
	"github.com/dchest/blake512"
)

var algorithms = map[int]func() hash.Hash{
	224: blake256.New224,
	256: blake256.New,
	384: blake512.New384,
	512: blake512.New,
}

var algoFlag = flag.Int("a", 256, "algorithm: 224, 256, 384, 512")

func calcSum(f *os.File, h hash.Hash) (sum []byte, err error) {
	h.Reset()
	_, err = io.Copy(h, f)
	sum = h.Sum(nil)
	return
}

func main() {
	flag.Parse()

	fn, ok := algorithms[*algoFlag]
	if !ok {
		flag.Usage()
		os.Exit(1)
	}

	h := fn()

	if flag.NArg() == 0 {
		// Read from stdin.
		sum, err := calcSum(os.Stdin, h)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%x\n", sum)
		os.Exit(0)
	}
	exitNo := 0
	for i := 0; i < flag.NArg(); i++ {
		filename := flag.Arg(i)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "(%s) %s\n", filename, err)
			exitNo = 1
			continue
		}
		sum, err := calcSum(f, h)
		f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "(%s) %s\n", filename, err)
			exitNo = 1
			continue
		}
		fmt.Printf("BLAKE-%d (%s) = %x\n", h.Size()*8, filename, sum)
	}
	os.Exit(exitNo)
}
