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

	workercount := GetLimit()

	workqueue := make(chan chan byte, workercount*2)
	results := make(chan chan byte, workercount*2)
	for i := 0; i < workercount; i++ {
		go encodeWorker(ctx, workqueue, results)
	}

	for {
		inchan := make(chan byte, BUFFERSIZE)
		bytecount, err := bufrd.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		for i := 0; i < bytecount; i++ {
			inchan <- buf[i]
		}
		close(inchan)
		workqueue <- inchan
		if err == io.EOF {
			break
		}
	}
	close(workqueue)

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

func encodeWorker(ctx context.Context, input chan chan byte, output chan chan byte) {
	for job := range input {
		outchan := make(chan byte, BUFFERSIZE*2)
		for bte := range job {
			b, e := YEnc(bte)
			if e {
				outchan <- '='
			}
			outchan <- b
		}
		close(outchan)
		output <- outchan
	}
}

const BUFFERSIZE = 4096
