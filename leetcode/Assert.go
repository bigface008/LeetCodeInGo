package leetcode

import "fmt"

func Assert(condition bool, name string) {
	if condition {
		print("[PASSED]")
	} else {
		print("[FAILED]")
	}
	fmt.Printf(" %v", name)
	println()
}
