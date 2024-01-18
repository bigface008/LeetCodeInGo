package leetcode

import "fmt"

func Assert(condition bool, format string, a ...any) {
	if condition {
		print("[PASSED] ")
	} else {
		print("[FAILED] ")
	}
	fmt.Printf(format, a...)
	println()
}
