package main

import "leetcode"

// https://leetcode.com/problems/maximum-path-quality-of-a-graph/description
func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	edgeMap := make(map[int]map[int]int)
	for _, edge := range edges {
		mp, ok := edgeMap[edge[0]]
		if !ok {
			edgeMap[edge[0]] = make(map[int]int)
			edgeMap[edge[0]][edge[1]] = edge[2]
		} else {
			mp[edge[1]] = edge[2]
		}
		mp, ok = edgeMap[edge[1]]
		if !ok {
			edgeMap[edge[1]] = make(map[int]int)
			edgeMap[edge[1]][edge[0]] = edge[2]
		} else {
			mp[edge[0]] = edge[2]
		}
	}
	maxVal := 0
	used := make([]bool, len(values))
	used[0] = true
	dfs(0, values[0], 0, &maxVal, maxTime, values, edgeMap, used)
	return maxVal
}

func dfs(currIdx int, currValSum int, currTimeSum int, maxVal *int, maxTime int, values []int, edgeMap map[int]map[int]int, used []bool) {
	if currTimeSum > maxTime {
		return
	}
	if currIdx == 0 {
		*maxVal = max(*maxVal, currValSum)
	}
	edges, ok := edgeMap[currIdx]
	if !ok {
		return
	}
	for neighbor, time := range edges {
		newVal := currValSum
		needRevert := false
		if !used[neighbor] {
			newVal += values[neighbor]
			used[neighbor] = true
			needRevert = true
		}
		dfs(neighbor,
			newVal,
			currTimeSum+time,
			maxVal,
			maxTime,
			values,
			edgeMap,
			used,
		)
		if needRevert {
			used[neighbor] = false
		}
	}
}

func main() {
	test([]int{0, 32, 10, 43}, "[[0,1,10],[1,2,15],[0,3,10]]", 49, 75)
	test([]int{5, 10, 15, 20}, "[[0,1,10],[1,2,10],[0,3,10]]", 30, 25)
	test([]int{1, 2, 3, 4}, "[[0,1,10],[1,2,11],[2,3,12],[1,3,13]]", 50, 7)
	test([]int{0, 1, 2}, "[[1,2,10]]", 10, 0)
}

func test(values []int, edgesStr string, maxTime int, expect int) {
	edges := leetcode.Make2DIntArray(edgesStr)
	output := maximalPathQuality(values, edges, maxTime)
	leetcode.Test(output == expect, "maximum path quality values=%v edges=%v maxTime=%v expect=%v output=%v", values, edges, maxTime, expect, output)
}
