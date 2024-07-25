package main

import (
	"leetcode"
	"slices"
)

// https://leetcode.com/problems/sort-an-array
func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivotIdx := partition(nums, start, end)
	sort(nums, start, pivotIdx-1)
	sort(nums, pivotIdx+1, end)
}

func partition(nums []int, start, end int) int {
	pivot := nums[start]
	i, j := start, end
	for i < j {
		for i < j && nums[j] > pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	return i
}

func main() {
	test([]int{5, 2, 3, 1}, []int{1, 2, 3, 5})
	test([]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5})
}

func test(nums, expect []int) {
	output := sortArray(nums)
	leetcode.Test(slices.Equal(output, expect), "sort array nums=%v expect=%v output=%v", nums, expect, output)
}
