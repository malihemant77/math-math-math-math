package main

import "testing"

func TestChecksum(t *testing.T) {
	cmp := func(a, b int) bool {
		return checksum(a) == b
	}

	tests := []int{
		0, 0,
		1, 1,
		10, 1,
		18, 9,
		19, 1,
		200, 2,
		1000, 1,
		1999, 1,
		99998888, 5,
	}

	for i := 0; i < len(tests); i += 2 {
		a := tests[i]
		b := tests[i+1]
		if !cmp(a, b) {
			t.Errorf("checksum(%d) is not %d", a, b)
		}
	}
}
