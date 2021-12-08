package yenc

import (
	"unsafe"

	"github.com/yesuu/simd"
)

func YEnc(input byte) (byte, bool) {
	escape := false

	input += uint8(42)

	if input == 0x00 || input == 0x0A || input == 0x0D || input == 0x3D {
		input += uint8(64)
		escape = true
	}

	return input, escape
}

func YEncPtr(input *[2]byte) {
	input[1] += 42

	if input[1] == 0x00 || input[1] == 0x0A || input[1] == 0x0D || input[1] == 0x3D {
		input[1] += 0x40
		input[0] = 0x3D
	}
}

func YEncHashmap(input byte) (byte, bool) {
	inp := hashmap[input]
	return inp.bte, inp.esc
}

func YEncLT(input byte) (byte, bool) {
	inp := lookupTable[int(input)]
	return inp.bte, inp.esc
}

func YEncBootlegSIMD(input [8]byte) []byte {
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

func YEncSIMD(input [16]byte) []byte {
	in := simd.Uint8x16{
		input[0],
		input[1],
		input[2],
		input[3],
		input[4],
		input[5],
		input[6],
		input[7],
		input[8],
		input[9],
		input[10],
		input[11],
		input[12],
		input[13],
		input[14],
		input[15],
	}

	in = simd.AddUint8x16(in, mask42)

	mask := simd.Uint8x16{}

	for i, v := range in {
		if v == 0x00 || v == 0x0A || v == 0x0D || v == 0x3D {
			mask[i] = 0x40
		}
	}

	in = simd.AddUint8x16(in, mask)

	var result []byte

	for i, v := range mask {
		if v != 0 {
			result = append(result, 0x3D)
		}
		result = append(result, in[i])
	}
	return result
}

var mask42 = simd.Uint8x16{
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
	0x2a,
}
