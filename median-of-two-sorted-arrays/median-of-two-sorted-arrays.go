package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int

	l1 := len(nums1)
	l2 := len(nums2)
	var longI, shortI int
	var longNums, shortNums []int
	if l1 > l2 {
		longI = l1
		shortI = l2
		longNums = nums1
		shortNums = nums2
	} else {
		longI = l2
		shortI = l1
		longNums = nums2
		shortNums = nums1
	}
	i1, i2 := 0, 0

	for i1 < longI {
		if i2 < shortI {
			if longNums[i1] < shortNums[i2] {
				nums = append(nums, longNums[i1])
				i1++
			} else {
				nums = append(nums, shortNums[i2])
				i2++
			}
		} else {
			nums = append(nums, longNums[i1])
			i1++
		}
	}
	if i2 < shortI {
		nums = append(nums, shortNums[i2:]...)
	}

	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2-1]+nums[l/2]) / 2
	} else {
		return float64(nums[l/2])
	}
}
