package leetcode

import (
	"fmt"
	"reflect"
)

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func TestQuickSort() {
	f := func(arr []int, expect []int) {
		originStr := fmt.Sprint(arr)
		QuickSort(arr)
		Assert(reflect.DeepEqual(arr, expect), "heap sort arr: %v output: %v expect: %v", originStr, arr, expect)
	}
	f([]int{23, 1, -1, 100, 50, 24, -13, 21}, []int{-13, -1, 1, 21, 23, 24, 50, 100})
	f([]int{2, 1, 3}, []int{1, 2, 3})
	f([]int{1}, []int{1})
	f([]int{}, []int{})
}

func quickSort(arr []int, begin int, end int) {
	if end <= begin {
		return
	}
	pivotIndex := partition(arr, begin, end)
	quickSort(arr, begin, pivotIndex-1)
	quickSort(arr, pivotIndex+1, end)
}

func partition(arr []int, begin int, end int) int {
	pivot := arr[begin]
	for begin < end {
		for begin < end && arr[end] > pivot {
			end--;
		}
		arr[begin] = arr[end]
		for begin < end && arr[begin] <= pivot {
			begin++;
		}
		arr[end] = arr[begin]
	}
	arr[begin] = pivot
	return begin
}
