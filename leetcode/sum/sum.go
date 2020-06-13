package sum

// simple

// 两个for 时：O(N^2)  空间复杂度： O(1）
func TwoSum1(nums []int, target int) (s []int) {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			m := nums[j]
			if num+m == target {
				s = []int{i, j}
				return
			}
		}
	}
	return
}

// map  时间：O(N) 空间：O(N)
func TwoSum2(nums []int, target int) (s []int) {
	m := map[int]int{}
	for i, num := range nums {
		if n, ok := m[target-num]; ok {
			s = []int{n, i}
			return
		}
		m[num] = i
	}
	return
}
