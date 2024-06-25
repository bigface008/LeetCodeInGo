package main

import (
	"leetcode"
	"reflect"
)

// https://leetcode.com/problems/find-a-good-subset-of-the-matrix/description/
func goodSubsetofBinaryMatrix(grid [][]int) []int {
	maskToIdx := map[int]int{}
	for i, row := range grid {
		mask := 0
		for j, x := range row {
			mask |= x << j
		}
		if mask == 0 {
			return []int{i}
		}
		maskToIdx[mask] = i
	}

	for x, i := range maskToIdx {
		for y, j := range maskToIdx {
			if x == y {
				continue
			}
			if x&y == 0 {
				return []int{min(i, j), max(i, j)}
			}
		}
	}
	return []int{}
}

func main() {
	test("[[0,1,1,0],[0,0,0,1],[1,1,1,1]]", "[0,1]")
	test("[[0]]", "[0]")
	test("[[1,1,1],[1,1,1]]", "[]")
}

func test(gridStr string, expectStr string) {
	grid := leetcode.Make2DIntArray(gridStr)
	expect := leetcode.MakeIntArray(expectStr)
	output := goodSubsetofBinaryMatrix(grid)
	leetcode.Test(reflect.DeepEqual(output, expect), "good subset of binary matrix: %v expect: %v output: %v", grid, expect, output)
}
