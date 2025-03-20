package main

import "leetcode"

// https://leetcode.cn/problems/satisfiability-of-equality-equations/
func equationsPossible(equations []string) bool {
	parents := make([]int, 26)
	ranks := make([]int, 26)
	for i := range 26 {
		parents[i] = i
	}

	findParent := func(a int) int {
		for a != parents[a] {
			parents[a] = parents[parents[a]]
			a = parents[a]
		}
		return a
	}

	merge := func(a int, b int) {
		pa := findParent(a)
		pb := findParent(b)
		if pa == pb {
			return
		}
		if ranks[pa] > ranks[pb] {
			parents[pb] = pa
		} else if ranks[pb] > ranks[pa] {
			parents[pa] = pb
		} else {
			parents[pb] = pa
			ranks[pa]++
		}
	}

	notEqualPairs := make([][]int, 0)
	for _, eq := range equations {
		a := int(eq[0] - 'a')
		b := int(eq[3] - 'a')
		if eq[1] == '!' {
			notEqualPairs = append(notEqualPairs, []int{a, b})
		} else {
			merge(a, b)
		}
	}
	for _, p := range notEqualPairs {
		a := p[0]
		b := p[1]
		if findParent(a) == findParent(b) {
			return false
		}
	}
	return true
}

func check(equations []string, expect bool) {
	output := equationsPossible(equations)
	leetcode.Test(expect == output, "expect=%v output=%v", expect, output)
}

func main() {
	//check([]string{"a==b", "b!=a"}, false)
	check([]string{"c==c", "b==d", "x!=z"}, true)
}
