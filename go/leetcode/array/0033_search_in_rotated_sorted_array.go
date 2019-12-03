package array

//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中不存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//示例 1:
//
//输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
//示例 2:
//
//输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	rotateIndex := findRotateIndex(nums)
	if nums[rotateIndex] > target {
		return bsearch(nums, target, 0, rotateIndex-1)
	} else if nums[rotateIndex] < target {
		return bsearch(nums, target, rotateIndex+1, len(nums)-1)
	} else {
		return rotateIndex
	}
	return -1
}

func bsearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func findRotateIndex(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] > nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
