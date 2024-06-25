package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	f := func(digits []int, expect []int) {
		output := plusOne(digits)
		Assert(t, reflect.DeepEqual(output, expect), "plus one digist %v expect %v output %v", digits, expect, output)
	}
	f([]int{1, 2, 3}, []int{1, 2, 4})
	f([]int{4, 3, 2, 1}, []int{4, 3, 2, 2})
	f([]int{9}, []int{1, 0})
}
