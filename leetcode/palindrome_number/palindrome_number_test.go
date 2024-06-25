package palindrome_number

import (
	"leetcode"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	f := func(x int, expect bool) {
		output := isPalindrome(x)
		leetcode.Assert(t, output == expect, "%v is palindrome %v output %v", x, expect, output)
	}
	f(121, true)
	f(-121, false)
	f(10, false)
}
