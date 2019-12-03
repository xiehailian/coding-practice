package array

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/sliding-window-maximum

func maxSlidingWindow(nums []int, k int) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	var res = make([]int, 0)
	for l := 0; l <= len(nums)-k; l++ {
		r := l + k
		var max = 0
		for i := l; i < r; i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		res = append(res, max)
	}
	return res
}

func maxSlidingWindow2(nums []int, k int) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	window := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k+1)

	for i, v := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		for len(window) > 0 && nums[window[len(window)-1]] < v {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
