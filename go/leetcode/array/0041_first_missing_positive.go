package array

//给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
//
//示例 1:
//
//输入: [1,2,0]
//输出: 3
//示例 2:
//
//输入: [3,4,-1,1]
//输出: 2
//示例 3:
//
//输入: [7,8,9,11,12]
//输出: 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/first-missing-positive

func firstMissingPositive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 1
	}

	n := len(nums)
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
