package main

import "leetcode"

// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/?envType=daily-question&envId=2025-02-12
func maximumSum(nums []int) int {
	sum2Idx := make(map[int][]int)
	for i, x := range nums {
		sum := 0
		for x > 0 {
			sum += x % 10
			x /= 10
		}
		sum2Idx[sum] = append(sum2Idx[sum], i)
	}
	ans := -1
	for _, v := range sum2Idx {
		if len(v) < 2 {
			continue
		}
		mx1 := -1
		mx2 := -1
		for _, x := range v {
			if mx1 == -1 || nums[x] > mx1 {
				mx2 = mx1
				mx1 = nums[x]
			} else if mx2 == -1 || nums[x] > mx2 {
				mx2 = nums[x]
			}
		}
		temp := mx1 + mx2
		if temp > ans {
			ans = temp
		}
	}
	return ans
}

func check(nums []int, expect int) {
	output := maximumSum(nums)
	leetcode.Test(output == expect, "nums=%v output=%d expect=%d", nums, output, expect)
}

func main() {
	check([]int{18, 43, 36, 13, 7}, 54)
}
