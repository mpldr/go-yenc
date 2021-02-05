package yenc

import (
	"testing"

	"context"
)

func TestEncodeWorker(t *testing.T) {
	input := make(chan chan byte, 1)
	inchan := make(chan byte, BUFFERSIZE)
	resultchan := make(chan chan byte, 1)
	for i := 0; i < 2048; i++ {
		inchan <- uint8(i)
	}
	close(inchan)

	input <- inchan
	close(input)

	encodeWorker(context.Background(), input, resultchan)

	result := <-resultchan
	i := uint8(0)
	for resultbyte := range result {
		checkbyte, escape := YEnc(i)
		if escape {
			if resultbyte != '=' {
				t.Errorf("Unescaped byte: %d; got: %d", i, resultbyte)
			}
			resultbyte = <-result
		}
		if resultbyte != checkbyte {
			t.Errorf("%d was encoded as %d, but %d was expected", i, resultbyte, checkbyte)
		}
		i++
	}
}
