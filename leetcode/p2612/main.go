package main

import (
	"leetcode"
	"slices"
)

func minReverseOperations(n int, p int, banned []int, k int) []int {
	dq := make([]int, 0, n)
	dq = append(dq, p)
	ban := make(map[int]bool)
	for _, x := range banned {
		ban[x] = true
	}
	ans := make([]int, n)
	for i := range n {
		ans[i] = -1
	}
	step := 0
	visited := make([]bool, n)
	visited[p] = true
	for len(dq) != 0 {
		levelSize := len(dq)
		for _ = range levelSize {
			node := dq[0]
			dq = dq[1:]
			ans[node] = step
			start := max(0, node-k+1)
			end := min(n-k, node)
			for i := start; i <= end; i++ {
				subNode := k - 1 - node + 2*i
				if subNode != node && !ban[subNode] && !visited[subNode] {
					visited[subNode] = true
					dq = append(dq, subNode)
				}
			}
		}
		step++
	}
	return ans
}

func check(n int, p int, banned []int, k int, expect []int) {
	output := minReverseOperations(n, p, banned, k)
	leetcode.Test(slices.Equal(output, expect), "n=%d p=%d banned=%v k=%d output=%v expect=%v", n, p, banned, k, output, expect)
}

func main() {
	check(4, 0, []int{1, 2}, 4, []int{0, -1, -1, 1})
}
