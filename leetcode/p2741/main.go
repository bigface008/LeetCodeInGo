package main

import (
	"leetcode"
)

type Pair struct {
	lastIdx int
	mask    int
}

const MOD int64 = 1000000007

// https://leetcode.cn/problems/special-permutations/?envType=daily-question&envId=2024-06-26
func specialPerm(nums []int) int {
	n := len(nums)
	dp := make([][]int64, 1<<n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		dp[1<<i][i] = 1
	}

	for state := 1; state < (1 << n); state++ {
		for i := 0; i < n; i++ {
			if state>>i&1 == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if i == j || state>>j&1 == 0 {
					continue
				}
				x := nums[i]
				y := nums[j]
				if x%y != 0 && y%x != 0 {
					continue
				}
				dp[state][i] = (dp[state][i] + dp[state^(1<<i)][j]) % MOD
			}
		}
	}

	var sum int64
	for i := 0; i < n; i++ {
		sum = (sum + dp[(1<<n)-1][i]) % MOD
	}
	return int(sum)
}

func checkLegal(nums []int, i0 int, i1 int) bool {
	return nums[i0]%nums[i1] == 0 || nums[i1]%nums[i0] == 0
}

func main() {
	test([]int{2, 3, 6}, 2)
	test([]int{1, 4, 3}, 2)
}

func test(nums []int, expect int) {
	output := specialPerm(nums)
	leetcode.Test(output == expect, "special perm: %v expect: %v output: %v", nums, expect, output)
}
