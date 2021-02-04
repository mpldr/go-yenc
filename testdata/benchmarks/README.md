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

## Raw Speed

Raw speed is calculated by running `go test -bench=.` 100 times and taking the
average. This is done to account for variations in CPU Usage as this test is
completed pretty quick.

| Algorithm    | ns/Op Escaped | ns/Op Unescaped | ns/Op (exp. avg.)¹ | *n*th fastest |
|--------------|---------------|-----------------|--------------------|---------------|
| naive        | 2.46          | 2.43            | 2.43               | 1             |
| lookup-table | 2.55          | 2.55            | 2.55               | 2             |

¹) assuming random distribution of bytes and that 4/256 bytes have to be escaped.

## Data Throughput
