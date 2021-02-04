package yenc

func YEnc(input byte) (byte, bool) {
	escape := false

	input += uint8(42)

	if input == 0x00 || input == 0x0A || input == 0x0D || input == 0x3D {
		input += uint8(64)
		escape = true
	}

	return input, escape
}

func YEncLT(input byte) (byte, bool) {
	inp := lookupTable[int(input)]
	return inp.bte, inp.esc
}
