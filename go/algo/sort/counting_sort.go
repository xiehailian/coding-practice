package sort

func CountingSort(arr []int) {

	n := len(arr)
	if n <= 1 {
		return
	}

	max := getMax(arr)
	count := make([]int, max+1)

	for i := 0; i < n; i++ {
		count[arr[i]]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[count[arr[i]] - 1] = arr[i]
		count[arr[i]]--
	}

	copy(arr, tmp)

}

