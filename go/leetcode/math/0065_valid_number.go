package math

import "strings"

//验证给定的字符串是否可以解释为十进制数字。
//
//例如:
//
//"0" => true
//" 0.1 " => true
//"abc" => false
//"1 a" => false
//"2e10" => true
//" -90e3   " => true
//" 1e" => false
//"e3" => false
//" 6e-1" => true
//" 99e2.5 " => false
//"53.5e93" => true
//" --6 " => false
//"-+3" => false
//"95a54e53" => false
//
//说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
//
//数字 0-9
//指数 - "e"
//正/负号 - "+"/"-"
//小数点 - "."
//当然，在输入中，这些字符的上下文也很重要。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/valid-number

var (
	blank  = 0 // 空格
	digit1 = 1 // 数字(0-9) 无前缀
	sign1  = 2 // +/- 无e前缀
	point  = 4 // '.'
	digit2 = 5 // 数字(0-9) 有符号前缀
	e      = 6 // 'e'
	sign2  = 7 // +/- 有e前缀
	digit3 = 8 // 数字(0-9) 有e前缀
)

func isNumber(s string) bool {
	s = strings.TrimRight(s, " ")
	dfa := [][]int{
		{blank, digit1, sign1, point, -1},
		{-1, digit1, -1, digit2, e},
		{-1, digit1, -1, point, -1},
		{-1, digit2, -1, -1, e},
		{-1, digit2, -1, -1, -1},
		{-1, digit2, -1, -1, e},
		{-1, digit3, sign2, -1, -1},
		{-1, digit3, -1, -1, -1},
		{-1, digit3, -1, -1, -1},
	}

	state := 0 // blank start
	for i := 0; i < len(s); i++ {
		var newState int
		switch s[i] {
		case ' ':
			newState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newState = 1
		case '+', '-':
			newState = 2
		case '.':
			newState = 3
		case 'e':
			newState = 4
		default:
			return false
		}
		state = dfa[state][newState]
		if state == -1 {
			return false
		}
	}
	return state == digit1 || state == digit2 || state == digit3
}

