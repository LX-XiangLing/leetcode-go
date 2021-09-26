package _278_First_Bad_Version

/**
 *
 *
 * 题目链接:https://leetcode-cn.com/problems/first-bad-version/
 *
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool { return false }

func firstBadVersion(n int) int {
	low, high := 1, n
	for low <= high {
		index := (high-low)/2 + low
		if isBadVersion(index) {
			high = index - 1
		} else {
			if isBadVersion(index + 1) {
				return index + 1
			} else {
				low = index + 1
			}
		}
	}
	return low
}
