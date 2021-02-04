package yenc

// YEnc takes one byte and returns it's yenc equivalent.
func YEnc(input byte) (byte, bool) {
	inp := lookupTable[int(input)]
	return inp.bte, inp.esc
}
