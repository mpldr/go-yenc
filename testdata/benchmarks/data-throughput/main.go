package main

import (
	"bufio"
	"os"

	yenc "git.sr.ht/~poldi1405/go-yenc/testdata/benchmarks"
)

func main() {
	var b byte
	var e bool

	infile, err := os.Open("indata.dat")
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	outfile, err := os.Create("outdata.dat")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	reader := bufio.NewReader(infile)
	writer := bufio.NewWriter(outfile)
	var eof bool

	for !eof {
		b, err = reader.ReadByte()
		if err != nil {
			eof = true
		}

		b, e = yenc.YEnc(b)
		if e {
			writer.Write([]byte{0x3D})
		}
		writer.Write([]byte{b})
	}
	writer.Flush()
}
