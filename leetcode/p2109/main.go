package main

import (
	"strings"
	"sync"
)

// https://leetcode.com/problems/adding-spaces-to-a-string/?envType=daily-question&envId=2024-12-03
func addSpaces(s string, spaces []int) string {
	ans := strings.Builder{}
	prevIdx := 0
	for _, idx := range spaces {
		//ans += s[prevIdx:idx] + " "
		ans.WriteString(s[prevIdx:idx])
		ans.WriteString(" ")
		prevIdx = idx
	}
	ans.WriteString(s[prevIdx:])
	return ans.String()
}

const (
	_ = iota
	MONEY_USD
	MONEY_RMB
	MONEY_JPA
	_
	MONEY_STH
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
}
