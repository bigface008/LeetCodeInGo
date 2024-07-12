package main

import (
	"leetcode"
	"slices"
)

// https://leetcode.com/problems/minimum-number-game/description/
func numberGame(nums []int) []int {
	return []int{}
}

func heapify() {

}

func main() {
	test([]int{5, 4, 2, 3}, []int{3, 2, 5, 4})
	test([]int{2, 5}, []int{5, 2})
}

func test(nums []int, expect []int) {
	output := numberGame(nums)
	leetcode.Test(slices.Equal(expect, output), "number game nums=%v expect=%v output=%v", nums, expect, output)
}
