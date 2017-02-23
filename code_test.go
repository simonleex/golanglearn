package leetcode_golang

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test8(t *testing.T) {
	strs := []string{" -00123f1","+-2","+", " +-1234", "11", "", "-123","0", "  -123124123123"}
	anss := []int{-123,0,0, 0, 11, 0, -123,0, math.MinInt32}

	for i, s := range strs {
		assert.Equal(t, anss[i], myAtoi(s))
	}

}
