package merge_k_lists

import (
	. "leetcode"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	f := func(lists [][]int, expect []int) {
		var input []*ListNode
		for _, list := range lists {
			input = append(input, NewList(list))
		}
		output := MergeKLists(input)
		Assert(t, output.EqualsArr(expect), "merge k lists: %v expect: %v output: %v", lists, expect, output)
	}
	f([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6})
	f([][]int{}, []int{})
	f([][]int{{}}, []int{})
}
