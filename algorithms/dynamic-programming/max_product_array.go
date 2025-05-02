package dynamicProgramming

// https://www.geeksforgeeks.org/maximum-product-subarray-in-c/
func MaxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxCurrent := nums[0]
	minCurrent := nums[0]
	maxTotal := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		temp := max(max(num, num*maxCurrent), num*minCurrent)
		minCurrent = min(min(num, num*maxCurrent), num*minCurrent)
		maxCurrent = temp
		maxTotal = max(maxTotal, maxCurrent)
	}

	return maxTotal
}
