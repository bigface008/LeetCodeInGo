package main

import (
	"fmt"
	"leetcode"
)

// https://leetcode.cn/problems/maximum-or/?envType=daily-question&envId=2025-03-21
func maximumOr(nums []int, k int) int64 {
	N := len(nums)
	suffix := make([]int, N)
	for i := N - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] | nums[i+1]
	}

	ans, pre := 0, 0
	for i, x := range nums {
		ans = max(ans, pre|(x<<k)|suffix[i])
		pre |= x
	}
	return int64(ans)
}

//func maximumOr(nums []int, k int) int64 {
//	N := len(nums)
//	suffix := make([]int64, N)
//	suffix[N - 1] = int64(nums[N - 1])
//	for i := N - 2; i >= 0; i-- {
//		suffix[i] = suffix[i + 1] | int64(nums[i + 1])
//	}
//
//	var ans int64
//	var pre int64
//	for i, x := range nums {
//		ans = max(ans, pre | int64(x << k) | suffix[i])
//		pre |= int64(x)
//	}
//	return ans
//}

// dp[1][0] | 4 << 1 = 10 | 8 = 10
// dp[1][1] | 4 << 0

// 10100
//  1000

// 10 20
// 10
// 14

func check(nums []int, k int, expect int64) {
	output := maximumOr(nums, k)
	leetcode.Test(output == expect, "nums=%v k=%d output=%d expect=%d", nums, k, output, expect)
}

func main() {
	//check([]int{10, 8, 4}, 1, 30)
	s1 := []int{5, 10, 20, 30}
	for x := range s1 {
		fmt.Println(x)
	}
}
