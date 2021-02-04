package main

import (
	_ "embed"
)

//go:generate dd if=/dev/urandom of=indata.dat bs=16M count=16 iflag=fullblock

//go:embed indata.dat
var indata []byte