package sortingPrimer

func MergeSort(arr []int) []int {
	arrLen := len(arr)
	if arrLen == 1 {
		return arr
	}

	mid := arrLen / 2
	left := MergeSort(arr[0:mid])
	right := MergeSort(arr[mid:arrLen])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	leftSize := len(left)
	rightSize := len(right)
	result := make([]int, leftSize+rightSize)

	idx := 0
	l, r := 0, 0
	for l < leftSize && r < rightSize {
		if left[l] <= right[r] {
			result[idx] = left[l]
			l++
		} else {
			result[idx] = right[r]
			r++
		}
		idx++
	}

	if l < leftSize {
		Copy(left, l, result, idx)
	} else {
		Copy(right, r, result, idx)
	}

	return result
}

func Copy(src []int, sIdx int, dest []int, dIdx int) {
	for sIdx < len(src) {
		dest[dIdx] = src[sIdx]
		dIdx++
		sIdx++
	}
}
