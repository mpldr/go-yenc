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

| Algorithm    | ns/Op Escaped | ns/Op Unescaped | ns/Op (exp. avg.)¹ | *n*th fastest |
|--------------|---------------|-----------------|--------------------|---------------|
| naive        | 2.40          | 2.39            | 2.39               | 1             |
| lookup-table | 2.51          | 2.51            | 2.51               | 2             |
| hashmap      | 21.05         | 20.99           | 20.99              | 4
| bootleg-simd | 13.95         | 8.48            | 8.57               | 3             |

¹) assuming random distribution of bytes and that 4/256 bytes have to be escaped.

## Data Throughput

`[data-throughput/benchmark.sh]`

Data Throughput is calculated by running the encoding function on a set of
randomly generated data which is compiled into the program.

| Algorithm    | Duration | Byte      | Throughput    | *n*th fastest | Speed relative to naive |
|--------------|----------|-----------|---------------|---------------|-------------------------|
| naive        |  0.9828  | 268435456 |  260.48 MiB/s | 2             | 1.00                    |
| lookup-table |  0.8363  | 268435456 |  306.11 MiB/s | 1             | 1.18                    |
| hashmap      |  8.828   | 268435456 | 28.9986 MiB/s | 4             | 0.11                    |
| bootleg-simd |  4.647   | 268435456 | 55.0893 MiB/s | 3             | 0.21                    |

<!--
There was an extreme improvement by removing the fmt.Print() statements. This
also lead to a new ranking and we have definitely met the 10 MiB/s
-->

Variations in speed may be due to changes in the input dataset and fluctuations
in computer activity.

## additional notes

To ensure maximum comparability, all fields are updated every time the benchmark
is run. This way they should give an estimate even when run on a different
system.
