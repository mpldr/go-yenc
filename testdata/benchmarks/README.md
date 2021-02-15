# Benchmarks

This is the lab. Here, various algorithms are competing for the crown of highest
encoding speed.

The benchmarks are performed using Go's integrated benchmarks and Hyperfine. The
categories are: raw speed and data throughput.

## Contenders

- naive implementation
- Lookup Table
- SIMD (not yet implemented)
- Bootleg SIMD (do it with a 32/64-bit integer and split it up)
- Bitwise Operations (not yet implemented)
- `io.Writer` implementation

## Raw Speed

`[benchmark.sh]` 

Raw speed is calculated by running the benchmark 100 times and taking the 
average. This is done to account for variations in CPU Usage as this test is
completed pretty quick.

| Algorithm    | ns/Op Escaped | ns/Op Unescaped | ns/Op (exp. avg.)¹ | *n*th fastest |
|--------------|---------------|-----------------|--------------------|---------------|
| naive        | 2.53          | 2.51            | 2.51               | 1             |
| lookup-table | 2.61          | 2.61            | 2.61               | 2             |
| bootleg-simd | 13.95         | 8.48            | 8.57               | 3             |

¹) assuming random distribution of bytes and that 4/256 bytes have to be escaped.

## Data Throughput

`[data-throughput/benchmark.sh]`

Data Throughput is calculated by running the encoding function on a set of
randomly generated data which is compiled into the program.

| Algorithm    | Duration | Byte      | Throughput    | *n*th fastest | Speed relative to naive |
|--------------|----------|-----------|---------------|---------------|-------------------------|
| naive        |  91.631  | 268435456 | 2.79381 MiB/s | 3             | 1.00 ± 0.00 times       |
| lookup-table |  89.875  | 268435456 | 2.84840 MiB/s | 2             | 1.01 ± 0.01 times       |
| bootleg-simd |  29.931  | 268435456 | 8.55301 MiB/s | 1             | 3.06 ± 0.01 times       |

// TODO: Try if using bufio in naive and LT levels the playingfield. Until then
// Bootleg-SIMD is the winner

## additional notes

To ensure maximum comparability, all fields are updated every time the benchmark
is run. This way they should give an estimate even when run on a different
system.
