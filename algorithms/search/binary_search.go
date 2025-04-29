package searchPrimer

func BinarySearch(arr []int, val, low, high int) (found bool, idx int) {
	if high < low {
		return false, 0
	}

	mid := low + (high-low)/2

	if arr[mid] == val {
		return true, mid
	}

	if arr[mid] < val {
		return BinarySearch(arr, val, mid+1, high)
	} else {
		return BinarySearch(arr, val, low, mid-1)
	}
}
