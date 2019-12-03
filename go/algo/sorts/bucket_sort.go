package sorts

func BucketSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	max := getMax(arr)
	buckets := make([][]int, n)

	var index int
	for i := 0; i < n; i++ {
		index = arr[i] * (n - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}

	cur := 0
	for i := 0; i < n; i++ {
		length := len(buckets[i])
		if length > 0 {
			QuickSort(buckets[i])
			copy(arr[cur:], buckets[i])
			cur += length
		}
	}
}

func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
