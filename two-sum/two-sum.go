package two_sum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for lI, l := range nums {
		r := target - l
		if rI, ok := m[r]; ok {
			return []int{rI, lI}
		}
		m[l] = lI
	}
	return nil
}
