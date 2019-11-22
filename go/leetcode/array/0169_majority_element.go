package array

//给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/majority-element
//

func majorityElement(nums []int) int {
	return recursion(nums, 0, len(nums)-1)
}

func recursion(nums []int, low, high int) int {
	if low == high {
		return nums[low]
	}

	mid := low + (high - low) / 2
	left := recursion(nums, low, mid)
	right := recursion(nums, mid+1, high)

	if left == right {
		return left
	}

	var leftCount, rightCount = 0, 0
	for i := low; i <= high; i++ {
		if nums[i] == left {
			leftCount++
		}
		if nums[i] == right {
			rightCount ++
		}
	}

	if leftCount > rightCount {
		return left
	} else {
		return right
	}
}


