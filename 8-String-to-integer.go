package leetcode_golang

import (
	"math"
)

/*
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

Update (2015-02-10):
The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.

spoilers alert... click to show requirements for atoi.

Requirements for atoi:
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.

Subscribe to see which companies asked this question.
*/

func myAtoi(str string) int {
	strLen := len(str)
	if strLen == 0 {
		return 0
	}
	var skip int
	for skip = 0; skip < strLen && str[skip] == ' '; skip++ {
	}

	str = str[skip:]
	if len(str) == 0 {
		return 0
	}

	flag := 1
	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		flag = -1
		str = str[1:]
	}
	if len(str) == 0 || str[0] == '+' || str[0] == '-' {
		return 0
	}

	var ans int64
	for i := 0; i < len(str); i++ {
		if str[i] <= '9' && str[i] >= '0' {
			ans = ans * 10 + int64(str[i] - '0')
			if ans > math.MaxInt32 + 10 {
				break
			}
		} else {
			break
		}
	}
	if flag == -1 {
		ans *= -1
	}

	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}

	return int(ans)
}
