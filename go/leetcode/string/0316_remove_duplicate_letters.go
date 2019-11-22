package string

//给定一个仅包含小写字母的字符串，去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
//示例 1:
//
//输入: "bcabc"
//输出: "abc"
//示例 2:
//
//输入: "cbacdcbc"
//输出: "acdb"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-duplicate-letters

func removeDuplicateLetters(s string) string {
	var stack []byte
	var counts [26]int

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	var insert [26]bool
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']--
		if insert[s[i]-'a'] {
			  continue
		}
		for j := len(stack)-1; j >= 0; j-- {
			if !(s[i] < stack[j] && counts[stack[j]-'a'] > 0) {
				break
			}
			insert[stack[j] - 'a'] = false
			stack = stack[:j]
		}
		insert[s[i]-'a'] = true
		stack = append(stack, s[i])
	}
	return string(stack)
}