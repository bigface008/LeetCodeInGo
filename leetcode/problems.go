package leetcode

// https://leetcode.com/problems/plus-one/description/?envType=study-plan-v2&envId=top-interview-150
func plusOne(digits []int) []int {
	size := len(digits)
	if size == 0 {
		return []int{1}
	}
	carry := 0
	for i := size - 1; i >= 0; i-- {
		temp := carry + digits[i]
		if i == size-1 {
			temp += 1
		}
		if temp >= 10 {
			temp -= 10
			digits[i] += temp
			carry = 1
		} else {
			digits[i] += temp
			carry = 0
		}
	}
	if carry == 1 {
		digits = append(digits, 0)
		copy(digits[1:], digits[0:])
		digits[0] = 1
	}
	return digits
}
