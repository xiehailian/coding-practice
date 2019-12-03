package sorts

func InsertSort(arr []int) {

	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
