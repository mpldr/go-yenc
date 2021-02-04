package main

import (
	"fmt"

	yenc "git.sr.ht/~poldi1405/go-yenc/testdata/benchmarks"
)

func main() {
	for i := 0; i < len(indata); i++ {
		fmt.Print(yenc.YEnc(indata[i]))
	}
}
