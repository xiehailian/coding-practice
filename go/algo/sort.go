package algo

import "strconv"

func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n; i ++ {
		flag := false
		for j := 0; j < n - i - 1; j ++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}

func InsertionSort(a []int) {
	n := len(a)

	for i := 0; i < n; i++ {
		for j := i; j > 0; j -- {
			if a[j]  < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func SelectionSort(a []int)  {

	n := len(a)

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

func ShellSort(a []int)  {
	gap := len(a)/2
	for gap >= 1 {
		for i := 0; i < gap; i++ {
			for j := i+gap; j < len(a); j++ {
				if a[j] < a[j-gap] {
					a[j], a[j-gap] = a[j-gap], a[j]
				}
			}
		}
		gap--
	}
}

func MergeSort(a []int)  {
	n := len(a)
	if n <= 1 {
		return
	}

	mergeSort(a, 0, n-1)
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
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}

	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], tmpArr)
}

func QuickSort(a []int) {
	separateSort(a, 0, len(a)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}


func BucketSort(a []int)  {
	num := len(a)
	if num <= 1 {
		return
	}

	max := getMax(a)
	buckets := make([][]int, num)

	index := 0
	for i := 0; i < num; i++ {
		index = a[i] * (num - 1) / max
		buckets[index] = append(buckets[index], a[i])
	}

	cur := 0
	for i := 0; i < num; i++ {
		l := len(buckets[i])
		if l > 0 {
			QuickSort(buckets[i])
			copy(a[cur:], buckets[i])
			cur += l
		}
	}

}

func getMax(a []int) int  {
	max := a[0]
	for i:= 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}


func CountingSort(a []int)  {
	num := len(a)

	if num <= 1 {
		return
	}

	max := getMax(a)

	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	r := make([]int, num)
	for i := range a {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	copy(a, r)
}

func RadixSort(a []int)  {
	buckets := make([][]int, 10)
	maxDigist := maxDigist(a)

	for i := 0; i < maxDigist; i++ {
		for _, v := range a {
			buckets[getDigist(v, i)] = append(buckets[getDigist(v, i)], v)
		}

		j := 0
		for index, value := range buckets {
			if len(value) > 0 {
				for _, v := range value {
					a[j] = v
					j++
				}
			}
			buckets[index] = nil
		}
	}
}

func maxDigist(a []int) int {
	maxDigist := 1
	for _, v := range a {
		if len(strconv.Itoa(v)) > maxDigist {
			maxDigist = len(strconv.Itoa(v))
		}
	}
	return maxDigist
}

func getDigist(num int, index int) int {
	strNum := strconv.Itoa(num)
	if index > len(strNum) - 1 {
		return 0
	}
	index = len(strNum) - 1 - index
	res, err := strconv.Atoi(string(strNum[index]))
	if err != nil {
		return -1
	} else {
		return res
	}
}

func HeapSort(a []int)  {
	buidHeap(a, len(a))

	k := len(a)
	for k >= 1 {
		swap(a, 1, k)
		heapifyUpToDown(a, 1, k-1)
		k--
	}
}

//build a heap
func buidHeap(a []int, n int) {

	//heapify from the last parent node
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(a, i, n)
	}
}

//heapify from up to down , node index = top
func heapifyUpToDown(a []int, top int, count int) {

	for i := top; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}

}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
