package leetcode_golang


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i,vi := range nums {
		if j, ok := m[vi]; ok {
			return []int{j, i}
		}
		m[target - vi] = i
	}
	return nil
}