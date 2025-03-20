package leetcode

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"unicode"
)

func Test(condition bool, format string, a ...any) {
	if condition {
		fmt.Print("[PASSED] ")
	} else {
		fmt.Print("[FAILED] ")
	}
	fmt.Printf(format, a...)
	fmt.Println()
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testArrayMaker() {
	arr1 := MakeIntArray("[1,2,3,4,5]")
	fmt.Println(arr1)
	arr1 = MakeIntArray("[-2,3,-4,5]")
	fmt.Println(arr1)

	arr2 := Make2DIntArray("[[1,2,3],[4,5,6],[7,8,9]]")
	fmt.Println(arr2)
	arr2 = Make2DIntArray("[[1],[-4],[7]]")
	fmt.Println(arr2)
	arr2 = Make2DIntArray("[[-1]]")
	fmt.Println(arr2)
	arr2 = Make2DIntArray("[[]]")
	fmt.Println(arr2)
	arr2 = Make2DIntArray("[]")
	fmt.Println(arr2)
}

func MakeIntArray(arrStr string) []int {
	res := []int{}
	i := 0
	for i < len(arrStr) {
		ch := arrStr[i]
		if unicode.IsDigit(rune(ch)) || ch == '-' { // number start
			j := i + 1
			for ; j < len(arrStr); j++ {
				if !unicode.IsDigit(rune(arrStr[j])) {
					break
				}
			}
			numStr := arrStr[i:j]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			res = append(res, num)
			i = j
		} else {
			i++
		}
	}
	return res
}

func Make2DIntArray(arrStr string) [][]int {
	res := [][]int{}
	level := 0
	i := 0
	for i < len(arrStr) {
		ch := arrStr[i]
		if ch == '[' {
			if level == 1 {
				res = append(res, []int{})
			}
			level++
			i++
		} else if ch == ']' {
			level--
			i++
		} else if unicode.IsDigit(rune(ch)) || ch == '-' { // number start
			j := i + 1
			for ; j < len(arrStr); j++ {
				if !unicode.IsDigit(rune(arrStr[j])) {
					break
				}
			}
			numStr := arrStr[i:j]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			end := len(res) - 1
			res[end] = append(res[end], num)
			i = j
		} else {
			i++
		}
	}
	return res
}

//func Make2DStrArray(arrStr string) [][]string {
//	res := [][]string{}
//	level := 0
//	i := 0
//	for i < len(arrStr) {
//		ch := arrStr[i]
//		if ch == '[' {
//			if level == 1 {
//				res = append(res, []byte{})
//			}
//			level++
//			i++
//		} else if ch == ']' {
//			level--
//			i++
//		} else if ch == '"' || ch == '\'' { // number start
//			end := len(res) - 1
//			res[end] = append(res[end], arrStr[i+1])
//			i += 3
//		} else {
//			i++
//		}
//	}
//	return res
//}

func Make2DCharArray(arrStr string) [][]byte {
	res := [][]byte{}
	level := 0
	i := 0
	for i < len(arrStr) {
		ch := arrStr[i]
		if ch == '[' {
			if level == 1 {
				res = append(res, []byte{})
			}
			level++
			i++
		} else if ch == ']' {
			level--
			i++
		} else if ch == '"' || ch == '\'' { // number start
			end := len(res) - 1
			res[end] = append(res[end], arrStr[i+1])
			i += 3
		} else {
			i++
		}
	}
	return res
}

func Assert(t *testing.T, condition bool, format string, a ...any) {
	if !condition {
		t.Errorf(format, a...)
	}
}
