package main

import (
	"leetcode"
	"math"
	"slices"
)

type pair struct {
	origin int
	mapped int
}

// https://leetcode.com/problems/sort-the-jumbled-numbers/description/?envType=daily-question&envId=2024-07-24
func sortJumbled(mapping []int, nums []int) []int {
	convMap := make(map[int][]int) // mapped -> [origin]
	for _, num := range nums {
		mapped := convert(num, mapping)
		if _, ok := convMap[mapped]; !ok {
			convMap[mapped] = []int{}
		}
		convMap[mapped] = append(convMap[mapped], num)
	}
	keys := []int{}
	for key, _ := range convMap {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	ans := []int{}
	for _, key := range keys {
		values := convMap[key]
		for _, val := range values {
			ans = append(ans, val)
		}
	}
	return ans
}

func convert(num int, mapping []int) int {
	if num == 0 {
		return mapping[0]
	}
	res := 0
	cnt := 0
	for num > 0 {
		remainder := num % 10
		after := mapping[remainder]
		res += int(math.Pow(10, float64(cnt))) * after
		cnt++
		num /= 10
	}
	return res
}

func main() {
	test([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}, []int{338, 38, 991})
	test([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}, []int{123, 456, 789})
	test([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
}

func test(mapping, nums, expect []int) {
	output := sortJumbled(mapping, nums)
	leetcode.Test(slices.Equal(output, expect), "sort jumbled mapping=%v nums=%v expect=%v output=%v", mapping, nums, expect, output)
}
