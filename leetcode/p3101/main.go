package main

import "leetcode"

// https://leetcode.cn/problems/count-alternating-subarrays/
func countAlternatingSubarrays(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int64, len(nums))
	dp[0] = 1
	ans := int64(1)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		ans += dp[i]
	}
	return ans
}

func main() {
	test([]int{0, 1, 1, 1}, 5)
	test([]int{1, 0, 1, 0}, 10)
}

func test(nums []int, expect int64) {
	output := countAlternatingSubarrays(nums)
	leetcode.Test(output == expect, "count alternating subarrays nums=%v output=%d expect=%d", nums, output, expect)
}
