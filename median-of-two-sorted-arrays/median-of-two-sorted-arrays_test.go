package median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	median := findMedianSortedArrays([]int{3}, []int{-2, -1})
	t.Log("median", median)
}
