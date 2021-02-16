package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

var cpuprof = flag.Bool("cpu", false, "create CPU Profile")
var memprof = flag.Bool("mem", false, "create Memory Profile")
var traceprof = flag.Bool("trace", false, "create Trace Profile (this will severely impact performance)")
var file = flag.String("file", "", "file to encode")

func main() {
	flag.Parse()

	if *cpuprof {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprof {
		f, err := os.Create("mem.prof")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		defer pprof.WriteHeapProfile(f)
	}

	if *traceprof {
		f, err := os.Create("trace.prof")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		err = trace.Start(f)
		if err != nil {
			log.Fatal(err)
		}
		defer trace.Stop()
	}

	if *file == "" {
		return
	}
}
