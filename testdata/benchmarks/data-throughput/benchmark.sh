#!/bin/env bash

rm -r build

mkdir build

echo "Generating random Data…"
${_GO_EXE:-"go"} generate

echo "Generating naive…"
${_GO_EXE:-"go"} build -o "build/compare.bench.naive"

for p in patches/*; do
	echo "Generating $(basename ${p%.*})…"
	git apply "$p"
	${_GO_EXE:-"go"} build -o "build/bench.$(basename ${p%.*})"
	git checkout -- main.go
done

hyperfine build/*
