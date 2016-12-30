package leetcode_golang

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := -1
	ans := 0
	for i, vi := range s {
		if j, ok := m[vi]; ok {
			if j > start {
				start = j
			}
		}
		if i - start > ans {
			ans = i - start
		}
		m[vi] = i
	}
	return ans
}