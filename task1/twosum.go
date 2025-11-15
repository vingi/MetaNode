package task1

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
// 你可以按任意顺序返回答案。
func TwoSum(nums []int, target int) []int {
	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 遍历数组中剩余的元素
		for j := i + 1; j < len(nums); j++ {
			// 如果找到目标值
			if nums[i]+nums[j] == target {
				// 返回下标
				return []int{i, j}
			}
		}
	}
	// 如果没有找到目标值，返回空数组
	return []int{}
}
