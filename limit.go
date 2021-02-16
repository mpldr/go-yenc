package yenc

import "runtime"

func GetLimit() int {
	goroutines := runtime.NumCPU()
	goroutines *= 4 // every CPU-Thread gets 4 Jobs on average

	if runtime.NumCPU() > 8 {
		goroutines *= 4 // since we have more than 4 threads assume a modern CPU
	}

	return goroutines
}
