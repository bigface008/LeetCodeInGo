package lotter_winner

import (
	"leetcode"
	"testing"
)

func TestGetMaximumLength(t *testing.T) {
	f := func(lottery string, winner string, k int, expect int) {
		output := GetMaximumLength(lottery, winner, k)
		leetcode.Assert(t, output == expect, "lottery '%v' winner '%v' k %v expect %v output %v", lottery, winner, k, expect, output)
	}
	f("basic", "amage", 8, 3)
	f("hackerrank", "fpelqanxyk", 8, 7)
}
