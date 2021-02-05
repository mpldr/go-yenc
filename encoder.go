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
	workercount := 1

	workqueue := make(chan [2]chan byte, workercount*2)
	results := make(chan chan byte, workercount*2)
	for i := 0; i < workercount; i++ {
		go encodeWorker(ctx, workqueue)
	}
	var writeerror error
	writingFinished := make(chan bool)
	go func() {err := resultWriter(ctx, output, results, cancel); writeerror = err; writingFinished <- true}()

	for {
		inchan := make(chan byte, BUFFERSIZE)
		outchan := make(chan byte, BUFFERSIZE)
		bytecount, err := bufrd.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		for i := 0; i < bytecount; i++ {
			inchan <- buf[i]
		}
		close(inchan)
		workqueue <- [2]chan byte{inchan, outchan}
		results <- outchan
		if err == io.EOF || bytecount <1 {
			break
		}
	}
	close(workqueue)
	close(results)

	<-writingFinished
	if writeerror != nil {
		return writeerror
	}
	
	return nil
}

func resultWriter(ctx context.Context, w io.Writer, results chan chan byte, done context.CancelFunc) error {
	for result := range results { //TODO: close when done
		var buf []byte
		for b := range result {
			buf = append(buf, b)
		}
		_, err := w.Write(buf)
		if err != nil {
			return err
		}
	}
	return nil
}

func encodeWorker(ctx context.Context, input chan [2]chan byte) {
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
