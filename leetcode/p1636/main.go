package main

import "slices"

// https://leetcode.com/problems/sort-array-by-increasing-frequency/?envType=daily-question&envId=2024-07-23
func frequencySort(nums []int) []int {
	numToFreq := make(map[int]int) // val to freq
	for _, num := range nums {
		numToFreq[num]++
	}
	freqToNum := make(map[int][]int)
	for num, freq := range numToFreq {
		if _, ok := freqToNum[freq]; !ok {
			freqToNum[freq] = []int{}
		}
		freqToNum[freq] = append(freqToNum[freq], num)
	}
	ans := make([]int, 0)
	frequencies := make([]int, 0)
	for freq, _ := range freqToNum {
		frequencies = append(frequencies, freq)
	}
	slices.Sort(frequencies)
	for _, freq := range frequencies {
		values := freqToNum[freq]
		slices.Sort(values)
		for i := len(values) - 1; i >= 0; i-- {
			for _ = range freq {
				ans = append(ans, values[i])
			}
		}
	}
	return ans
}
