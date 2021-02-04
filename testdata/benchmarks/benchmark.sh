#!/bin/env bash

runs=100

mkdir results.d

for i in $(seq $runs); do
	echo -ne " running benchmark: $i/$runs\r"
	go test -bench=. | head -n -2 | tail -n +4 > results.d/$i
done

echo " running benchmark: completed"

awk 'NR==FNR{a[$1]=$3+" ";next;} {a[$1]=($1 in a)?a[$1] $3 " ":$3 " "}END{for(x in a)print x, a[x]}' results.d/* > results.d/joined

awk 'BEGIN{FS=" "}{ n=0; sum=0; for(i=1;i<NF;++i) { if( $i ) { ++n; sum += $i; } } print $1 ": " sum/n; }' results.d/joined | sort
