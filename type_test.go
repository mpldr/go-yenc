package yenc

import "testing"

func TestNewEncoder(t *testing.T) {
	// Did I say GetLimit was useless? This one is even less useful… but I
	// absolutely hate red lines in the Coverage Report
	y := NewEncoder()
	if y.LineLength != 128 {
		t.Fail()
	}
}
