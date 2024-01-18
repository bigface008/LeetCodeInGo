package leetcode

import (
	"fmt"
	"reflect"
)

func HeapSort(arr []int) {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, len(arr) - 1)
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		maxHeapify(arr, 0, i - 1)
	}
}

func TestHeapSort() {
	f := func(arr []int, expect []int) {
		originStr := fmt.Sprint(arr)
		HeapSort(arr)
		Assert(reflect.DeepEqual(arr, expect), "heap sort arr: %v output: %v expect: %v", originStr, arr, expect)
	}
	f([]int{23, 1, -1, 100, 50, 24, -13, 21}, []int{-13, -1, 1, 21, 23, 24, 50, 100})
	f([]int{2, 1, 3}, []int{1, 2, 3})
	f([]int{1}, []int{1})
	f([]int{}, []int{})
}

func maxHeapify(arr []int, begin int, end int) {
	parent := begin
	child := 2 * parent + 1
	for child <= end {
		if child + 1 <= end && arr[child + 1] > arr[child] {
			child++;
		}
		if arr[child] < arr[parent] {
			return
		} else {
			arr[parent], arr[child] = arr[child], arr[parent]
			parent = child
			child = 2 * parent + 1
		}
	}
}
