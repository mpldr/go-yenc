package yenc

import (
	"testing"
)

func TestEncoding(t *testing.T) {
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"lookup-table", YEncLT},
		{"naive", YEnc},
		{"hashmap", YEncHashmap},
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
	encoderPtr := []struct {
		name string
		fn   func(*[2]byte)
	}{
		{"naive-pointer", YEncPtr},
		{"cgo", CYEnc},
	}
	for _, enc := range encoderPtr {
		t.Run(enc.name, func(t *testing.T) {
			for i := 0; i < 256; i++ {
				b := [2]byte{0, uint8(i)}
				enc.fn(&b)
				if b[1] != lookupTable[i].bte || (b[0] == 0x3D) != lookupTable[i].esc {
					t.Logf("Encoding of %x returned %x but %x was expected", i, b, lookupTable[i].bte)
					t.Fail()
				}
			}
		})
	}
}

func TestEncodingMultibyte(t *testing.T) {
	encoder := []struct {
		name string
		fn   func([8]byte) []byte
	}{
		{"bootleg-simd", YEncBootlegSIMD},
	}
	for _, enc := range encoder {
		t.Run(enc.name, func(t *testing.T) {
			var input [8]byte
			for i := 0; i < 256; i++ {
				index := i % 8
				input[index] = uint8(i)

				if index == 7 {
					for j := 0; j < 8; j++ {
						// TODO: add test
						_ = enc.fn([8]byte{5, 19, 0, 20, 18, 128, 64})
						/*if b != lookupTable[i].bte || esc != lookupTable[i].esc {
							t.Logf("Encoding of %x returned %x but %x was expected", i, b, lookupTable[i].bte)
							t.Fail()
						}*/
					}
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
		{"hashmap", YEncHashmap},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b, e := enc.fn(5)
				if e {
					_ = []byte{'=', b}
				} else {
					_ = []byte{b}
				}
			}
		})
	}
	encoderPtr := []struct {
		name string
		fn   func(*[2]byte)
	}{
		{"naive-pointer", YEncPtr},
		{"cgo", CYEnc},
	}
	for _, enc := range encoderPtr {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < 256; i++ {
				b := [2]byte{0, uint8(i)}
				enc.fn(&b)
				_ = b
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
		{"hashmap", YEncHashmap},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = enc.fn(19)
			}
		})
	}
}

func BenchmarkEncoding8Byte(b *testing.B) {
	indata := [8]byte{5, 5, 5, 5, 5, 5, 5, 5}
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
		{"hashmap", YEncHashmap},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 8; j++ {
					_, _ = enc.fn(indata[j])
				}
			}
		})
	}

	multiencoder := []struct {
		name string
		fn   func([8]byte) []byte
	}{
		{"bootleg-simd", YEncBootlegSIMD},
	}
	for _, enc := range multiencoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = enc.fn(indata)
			}
		})
	}
}

func BenchmarkEncoding8ByteEscaped(b *testing.B) {
	indata := [8]byte{19, 19, 19, 19, 19, 19, 19, 19}
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
		{"hashmap", YEncHashmap},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 8; j++ {
					_, _ = enc.fn(indata[j])
				}
			}
		})
	}

	multiencoder := []struct {
		name string
		fn   func([8]byte) []byte
	}{
		{"bootleg-simd", YEncBootlegSIMD},
	}
	for _, enc := range multiencoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = enc.fn(indata)
			}
		})
	}
}

func BenchmarkEncoding16Byte(b *testing.B) {
	indata := [16]byte{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
		{"hashmap", YEncHashmap},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 16; j++ {
					_, _ = enc.fn(5)
				}
			}
		})
	}

	multiencoder := []struct {
		name string
		fn   func([8]byte) []byte
	}{
		{"bootleg-simd", YEncBootlegSIMD},
	}
	for _, enc := range multiencoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = enc.fn([8]byte{5, 5, 5, 5, 5, 5, 5, 5})
				_ = enc.fn([8]byte{5, 5, 5, 5, 5, 5, 5, 5})
			}
		})
	}

	simdEncoder := []struct {
		name string
		fn   func([16]byte) []byte
	}{
		{"simd", YEncSIMD},
	}
	for _, enc := range simdEncoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = enc.fn(indata)
			}
		})
	}
}

func BenchmarkEncoding16ByteEscaped(b *testing.B) {
	indata := [16]byte{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19}
	encoder := []struct {
		name string
		fn   func(byte) (byte, bool)
	}{
		{"naive", YEnc},
		{"lookup-table", YEncLT},
		{"hashmap", YEncHashmap},
	}
	for _, enc := range encoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for j := 0; j < 16; j++ {
					_, _ = enc.fn(indata[j])
				}
			}
		})
	}

	multiencoder := []struct {
		name string
		fn   func([8]byte) []byte
	}{
		{"bootleg-simd", YEncBootlegSIMD},
	}
	for _, enc := range multiencoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = enc.fn([8]byte{19, 19, 19, 19, 19, 19, 19, 19})
			}
		})
	}

	simdEncoder := []struct {
		name string
		fn   func([16]byte) []byte
	}{
		{"simd", YEncSIMD},
	}
	for _, enc := range simdEncoder {
		b.Run(enc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = enc.fn(indata)
			}
		})
	}
}
