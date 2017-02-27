package leetcode_golang



func searchInsert(nums []int, target int) int {
	var l, r , mid int = 0, len(nums) - 1 , 0
	for l < r {
		mid = (l + r)/ 2
		if target < nums[mid] {
			r = mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	if target > nums[l] {
		l ++
	}
	return l
}
