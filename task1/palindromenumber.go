package task1

// 判断一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数
func Palindrome(x int) bool {
	// 负数和以0结尾的非零数字不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedNum := 0
	// 反转后半部分数字
	for x > reversedNum {
		reversedNum = reversedNum*10 + x%10
		x = x / 10
	}

	// 当数字长度为奇数时，需要去掉中间位
	// 例如：12321，x=12，reversedNum=123，去掉最后一位3
	return x == reversedNum || x == reversedNum/10
}
