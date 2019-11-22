package string


//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//
//换句话说，第一个字符串的排列之一是第二个字符串的子串。
//
//示例1:
//
//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
// 
//
//示例2:
//
//输入: s1= "ab" s2 = "eidboaoo"
//输出: False
// 
//
//注意：
//
//输入的字符串只包含小写字母
//两个字符串的长度都在 [1, 10,000] 之间
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/permutation-in-string

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1map := [26]int{}
	for i := 0; i < len(s1); i++ {
		s1map[s1[i] - 'a']++
	}
	for i := 0; i < len(s2) - len(s1); i++ {
		s2map := [26]int{}
		for j := 0; j < len(s1); j++ {
			s2map[s2[i+j] - 'a']++
		}

		if matches(s1map, s2map) {
			return true
		}
	}
	return false
}

func matches(s1map, s2map [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1map[i] != s2map[i] {
			return false
		}
	}
	return true
}


func checkInclusion2(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	s1map := [26]int{}
	s2map := [26]int{}
	for i := 0; i < len(s1); i++ {
		s1map[s1[i] - 'a']++
		s2map[s2[i] - 'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if s1map[i] == s2map[i] {
			count++
		}
	}

	for i := 0; i < len(s2) - len(s1); i++ {
		left := s2[i] - 'a'
		right := s2[i + len(s1)] - 'a'
		if count == 26 {
			return true
		}

		s2map[right]++
		if s2map[right] == s1map[right] {
			count++
		} else if s2map[right] == s1map[right] + 1 {
			count--
		}

		s2map[left]--
		if s2map[left] == s1map[left] {
			count++
		} else if s2map[left] == s1map[left] - 1 {
			count--
		}
	}
	return count == 26
}