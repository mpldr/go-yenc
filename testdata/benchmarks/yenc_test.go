package yenc

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
	}
	for _, enc := range encoder {
		t.Run(enc.name, func(t *testing.T) {
			for i := 0; i < 256; i++ {
				b, esc := enc.fn(uint8(i))
				if b != lookupTable[i].bte || esc != lookupTable[i].esc {
					t.Logf("Encoding of %x returned %x but %x was expected", i, b, lookupTable[i].bte)
					t.Fail()
				}
			}
		})
	}
}

func BenchmarkEncoding(b *testing.B) {
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = enc.fn(5)
			}
		})
	}
}

func BenchmarkEncodingEscape(b *testing.B) {
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = enc.fn(19)
			}
		})
	}
}
