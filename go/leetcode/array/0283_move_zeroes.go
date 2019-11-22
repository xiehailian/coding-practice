package array

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/move-zeroes

func moveZeroes(nums []int)  {

	if nums == nil || len(nums) <= 1 {
		return
	}

	//var count = 0
	//for i := 0; i < len(nums); i++ {
	//	if nums[i] == 0 {
	//		for j := i; j < len(nums) - count; j++ {
	//			nums[j+1], nums[j] = nums[j], nums[j+1]
	//		}
	//	}
	//	count++
	//}

	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

}