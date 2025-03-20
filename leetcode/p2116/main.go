package p2116

import "leetcode"

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	x := 0
	for i, ch := range s {
		if ch == '(' || locked[i] == '0' {
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}
	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}
	return true
}

func check(s string, locked string, expect bool) {
	output := canBeValid(s, locked)
	leetcode.Test(output == expect, "s=%s locked=%s output=%v expect=%v", s, locked, output, expect)
}

func main() {
	check(")", "0", false)
}
