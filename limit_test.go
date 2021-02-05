package yenc

import (
	"runtime"
	"testing"
)

func TestGetLimit(t *testing.T) {
	// this test is essentially useless, but hey… free coverage.

	limit := GetLimit()
	if runtime.NumCPU() > 8 {
		if limit != runtime.NumCPU()*16 {
			t.Fail()
		}
	} else {
		if limit != runtime.NumCPU()*4 {
			t.Fail()
		}
	}
}
