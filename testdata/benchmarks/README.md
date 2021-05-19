# Benchmarks

This is the lab. Here, various algorithms are competing for the crown of highest
encoding speed.

The benchmarks are performed using Go's integrated benchmarks and Hyperfine. The
categories are: raw speed and data throughput.

## Contenders

- naive implementation
- Lookup Table
	- Slice containing struct
	- Hashmap with byte-key
- Bootleg SIMD (do it with a 32/64-bit integer and split it up)

### not yet participating

- Bitwise Operations
- SIMD
- `io.Writer` implementation

## Raw Speed

`[benchmark.sh]` 

Raw speed is calculated by running the benchmark 100 times and taking the 
average. This is done to account for variations in CPU Usage as this test is
completed pretty quick.

| Algorithm     | ns/Op Escaped | ns/Op Unescaped | ns/Op (exp. avg.)¹ | *n*th fastest |
|---------------|---------------|-----------------|--------------------|---------------|
| naive         | 2.42          | 2.28            | 2.28               | 2             |
| naive-pointer | 22.72         | 22.70           | 22.72              | 6             |
| lookup-table  | 2.20          | 2.20            | 2.20               | 1             |
| hashmap       | 20.01         | 19.69           | 19.70              | 5             |
| bootleg-simd  | 16.49         | 10.52           | 10.62              | 3             |
| simd          | 15.42         | 11.83           | 11.77              | 4             |

¹) assuming random distribution of bytes and that 4/256 bytes have to be escaped.

## Data Throughput

`[data-throughput/benchmark.sh]`

Data Throughput is calculated by running the encoding function on a set of
randomly generated data which is written to a file. This operation is performed
on a ramdisk to get raw numbers.

| Algorithm     | Duration | Byte       | Throughput    | *n*th fastest | Speed relative to naive |
|---------------|----------|------------|---------------|---------------|-------------------------|
| naive         |  30.516  | 1073741824 | 33.5562 MiB/s | 3             | 1.00                    |
| naive-pointer |  30.752  | 1073741824 | 33.2986 MiB/s | 5             | 0.99                    |
| lookup-table  |  30.569  | 1073741824 |  33.498 MiB/s | 4             | 1.00                    |
| hashmap       |  63.524  | 1073741824 | 16.1199 MiB/s | 6             | 0.48                    |
| bootleg-simd  |   4.310  | 1073741824 | 237.587 MiB/s | 1             | 7.08                    |
| simd          |   4.314  | 1073741824 | 237.367 MiB/s | 2             | 7.07                    |

<!--
No idea why SIMD changes from coming out almost last to placing first. I'm not
complaining, but I am confused.
-->

Variations in speed may be due to changes in the input dataset and fluctuations
in computer activity.

## additional notes

To ensure maximum comparability, all fields are updated every time the benchmark
is run. This way they should give an estimate even when run on a different
system.
