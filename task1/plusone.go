package task1

// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func PlusOne(digits []int) []int {
	// 先将数组转为整数
	num := 0
	for _, digit := range digits {
		num = num*10 + digit
	}
	// 加 1
	num++
	// 将整数转为数组
	digits = []int{}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		// fmt.Printf("num: %d, digits: %v\n", num, digits)
		num /= 10
	}
	return digits
}
