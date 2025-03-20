package p1742

// https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/?envType=daily-question&envId=2025-02-13
func countBalls(lowLimit int, highLimit int) int {
	box2Num := make(map[int]int)
	ans := 0
	for i := lowLimit; i <= highLimit; i++ {
		digit := i
		num := 0
		for digit > 0 {
			num += digit % 10
			digit /= 10
		}
		box2Num[num]++
		if box2Num[num] > ans {
			ans = box2Num[num]
		}
	}
	return ans
}
