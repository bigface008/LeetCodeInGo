package main

import (
	"leetcode"
)

// https://leetcode.com/problems/target-sum/description
func findTargetSumWays(nums []int, target int) int {
	cnt := 0
	used := make([]bool, len(nums))
	backtrack(&cnt, target, 0, 0, used, nums)
	return cnt
}

func backtrack(cnt *int, target int, sum int, idx int, used []bool, nums []int) {
	if idx == len(nums) {
		if sum == target {
			*cnt++
		}
		return
	}
	used[idx] = true
	backtrack(cnt, target, sum+nums[idx], idx+1, used, nums)
	backtrack(cnt, target, sum-nums[idx], idx+1, used, nums)
	used[idx] = false
}

func findTargetSumWaysV2(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	neg := (sum - target)
	if neg < 0 || neg%2 == 1 {
		return 0 // impossible
	}
	neg /= 2

	N := len(nums)
	dp := make([][]int, N+1)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, neg+1)
	}
	// 1st line
	for j := 0; j < neg; j++ {
		if j == 0 {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[N][neg]
}

// func rec(remain int, nums []int, memo map[int]int) int {
// 	if remain == 0 {
// 		if len(nums) == 0 {
// 			return 1
// 		}
// 	}
// }

func main() {
	test([]int{1, 1, 1, 1, 1}, 3, 5)
	test([]int{1}, 1, 1)
}

func test(nums []int, target int, expect int) {
	output := findTargetSumWays(nums, target)
	leetcode.Test(output == expect, "find target sum ways: %v expect: %v output: %v", nums, expect, output)
}
