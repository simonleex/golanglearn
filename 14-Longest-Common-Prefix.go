package leetcode_golang



func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := 0
	endloop :
	for {
		for _, v := range strs {
			if len(v) == i || v[i] != strs[0][i] {
				break endloop
			}
		}
		i ++
	}
	return strs[0][:i]
}
