package yenc

import (
	"bytes"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"sync"
)

type Encoder struct {
	file     io.Writer
	buffer   bytes.Buffer
	writeMtx sync.Mutex

	LineLength int
	BlockSize  int
	CRC        hash.Hash32
}

func NewEncoder(w io.Writer) *Encoder {
	enc := &Encoder{
		file:       w,
		LineLength: 128,
		BlockSize:  0,
		CRC:        crc32.NewIEEE(),
	}
	return enc
}

func (e *Encoder) Write(slice []byte) (int, error) {
	e.writeMtx.Lock()
	defer e.writeMtx.Unlock()

	var bytesWritten int
	remainder := len(slice) % 8
	parts := getParts(&slice)

	maxindex := len(*parts) - 1

	var err error
	var n int
	var encoded []byte

	for i, p := range *parts {
		encoded = yenc(&p)

		if i == maxindex {
			break
		}

		n, err = e.buffer.Write(encoded)
		if err != nil {
			return bytesWritten, fmt.Errorf("failed writing yenc to output-buffer: %v")
		}
		bytesWritten += n
	}

	n, err = e.buffer.Write(encoded[:len(encoded)-(8-remainder)])
	if err != nil {
		return bytesWritten, fmt.Errorf("failed writing yenc to output-buffer: %v")
	}
	bytesWritten += n

	err = e.writeBlock(false)
	if err != nil {
		return n, err
	}

	return bytesWritten, nil
}

func (e *Encoder) writeBlock(closing bool) error {
	return nil
}
