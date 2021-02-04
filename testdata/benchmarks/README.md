# Benchmarks

This is the lab. Here, various algorithms are competing for the crown of highest
encoding speed.

The benchmarks are performed using Go's integrated benchmarks and Hyperfine. The
categories are: raw speed and data throughput.

## Contenders

- naive implementation
- Lookup Table
- SIMD (not yet implemented)
- Bootleg SIMD (do it with a 32/64-bit integer and split it up) (not yet implemented)
- Bitwise Operations (not yet implemented)
- `io.Writer` implementation

## Raw Speed

`[benchmark.sh](../benchmark.sh)` 

Raw speed is calculated by running the benchmark 100 times and taking the 
average. This is done to account for variations in CPU Usage as this test is
completed pretty quick.

| Algorithm    | ns/Op Escaped | ns/Op Unescaped | ns/Op (exp. avg.)¹ | *n*th fastest |
|--------------|---------------|-----------------|--------------------|---------------|
| naive        | 2.46          | 2.43            | 2.43               | 1             |
| lookup-table | 2.55          | 2.55            | 2.55               | 2             |

¹) assuming random distribution of bytes and that 4/256 bytes have to be escaped.

## Data Throughput

`[data-throughput/benchmark.sh](../data-throughput/benchmark.sh)`

Data Throughput is calculated by running the encoding function on a set of
randomly generated data which is compiled into the program.

| Algorithm    | Duration | Byte      | Throughput    | *n*th fastest | Difference to naive |
|--------------|----------|-----------|---------------|---------------|---------------------|
| naive        | 100.010  | 268435456 | 2.60645 MiB/s | 2             | 1.00 ± 0.00 times   |
| lookup-table |  98.875  | 268435456 | 2.65039 MiB/s | 1             | 1.01 ± 0.01 times   |

The reason for different results is not yet clear.
