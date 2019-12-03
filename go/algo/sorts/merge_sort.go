package sorts

func MergeSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	i, j, k := start, mid+1, 0

	for ; i <= mid && j <= end; k++ {
		if arr[i] > arr[j] {
			tmp[k] = arr[j]
			j++
		} else {
			tmp[k] = arr[i]
			i++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = arr[i]
		k++
	}

	for ; j <= end; j++ {
		tmp[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmp)
}
