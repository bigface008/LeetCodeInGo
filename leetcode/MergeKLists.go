package leetcode

func MergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	newLength := length / 2
	left := MergeKLists(lists[:newLength])
	right := MergeKLists(lists[newLength:])
	return Merge2Lists(left, right)
}

func Merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	node := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	} else {
		node.Next = l2
	}
	return res.Next
}

func TestMergeKLists() {
	f := func(lists [][]int, expect []int) {
		input := []*ListNode{}
		for _, list := range lists {
			input = append(input, NewList(list))
		}
		output := MergeKLists(input)
		Assert(output.EqualsArr(expect),"merge k lists: %v expect: %v output: %v", lists, expect, output)
	}
	f([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6})
	f([][]int{}, []int{})
	f([][]int{{}}, []int{})
}
