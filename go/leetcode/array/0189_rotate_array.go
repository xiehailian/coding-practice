package array

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//说明:
//
//尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
//要求使用空间复杂度为 O(1) 的 原地 算法。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/rotate-array

func rotate(nums []int, k int) []int {

	if nums == nil || len(nums) <= 1 || k < 0 {
		return nums
	}

	n := len(nums)
	for i := 0; i < k; i++ {
		previous := nums[n-1]
		for j := 0; j < n; j++ {
			tmp := nums[j]
			nums[j] = previous
			previous = tmp
		}
	}
	return nums
}

func rotate2(nums []int, k int) []int {

	if nums == nil || len(nums) <= 1 || k < 0 {
		return nums
	}

	n := len(nums)
	k = k % n
	var count = 0
	for i := 0; count < n; i++ {
		cur := i
		prev := nums[i]
		for i != cur {
			next := (cur + k) % n
			tmp := nums[next]
			nums[next] = prev
			prev = tmp
			cur = next
			count++
		}
	}
	return nums
}

func rotate3(nums []int, k int) []int {
	if nums == nil || len(nums) <= 1 || k < 0 {
		return nums
	}

	n := len(nums)
	k = k % n
	nums = reverse(nums, 0, n-1)
	nums = reverse(nums, 0, k-1)
	nums = reverse(nums, k, n-1)
	return nums
}

func reverse(nums []int, start, end int) []int {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}
