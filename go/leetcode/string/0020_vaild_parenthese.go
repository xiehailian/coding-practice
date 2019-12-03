package string

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-parentheses

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	var arr []byte
	symbol := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			arr = append(arr, byte(v))
		} else {
			if len(arr) > 0 {
				if _, ok := symbol[byte(v)]; ok {
					if arr[len(arr)-1] != symbol[byte(v)] {
						return false
					}
					arr = arr[:len(arr)-1]
				}
			} else {
				return false
			}
		}

	}
	return len(arr) == 0
}
