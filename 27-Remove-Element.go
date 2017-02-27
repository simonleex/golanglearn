package leetcode_golang


func removeElement(nums []int, val int) int {
	var l, r int
	for r < len(nums) {
		if val != nums[r] {
			nums[l] = nums[r]
			l ++
			r ++
		} else {
			r ++
		}
	}
	nums = nums[:l]
	return len(nums)
}
