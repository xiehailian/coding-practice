package sorts

import "math"

func RadixSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	buckets := make([][]int, 10)
	maxDigitLength := 1
	for i := 0; i < n; i++ {
		if getDigitLength(arr[i]) > maxDigitLength {
			maxDigitLength = getDigitLength(arr[i])
		}
	}

	for i := 0; i < maxDigitLength; i++ {
		for _, v := range arr {
			buckets[getDigit(v, i)] = append(buckets[getDigit(v, i)], v)
		}

		k := 0
		for j := 0; j < 10; j++ {
			if len(buckets[j]) > 0 {
				for _, v := range buckets[j] {
					arr[k] = v
					k++
				}
			}
			buckets[j] = nil
		}
	}
}

func getDigitLength(digit int) int {
	if digit == 0 {
		return 1
	} else {
		return int(math.Log10(float64(digit))) + 1
	}
}

func getDigit(num int, index int) int {
	return num / int(math.Pow10(index)) % 10
}
