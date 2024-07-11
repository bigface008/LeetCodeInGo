package main

import "leetcode"

func incremovableSubarrayCount(nums []int) int64 {
	N := len(nums)
	firstAscendIdx := 0
	for i := 1; i < N; i++ {
		if nums[i] > nums[i-1] {
			firstAscendIdx = i
		} else {
			break
		}
	}
	if firstAscendIdx == N-1 {
		return int64(N * (N + 1) / 2)
	}
	var ans int64
	ans = int64(firstAscendIdx + 2)
	i := firstAscendIdx
	j := N - 1
	for j == N-1 || nums[j] < nums[j+1] {
		for i >= 0 && nums[i] >= nums[j] {
			i--
		}
		ans += int64(i + 2)
		j--
	}
	return ans
}

func main() {
	test([]int{1, 2, 3, 4}, 10)
	test([]int{6, 5, 7, 8}, 7)
	test([]int{8, 7, 6, 6}, 3)
}

func test(nums []int, expect int64) {
	output := incremovableSubarrayCount(nums)
	leetcode.Test(output == expect, "incremovable subarray count nums=%v expect=%d output=%d", nums, expect, output)
}
