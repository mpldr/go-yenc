package main

import (
	"fmt"
	"os"

	"git.sr.ht/~poldi1405/go-yenc"
)

func main() {
	y := yenc.NewEncoder()

	fh, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	outfile, err := os.Create("outfile")
	if err != nil {
		panic(err)
	}

	fmt.Println(y.EncodeFile(fh, outfile))
}
