package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {

    var noneZeros []int
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            noneZeros = append(noneZeros, nums[i])
        }

    }

    for i := 0; i < len(nums) ; i++ {
        if i < len(noneZeros) {
            nums[i] = noneZeros[i]
        } else {
            nums[i] = 0
        }
    }
}

func lengthOfLongestSubstring(s string) int {
    var m = make(map[string]int)

    for _, j := range s {
        if _, ok := m[string(j)]; ok {
            continue
        } else {
            m[string(j)] = 0
        }
    }
    fmt.Println(m)
    return len(m)
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}


func main() {
	//lengthOfLongestSubstring("pwwkew")
	//fmt.Println(maxDepth(100))
	//str := "谢hehe"
	//fmt.Println(len([]rune(str)),strings.Count(str,""), len(str))
	//fmt.Println(string([]rune{122, 212, 333, 125}))

	s := []int{1}
	fmt.Println(s[1:], s[:0])

	//fmt.Println(strconv.Itoa(1))

	//fmt.Println(100 / 8, 100 % 8, 1 << 4, 1 & 1)
	//var b = make([]byte, 10)
	//b[8] |= 1 << 4
	//fmt.Println(unsafe.Sizeof(b[0]))

	fmt.Println(^uint32(0))
}
