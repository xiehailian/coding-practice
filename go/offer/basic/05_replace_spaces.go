package basic

// 面试题5：替换空格
// 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
// 则输出“We%20are%20happy.”。

func ReplaceSpaces(str string)  string {
	strList := []rune(str)
	length := len(strList)
	count := 0

	for i := 0; i < length; i++ {
		if strList[i] == ' ' {
			count++
		}
	}

	newLength := length + count * 2
	newList := make([]rune, newLength)
	for i, j := length - 1, newLength - 1; i >= 0 && j >= 0; {
		if strList[i] == ' '{
			newList[j] = '0'
			j--
			newList[j] = '2'
			j--
			newList[j] = '%'
			j--
		} else {
			newList[j] = strList[i]
			j--
		}
		i--
	}
	return string(newList)
}