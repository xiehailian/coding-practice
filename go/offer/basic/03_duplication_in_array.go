package basic

// 面试题3（一）：找出数组中重复的数字
// 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
// 那么对应的输出是重复的数字2或者3。

func DuplicationInArray(arr []int) int {

	if arr == nil || len(arr) <= 0 {
		return -1
	}

	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > n - 1 || arr[i] < 0 {
			return -1
		}
	}

	for i := 0; i < n; i++ {
		for arr[i] != i {
			if arr[i] == arr[arr[i]] {
				return arr[i]
			} else {
				arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
			}
		}
	}
	return -1
}

// 面试题3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。

func DuplicationInArrayNoEdit(arr []int) int{
	if arr == nil || len(arr) <= 0 {
		return -1
	}

	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > n || arr[i] < 1 {
			return -1
		}
	}

	left := 1
	right := n - 1

	for left < right {
		mid := (left + right) >> 1
		count := 0
		for _, value := range(arr) {
			if value <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}