package _001_Two_Sum

/**
题目链接:https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, v := range nums {
		if k, ok := sumMap[target-v]; ok {
			return []int{k, i}
		}
		sumMap[v] = i
	}
	return nil
}
