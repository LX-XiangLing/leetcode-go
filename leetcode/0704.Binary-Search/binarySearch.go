package _704_Binary_Search

/**
 * 题目链接:https://leetcode-cn.com/problems/binary-search/
 */

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		index := (high-low)/2 + low
		num := nums[index]
		if num == target {
			return index
		}
		if num > target {
			high = index - 1
		}
		if num < target {
			low = index + 1
		}
	}
	return -1
}
