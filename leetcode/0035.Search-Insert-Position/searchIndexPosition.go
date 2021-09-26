package _035_Search_Insert_Position

/**
 * 题目链接:https://leetcode-cn.com/problems/search-insert-position/
 */

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		index := (high-low)/2 + low
		if nums[index] == target {
			return index
		}
		if nums[index] > target {
			high = index - 1
		}
		if nums[index] < target {
			low = index + 1
		}
	}
	return low
}
