package task1

import "sort"

// 56. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func Merge(intervals [][]int) [][]int {
	// 边界情况：空数组或只有一个区间
	if len(intervals) <= 1 {
		return intervals
	}

	// 第一步：按区间起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 第二步：合并重叠的区间
	result := [][]int{}
	for i := 0; i < len(intervals); i++ {
		// 获取当前区间
		currInterval := intervals[i]
		
		// 如果结果数组为空，或者当前区间与最后一个区间不重叠
		if len(result) == 0 || currInterval[0] > result[len(result)-1][1] {
			// 直接添加当前区间到结果中
			result = append(result, []int{currInterval[0], currInterval[1]})
		} else {
			// 当前区间与最后一个区间重叠，更新结束点
			if currInterval[1] > result[len(result)-1][1] {
				result[len(result)-1][1] = currInterval[1]
			}
		}
	}

	return result
}
