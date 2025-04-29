package sortingPrimer

func InsertionSort(arr []int) []int {
	result := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		result[i] = arr[i]

		for j := i; j > 0; j-- {
			if result[j-1] > result[j] {
				result[j-1], result[j] = result[j], result[j-1]
			} else {
				break
			}
		}
	}

	return result
}
