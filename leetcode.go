package leetcode_golang

// 6. ZigZag Conversion
func convert(s string, numRows int) string {
	// be careful with numRows equals to 1 or 2
	if numRows == 1 {
		return s
	}

	ans := []byte{}
	length := len(s)
	numsPerPhase := numRows*2 - 2
	phases := length / numsPerPhase
	if phases*numsPerPhase != length {
		phases++
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < phases; j++ {
			diff := j*numsPerPhase + i
			if diff >= length {
				continue
			}
			ans = append(ans, byte(s[diff]))
			if i == 0 || i == numRows-1 {
				continue
			}
			diff1 := numsPerPhase - i + j*numsPerPhase
			if diff1 >= length {
				continue
			}
			ans = append(ans, byte(s[diff1]))
		}
	}

	return string(ans)
}
