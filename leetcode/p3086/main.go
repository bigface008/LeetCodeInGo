package main

import (
	"leetcode"
	"math"
)

func minimumMoves(nums []int, k int, maxChanges int) int64 {
	// 1. calculate max continuous one group shorter than 3
	c := 0
	pos := make([]int, 0)
	for i, num := range nums {
		if num != 1 {
			continue
		}
		pos = append(pos, i)
		c = max(1, c)
		if i >= 1 && nums[i-1] == 1 {
			if i >= 2 && nums[i-2] == 1 {
				c = 3
			} else {
				c = max(c, 2)
			}
		}
	}

	// 2. check maxChange
	c = min(c, k)
	if maxChanges >= k-c {
		return int64(max(c-1, 0) + (k-c)*2)
	}

	// 3. build pre sum
	N := len(pos)
	preSum := make([]int, N+1)
	for i := range N {
		preSum[i+1] = preSum[i] + pos[i]
	}

	// 3. find smallest distance sum in each sliding window
	ans := math.MaxInt
	windowSize := k - maxChanges
	for right := windowSize; right <= N; right++ {
		left := right - windowSize
		i := left + windowSize/2
		midIdx := pos[i]
		leftSum := midIdx*(i-left) - (preSum[i] - preSum[left])
		rightSum := preSum[right] - preSum[i] - midIdx*(right-i)
		ans = min(ans, leftSum+rightSum)
	}
	return int64(ans + maxChanges*2)
}

func main() {
	test([]int{1, 0, 1}, 5, 3, 8)
}

func test(nums []int, k int, maxChanges int, expect int64) {
	output := minimumMoves(nums, k, maxChanges)
	leetcode.Test(output == expect, "minimum moves: %v %v %v expect: %v output: %v", nums, k, maxChanges, expect, output)
}
