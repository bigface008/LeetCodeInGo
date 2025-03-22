package main

import (
	"leetcode"
	"slices"
)

// https://leetcode.cn/problems/number-of-good-paths/description/
func findParent(node int, parents []int) int {
	for node != parents[node] {
		parents[node] = parents[parents[node]]
		node = parents[node]
	}
	return node
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	N := len(vals)
	graph := make([][]int, N)
	for i := range N {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	ids := make([]int, N)
	parents := make([]int, N)
	for i := range N {
		ids[i] = i
		parents[i] = i
	}
	sizes := make([]int, N)
	for i := range N {
		sizes[i] = 1
	}
	slices.SortFunc(ids, func(x, y int) int {
		return vals[x] - vals[y]
	})

	ans := N
	for _, x := range ids {
		vx := vals[x]
		fx := findParent(x, parents)
		for _, y := range graph[x] {
			y = findParent(y, parents)
			if y == fx || vals[y] > vx {
				continue
			}
			if vals[y] == vx {
				ans += sizes[fx] * sizes[y]
				sizes[fx] += sizes[y]
			}
			parents[y] = fx
		}
	}
	return ans
}

func check(vals []int, edges [][]int, expect int) {
	output := numberOfGoodPaths(vals, edges)
	leetcode.Test(output == expect, "vals=%v edges=%v output=%d expect=%d", vals, edges, output, expect)
}

func main() {
	check([]int{1, 3, 2, 1, 3}, leetcode.Make2DIntArray("[[0,1],[0,2],[2,3],[2,4]]"), 6)
}
