package sorts

func ShellSort(arr []int) {

	n := len(arr)
	gap := n / 2

	for gap >= 1 {
		for i := 0; i < gap; i++ {
			for j := i + gap; j < n; j++ {
				if arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}
			}
		}
		gap--
	}
}
