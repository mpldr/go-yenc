package yenc

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	for i := 0; i <= 255; i++ {
		b, esc := YEnc(uint8(i))
		if b != lookupTable[i].bte || esc != lookupTable[i].esc {
			t.Logf("Encoding of %x returned %x but %x was expected", i, b, lookupTable[i].bte)
			t.Fail()
		}
	}
}
