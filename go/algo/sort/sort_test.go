package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	BubbleSort(arr)
	fmt.Println("排序后：", arr)
}


func TestInsertSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	InsertSort(arr)
	fmt.Println("排序后：", arr)
}

func TestSelectSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	SelectSort(arr)
	fmt.Println("排序后：", arr)
}

func TestShellSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	ShellSort(arr)
	fmt.Println("排序后：", arr)
}

func TestMergeSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	MergeSort(arr)
	fmt.Println("排序后：", arr)
}

func TestQuickSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	QuickSort(arr)
	fmt.Println("排序后：", arr)
}

func TestBucketSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,10000}
	fmt.Println("排序前：", arr)
	BucketSort(arr)
	fmt.Println("排序后：", arr)
}

func TestCountingSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	CountingSort(arr)
	fmt.Println("排序后：", arr)
}

func TestRadixSort(t *testing.T) {
	arr := []int{9,4,5,2,7,2,1}
	fmt.Println("排序前：", arr)
	RadixSort(arr)
	fmt.Println("排序后：", arr)
}

//func TestHeapSort(t *testing.T) {
//	arr := []int{9,4,5,2,7,2,1}
//	fmt.Println("排序前：", arr)
//	HeapSort(arr)
//	fmt.Println("排序后：", arr)
//}