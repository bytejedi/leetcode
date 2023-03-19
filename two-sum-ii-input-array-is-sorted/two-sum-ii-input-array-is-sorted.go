package two_sum_ii_input_array_is_sorted

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
