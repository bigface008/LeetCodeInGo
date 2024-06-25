package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	f := func(arr []int, expect []int) {
		originStr := fmt.Sprint(arr)
		HeapSort(arr)
		Assert(t, reflect.DeepEqual(arr, expect), "heap sort arr: %v output: %v expect: %v", originStr, arr, expect)
	}
	f([]int{23, 1, -1, 100, 50, 24, -13, 21}, []int{-13, -1, 1, 21, 23, 24, 50, 100})
	f([]int{2, 1, 3}, []int{1, 2, 3})
	f([]int{1}, []int{1})
	f([]int{}, []int{})
}

func TestQuickSort(t *testing.T) {
	f := func(arr []int, expect []int) {
		originStr := fmt.Sprint(arr)
		QuickSort(arr)
		Assert(t, reflect.DeepEqual(arr, expect), "heap sort arr: %v output: %v expect: %v", originStr, arr, expect)
	}
	f([]int{23, 1, -1, 100, 50, 24, -13, 21}, []int{-13, -1, 1, 21, 23, 24, 50, 100})
	f([]int{2, 1, 3}, []int{1, 2, 3})
	f([]int{1}, []int{1})
	f([]int{}, []int{})
}
