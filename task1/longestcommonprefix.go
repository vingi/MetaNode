package task1

// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func LongestCommonPrefix(arrays []string) string {
	// 如果字符串数组为空，直接返回空字符串
	if len(arrays) == 0 {
		return ""
	}
	var prefix string = ""
	var prefixarray []byte
	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(arrays[i]); j++ {
			if i == 0 {
				prefixarray = append(prefixarray, arrays[i][j])
			} else {
				if j >= len(prefixarray) {
					break
				} else {
					if arrays[i][j] != prefixarray[j] {
						prefixarray = prefixarray[:j]
					}
				}
			}
		}
	}
	prefix = string(prefixarray)
	return prefix
}
