package _005_Longest_Palindromic_Substring

/*
	题目链接:https://leetcode-cn.com/problems/longest-palindromic-substring/
	给你一个字符串 s，找到 s 中最长的回文子串。
	示例 1：
	输入：s = "babad"
	输出："bab"
	解释："aba" 同样是符合题意的答案。

	示例 2：
	输入：s = "cbbd"
	输出："bb"

	示例 3：
	输入：s = "a"
	输出："a"

	示例 4：
	输入：s = "ac"
	输出："a"
*/

// longestPalindrome 中心扩散
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

// 动态规划
// case1：自己本身肯定是回文
// case2：长度为2的字符串首尾相同就是回文
// case3：长度>2的字符串，首尾相同，去掉首尾是回文的话就是回文 状态转移方程 dp[start][end]=dp[start+1][end-1]
func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	res := s[0:1]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		// 自己本身肯定是回文
		dp[i][i] = true
	}
	for length := 2; length <= len(s); length++ {
		for start := 0; start < len(s)-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] { //首尾不同则不可能为回文
				continue
			} else if length < 3 {
				// 偶数回文情况
				dp[start][end] = true //即 case 2的判断
			} else {
				dp[start][end] = dp[start+1][end-1] //状态转移
			}
			if dp[start][end] && (end-start+1) > len(res) { //记录最大值
				res = s[start : end+1]
			}
		}
	}
	return res
}
