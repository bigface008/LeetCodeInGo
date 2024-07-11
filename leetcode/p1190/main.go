package main

import (
	"leetcode"
)

// https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/?envType=daily-question&envId=2024-07-11
func reverseParentheses(s string) string {
	stk := make([][]byte, 0)
	stk = append(stk, []byte{})
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case '(':
			stk = append(stk, []byte{})
		case ')':
			top := stk[len(stk)-1]
			halfTopLen := len(top) / 2
			for i := 0; i < halfTopLen; i++ {
				top[i], top[len(top)-i-1] = top[len(top)-i-1], top[i]
			}
			stk = stk[:len(stk)-1]
			stk[len(stk)-1] = append(stk[len(stk)-1], top...)
		default:
			stk[len(stk)-1] = append(stk[len(stk)-1], ch)
		}
	}
	ans := string(stk[0][:])
	return ans
}

func main() {
	test("(abcd)", "dcba")
	test("(u(love)i)", "iloveu")
	test("(ed(et(oc))el)", "leetcode")
}

func test(s string, expect string) {
	output := reverseParentheses(s)
	leetcode.Test(output == expect, "reverse parentheses s=%s expect=%s output=%s", s, expect, output)
}
