package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	half := len(s) / 2
	for i := 0; i <= half; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
