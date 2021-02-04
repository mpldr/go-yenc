# go-yenc

A yenc encoder/decoder who wants to be the fastest. (except yencode) (see 
[benchmarks](https://git.sr.ht/~poldi1405/go-yenc/tree/master/item/testdata/benchmarks/README.md))

## Objective

The current objective is a throughput of at least 10 MiB/s without causing a
CPU-Meltdown or stealing too much RAM from Chrome.

Current Speed: 2.65 MiB/s
