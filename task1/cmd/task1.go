package main

import (
	"MetaNode/task1"
	"fmt"
)

func main() {
	// 测试 task1 包中的 singleNumber 函数
	nums := []int{2, 2, 1}
	result := task1.SingleNumber(nums)
	fmt.Println("只出现一次的数字是:", result)

	// 测试 task1 包中的 isPalindrome 函数
	num := 121
	isPalindrome := task1.Palindrome(num)
	fmt.Println("数字", num, "是否为回文数:", isPalindrome)

	// 测试 task1 包中的 ValidParentheses 函数
	testCases := []string{
		"()",     // 有效
		"()[]{}", // 有效
		"(]",     // 无效
		"([)]",   // 无效
		"{[]}",   // 有效
		"",       // 有效（空字符串）
	}

	fmt.Println("\n有效括号测试:")
	for _, testCase := range testCases {
		isValid := task1.ValidParentheses(testCase)
		fmt.Printf("'%s' 是否有效: %v\n", testCase, isValid)
	}

	// 测试 task1 包中的 LongestCommonPrefix 函数
	fmt.Println("\n最长公共前缀测试:")
	// 测试用例 1: 正常情况
	testCase1 := []string{"flower", "flow", "flight"}
	prefix1 := task1.LongestCommonPrefix(testCase1)
	fmt.Printf("'%v' 的最长公共前缀: %s\n", testCase1, prefix1)

	// 测试用例 2: 无公共前缀
	testCase2 := []string{"dog", "racecar", "car"}
	prefix2 := task1.LongestCommonPrefix(testCase2)
	fmt.Printf("'%v' 的最长公共前缀: %s\n", testCase2, prefix2)

	// 测试用例 3: 空字符串数组
	testCase3 := []string{}
	prefix3 := task1.LongestCommonPrefix(testCase3)
	fmt.Printf("'%v' 的最长公共前缀: %s\n", testCase3, prefix3)

	// 测试 PlusOne 函数
	fmt.Println("\nPlusOne 测试:")
	testDigits := [][]int{
		{1, 2, 3},    // 123 -> 124
		{9, 9, 9},    // 999 -> 1000
		{1, 9, 9},    // 199 -> 200
		{4, 3, 2, 1}, // 4321 -> 4322
	}

	for _, digits := range testDigits {
		result := task1.PlusOne(digits)
		fmt.Printf("%v + 1 = %v\n", digits, result)
	}

	// 测试 RemoveDuplicates 函数
	fmt.Println("\nRemoveDuplicates 测试:")
	testArrays := [][]int{
		{1, 1, 2, 2, 3, 3, 4}, // -> [1 2 3 4 4 4 4]
		{1, 1, 1, 1},          // -> [1 1 1 1]
		{1, 2, 3, 4, 5},       // -> [1 2 3 4 5]
		{},                    // -> []
		{1},                   // -> []
	}

	for _, arr := range testArrays {
		// 创建副本以便在测试中保护原数据
		testArr := make([]int, len(arr))
		copy(testArr, arr)
		length := task1.RemoveDuplicates(testArr)
		fmt.Printf("%v 去重后，长度=%d: %v\n", arr, length, testArr[:length])
	}

	// 测试 TwoSum 函数
	fmt.Println("\nTwoSum 测试:")
	testTwoSum := [][]int{
		{2, 7, 11, 15}, // 目标值 9，返回 [0, 1]
		{3, 2, 4},      // 目标值 6，返回 [1, 2]
		{3, 3},         // 目标值 6，返回 [0, 1]
	}
	targets := []int{9, 6, 6}

	for i, testCase := range testTwoSum {
		result := task1.TwoSum(testTwoSum[i], targets[i])
		fmt.Printf("在 %v 中找到目标值 %d 的下标: %v\n", testCase, targets[i], result)
	}

	// 测试 Merge 函数
	fmt.Println("\nMerge 测试:")
	testIntervals := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, // -> [[1, 6], [8, 10], [15, 18]]
		{{1, 4}, {4, 5}},                    // -> [[1, 5]]
		{{1, 4}, {2, 3}},                    // -> [[1, 4]]
		{{1, 4}},                            // -> [[1, 4]]
		{},                                  // -> []
	}

	for i, intervals := range testIntervals {
		// 创建副本
		testIntervalsCopy := make([][]int, len(intervals))
		for j, interval := range intervals {
			testIntervalsCopy[j] = make([]int, len(interval))
			copy(testIntervalsCopy[j], interval)
		}
		result := task1.Merge(testIntervalsCopy)
		fmt.Printf("测试用例 %d: %v\n  合并结果: %v\n", i+1, intervals, result)
	}
}
