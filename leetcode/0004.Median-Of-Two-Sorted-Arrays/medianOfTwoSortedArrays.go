package _004_Median_Of_Two_Sorted_Arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	l := m + n
	left, right := -1, -1
	as, bs := 0, 0
	for i := 0; i <= l/2; i++ {
		left = right
		if as < m && (bs >= n || nums1[as] < nums2[bs]) {
			right = nums1[as]
			as++
		} else {
			right = nums2[bs]
			bs++
		}
	}
	if l&1 == 0 {
		return float64(left+right) / 2.0
	}
	return float64(right)
}
