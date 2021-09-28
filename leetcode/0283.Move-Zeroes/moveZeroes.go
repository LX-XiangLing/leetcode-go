package _283_Move_Zeroes

/**
 * 题目链接: https://leetcode-cn.com/problems/move-zeroes/
 */
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
