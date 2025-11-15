package task1

// 26. 删除有序数组中的重复项
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
func RemoveDuplicates(nums []int) int {
	// 为空判断
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		// fmt.Printf("fast: %d, slow: %d, nums[fast]: %d, nums[slow]: %d\n", fast, slow, nums[fast], nums[slow])
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
