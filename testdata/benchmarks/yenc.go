package yenc

import "unsafe"

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

func YencBootlegSIMD(input [8]byte) []byte {
	var mask [8]bool

	// add 42 (0x2a) to every byte
	*(*uint64)(unsafe.Pointer(&input)) += 0x2a2a2a2a2a2a2a2a

	for i := 0; i < 8; i++ {
		bte := (*(*[8]byte)(unsafe.Pointer(&input)))[i]

		if bte == 0x00 || bte == 0x0A || bte == 0x0D || bte == 0x3D {
			(*(*[8]byte)(unsafe.Pointer(&input)))[i] += uint8(64)
			mask[i] = true
		}
	}

	var result []byte
	for i := 0; i < 8; i++ {
		if mask[i] {
			result = append(result, '=')
		}
		result = append(result, input[i])
	}
	return result
}
