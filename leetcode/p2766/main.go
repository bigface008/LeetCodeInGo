package main

import (
	"leetcode"
	"slices"
)

// https://leetcode.com/problems/relocate-marbles/description/
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	posSet := make(map[int]bool)
	for _, pos := range nums {
		posSet[pos] = true
	}
	for i := 0; i < len(moveFrom); i++ {
		from := moveFrom[i]
		to := moveTo[i]
		if !posSet[from] {
			continue
		}
		delete(posSet, from)
		posSet[to] = true
	}
	ans := make([]int, 0)
	for pos, _ := range posSet {
		ans = append(ans, pos)
	}
	slices.Sort(ans)
	return ans
}

func relocateMarblesV1(nums []int, moveFrom []int, moveTo []int) []int {
	posToIdx := make(map[int][]int)
	maxSize := len(nums)
	for i, pos := range nums {
		if _, ok := posToIdx[pos]; !ok {
			posToIdx[pos] = []int{}
		}
		posToIdx[pos] = append(posToIdx[pos], i)
	}
	for i := 0; i < len(moveFrom); i++ {
		from := moveFrom[i]
		to := moveTo[i]
		maxSize = max(from, maxSize)
		maxSize = max(to, maxSize)
		indices, ok := posToIdx[from]
		if !ok {
			continue
		}
		delete(posToIdx, from)
		if _, ok = posToIdx[to]; !ok {
			posToIdx[to] = []int{}
		}
		posToIdx[to] = append(posToIdx[to], indices...)
	}
	ans := make([]int, 0)
	for pos, _ := range posToIdx {
		ans = append(ans, pos)
	}
	slices.Sort(ans)
	return ans
}

func main() {
	test([]int{1, 6, 7, 8}, []int{1, 7, 2}, []int{2, 9, 5}, []int{5, 6, 8, 9})
	test([]int{1, 1, 3, 3}, []int{1, 3}, []int{2, 2}, []int{2})
}

func test(nums []int, moveFrom []int, moveTo []int, expect []int) {
	output := relocateMarbles(nums, moveFrom, moveTo)
	leetcode.Test(slices.Equal(output, expect), "relocate marbles nums=%v moveFrom=%v moveTo=%v expect=%v output=%v", nums, moveFrom, moveTo, expect, output)
}
