package string

//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//说明:
//
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
//示例 1:
//
//输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/regular-expression-matching

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	first := (len(s) != 0 && (p[0] == s[0] || p[0] == '.'))

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (first && isMatch(s[1:], p))
	} else {
		return first && isMatch(s[1:], p[1:])
	}
}

func isMatch2(s string, p string) bool {
	var dp = make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			first := i < len(s) && (p[j] == s[i] || p[j] == '.')
			if j + 1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (first && dp[i+1][j])
			} else {
				dp[i][j] = first && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}