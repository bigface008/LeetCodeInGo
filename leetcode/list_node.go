package leetcode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(arr []int) *ListNode {
	extra := &ListNode{}
	node := extra
	for _, v := range arr {
		node.Next = &ListNode{v, nil}
		node = node.Next
	}
	return extra.Next
}

func (n *ListNode) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	node := n
	for node != nil {
		builder.WriteString(strconv.Itoa(node.Val))
		if node.Next != nil {
			builder.WriteString(" -> ")
		}
		node = node.Next
	}
	builder.WriteString("]")
	return builder.String()
}

func (n *ListNode) Equals(other *ListNode) bool {
	n1, n2 := n, other
	for n1 != nil && n2 != nil {
		if n1.Val != n2.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	if n1 != nil || n2 != nil {
		return false
	}
	return true
}

func (n *ListNode) EqualsArr(other []int) bool {
	node := n
	for _, v := range other {
		if node == nil || node.Val != v {
			return false
		}
		node = node.Next
	}
	return node == nil
}
