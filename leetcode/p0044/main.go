package main

import "leetcode"

// https://leetcode.com/problems/wildcard-matching/description/
func isMatch(s string, p string) bool {
	// dp[i][j] 表示包括了s中前i个字符串的子串能否被p中前j个字符串的子串所匹配
	dp := make([][]bool, len(s)+1)
	for i := range len(dp) {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for i := 1; i <= len(p); i++ {
		if dp[0][i-1] {
			dp[0][i] = p[i-1] == '*'
		} else {
			dp[0][i] = false
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {

		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	test("aa", "a", false)
	test("aa", "*", true)
	test("cb", "?a", false)
}

func test(s string, p string, expect bool) {
	output := isMatch(s, p)
	leetcode.Test(output == expect, "is match s=%s p=%s expect=%t output=%t", s, p, expect, output)
}
