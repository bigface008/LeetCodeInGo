package main

import (
	"leetcode"
	"math"
	"reflect"
	"slices"
)

// https://leetcode.com/problems/block-placement-queries/description/
func getResults(queries [][]int) []bool {
	obstacles := make([]int, 0)
	obstacleMap := make(map[int]int) // obs -> next obs
	obstacles = append(obstacles, 0)
	obstacleMap[0] = math.MaxInt
	ans := make([]bool, 0)
	for _, q := range queries {
		if q[0] == 1 {
			// binary search & insert
			newObs := q[1]
			start, end := 0, len(obstacles)-1
			for start <= end {
				mid := (start + end) / 2
				midVal := obstacles[mid]
				if midVal >= newObs {
					end = mid - 1
				} else {
					start = mid + 1
				}
			}
			obstacles = slices.Insert(obstacles, start, newObs)
			if start+1 < len(obstacles) {
				obstacleMap[newObs] = obstacles[start+1]
			} else {
				obstacleMap[newObs] = math.MaxInt
			}
			if start-1 >= 0 {
				obstacleMap[obstacles[start-1]] = newObs
			}
		} else {
			x, sz := q[1], q[2]
			res := false
			for obs, nextObs := range obstacleMap {
				if obs <= x-sz {
					if nextObs-obs >= sz {
						res = true
						break
					}
				}
			}
			ans = append(ans, res)
		}
	}
	return ans
}

func main() {
	test("[[1,2],[2,3,3],[2,3,1],[2,2,2]]", []bool{false, true, true})
	test("[[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]", []bool{true, true, false})
}

func test(queriesStr string, expect []bool) {
	queries := leetcode.Make2DIntArray(queriesStr)
	output := getResults(queries)
	leetcode.Test(reflect.DeepEqual(output, expect), "block placement queries=%s expect=%v output=%v", queriesStr, expect, output)
}
