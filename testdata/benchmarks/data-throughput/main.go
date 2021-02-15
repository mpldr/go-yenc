package main

import (
	"bufio"
	"bytes"
	"fmt"

	yenc "git.sr.ht/~poldi1405/go-yenc/testdata/benchmarks"
)

func main() {
	var b byte
	var e bool

	reader := bufio.NewReader(bytes.NewReader(indata))
	var eof bool
	var err error

	for !eof {
		b, err = reader.ReadByte()
		if err != nil {
			eof = true
		}

		b, e = yenc.YEnc(b)
		if e {
			fmt.Print('=')
		}
		fmt.Print(b)
	}
}
