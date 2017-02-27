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

func Test19(t *testing.T) {

}

func Test35(t *testing.T) {
	in1 := [][]int {{1,3,5,6}, {1,3,5,6}, {1,3,5,6}, {1,3,5,6}}
	in2 := []int {5,2,7,0}
	out := []int {2,1,4,0}
	for i, v := range in1 {
		assert.Equal(t, out[i], searchInsert(v, in2[i]))
	}
}


func Test21(t *testing.T) {
	l1 := &ListNode{2, nil}
	l2 := &ListNode{1, nil}

	println(mergeTwoLists(l1, l2))


}