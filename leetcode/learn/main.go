package main

import "fmt"

func main() {
	//slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := slice[2:5]
	//s2 := slice[3:6]
	//s2 = append(s2, 100)
	//s1 = append(s1, 100)
	//s1[2] = 20
	//fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(slice)

	var slice []int
	if slice == nil {
		fmt.Println("Is Nil")
	}
	slice = append(slice, 1)
	fmt.Println(slice)
}
