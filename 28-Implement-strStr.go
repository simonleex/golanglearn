package leetcode_golang



func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i ++ {
		if i + len(needle) > len(haystack) {
			return -1
		}
		if haystack[i: i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
