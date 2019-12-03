package string

//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
//
//示例 1:
//
//输入: "abc"
//输出: 3
//解释: 三个回文子串: "a", "b", "c".
//示例 2:
//
//输入: "aaa"
//输出: 6
//说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/palindromic-substrings

func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	dp := make([][]bool, len(s))
	count := 0

	for end := 0; end < len(s); end++ {
		for start := 0; start <= end; start++ {
			if start == end {
				dp[start] = make([]bool, len(s))
				dp[start][end] = true
				count++
			} else {
				if s[start] == s[end] && (end-start == 1 || dp[start+1][end-1]) {
					dp[start][end] = true
					count++
				}
			}
		}
	}
	return count
}
