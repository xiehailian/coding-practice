package sort

func SelectSort(arr []int)  {

	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i; j < n - i; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}