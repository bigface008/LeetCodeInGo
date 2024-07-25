package main

import "leetcode"

// https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/
func minimumOperations(num string) int {
	zeroIdx, fiveIdx := -1, -1
	zeroCnt := 0
	for i := len(num) - 1; i >= 0; i-- {
		ch := num[i]
		switch ch {
		case '0':
			zeroIdx = i
			zeroCnt++
			if zeroCnt == 2 {
				cnt := 0
				for j := i + 1; j < len(num); j++ {
					if num[j] != '0' {
						cnt++
					}
				}
				return cnt
			}
		case '2':
			if fiveIdx != -1 {
				cnt := 0
				for j := i + 1; j < len(num); j++ {
					if j == fiveIdx {
						continue
					}
					cnt++
				}
				return cnt
			}
		case '5':
			if zeroIdx != -1 {
				cnt := 0
				for j := i + 1; j < len(num); j++ {
					if j == zeroIdx {
						continue
					}
					cnt++
				}
				return cnt
			}
			fiveIdx = i
		case '7':
			if fiveIdx != -1 {
				cnt := 0
				for j := i + 1; j < len(num); j++ {
					if j == fiveIdx {
						continue
					}
					cnt++
				}
				return cnt
			}
		}
	}
	return len(num) - zeroCnt
}

func main() {
	test("2245047", 2)
	test("2908305", 3)
	test("10", 1)
}

func test(num string, expect int) {
	output := minimumOperations(num)
	leetcode.Test(output == expect, "minimum operations num=%s expect=%d output=%d", num, expect, output)
}
