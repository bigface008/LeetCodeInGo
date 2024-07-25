package main

import "slices"

type pair struct {
	name   string
	height int
}

// https://leetcode.com/problems/sort-the-people/?envType=daily-question&envId=2024-07-22
func sortPeople(names []string, heights []int) []string {
	//sort(names, heights, 0, len(names) - 1)
	N := len(heights)
	pairs := make([]pair, N)
	for i := range N {
		pairs[i] = pair{name: names[i], height: heights[i]}
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		return b.height - a.height
	})
	for i := range N {
		names[i] = pairs[i].name
	}
	return names
}

//func partition(names []string, heights []int, start, end int) int {
//	pivot := heights[start]
//	i, j := start, end
//	for i <= j {
//		for i <= j && heights[]
//	}
//	return i
//}
//
//func sort(names []string, heights []int, start, end int) {
//	if (start >= end) {
//		return
//	}
//	pivotIdx := partition(names, heights, start, end)
//	sort(names, heights, start, pivotIdx)
//	sort(names, heights, pivotIdx + 1, end)
//}
