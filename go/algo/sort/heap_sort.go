package sort

func HeapSort(arr []int) {

	n := len(arr)
	if n <= 1 {
		return
	}

	buildHeap(arr, n)

	for n >= 1 {
		arr[1], arr[n] = arr[n], arr[1]
		heapifyUpToDown(arr, 1, n-1)
		n--
	}
}

func buildHeap(arr []int, n int) {
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDown(arr, i, n)
	}
}

func heapifyUpToDown(arr []int, top, count int)  {
	//for i := top; i <= count/2; {
	//	maxIndex := i
	//	if arr[i] < arr[i*2] {
	//		maxIndex = i * 2
	//	}
	//}
}