package main

import (
	"leetcode"
	"slices"
)

// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/?envType=daily-question&envId=2025-03-21
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	dq := supplies
	deg := make(map[string]int, len(recipes))
	graph := make(map[string][]string)
	for i, r := range recipes {
		for _, s := range ingredients[i] {
			graph[s] = append(graph[s], r)
		}
		deg[r] = len(ingredients[i])
	}

	ans := make([]string, 0)
	for len(dq) != 0 {
		node := dq[0]
		dq = dq[1:]
		for _, s := range graph[node] {
			deg[s]--
			if deg[s] == 0 {
				dq = append(dq, s)
				ans = append(ans, s)
			}
		}
	}
	return ans
}

func check(recipes []string, ingredients [][]string, supplies []string, expect []string) {
	output := findAllRecipes(recipes, ingredients, supplies)
	leetcode.Test(slices.Equal(output, expect), "recipes=%v ingredients=%v supplies=%v output=%v expect=%v", recipes, ingredients, supplies, output, expect)
}

func main() {
	//check([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast"}, []string{})
	//check([]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}, []string{"bread"})
	check([]string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}, [][]string{{"d"}, {"hveml", "f", "cpivl"}, {"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"}, {"cpivl", "hveml", "zpmcz", "ju", "h"}, {"h", "fzjnm", "e", "q", "x"}, {"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, {"f", "hveml", "cpivl"}}, []string{"f", "hveml", "cpivl", "d"}, []string{})
}
