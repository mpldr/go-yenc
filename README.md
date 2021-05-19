# go-yenc

A yenc encoder/decoder who wants to be the fastest. (except yencode) (see 
[benchmarks](https://git.sr.ht/~poldi1405/go-yenc/tree/master/item/testdata/benchmarks/README.md))

## Objective

~~The current objective is a single-threaded throughput of at least 10 MiB/s
without causing a CPU-Meltdown or stealing too much RAM from Chrome.~~

I think we can safely say that we managed to hit this goal. Now it's time for
the actual implementation. The details may change if a faster way occurs to me.

![current_speed=237.59MiB](https://img.shields.io/badge/current_speed-237.59_MiB%2Fs-green)
![ram_usage=6424KiB](https://img.shields.io/badge/RAM_Usage-6.27_KiB-green)
![cpu_usage=100%](https://img.shields.io/badge/CPU_Usage-1_Core-green)

## License

The code is Licensed under the MPL 2.0

Copyright (c) 2020 Moritz Poldrack
