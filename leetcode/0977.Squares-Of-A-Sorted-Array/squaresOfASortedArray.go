package _977_Squares_Of_A_Sorted_Array

func sortedSquares(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	ans := make([]int, n)
	for index := n - 1; index >= 0; index-- {
		if v, x := nums[left]*nums[left], nums[right]*nums[right]; v > x {
			ans[index] = v
			left++
		} else {
			ans[index] = x
			right--
		}
	}
	return ans
}
