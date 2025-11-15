package task1

// - 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 使用异或特性：a^a=0, a^0=a，成对数字抵消，剩下落单的。
func SingleNumber(nums []int) int {
	var res int = 0
	for _, value := range nums {
		res ^= value
	}
	return res
}
