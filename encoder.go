package yenc

import (
	"bufio"
	"context"
	"io"
	"os"
)

func (y *Encoder) EncodeFile(fh *os.File, output io.Writer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bufrd := bufio.NewReader(fh)
	var buf []byte

	// TODO: make it dynamic
	workercount := 16

	workqueue := make(chan [2]chan byte, workercount*2)
	results := make(chan chan byte, workercount*2)
	for i := 0; i < workercount; i++ {
		go encodeWorker(ctx, workqueue, results)
	}

	for {
		inchan := make(chan byte, BUFFERSIZE)
		outchan := make(chan byte, BUFFERSIZE*2)
		bytecount, err := bufrd.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		for i := 0; i < bytecount; i++ {
			inchan <- buf[i]
		}
		close(inchan)
		workqueue <- [2]chan byte{inchan, outchan}
		if err == io.EOF {
			break
		}
	}

	for result := range results {
		var buf []byte
		i := 0
		for b := range result {
			buf[i] = b
			i++
		}
		_, err := output.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
}

func encodeWorker(ctx context.Context, input chan [2]chan byte, output chan chan byte) {
	for job := range input {
		for bte := range job[0] {
			b, e := YEnc(bte)
			if e {
				job[1] <- '='
			}
			job[1] <- b
		}
		close(job[1])
	}
}

const BUFFERSIZE = 4096
