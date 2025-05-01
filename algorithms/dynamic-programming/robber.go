package dynamicProgramming

// https://leetcode.com/problems/house-robber/description/
func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	m := len(nums)
	maxStepValue := make([]int, m)
	maxStepValue[m-1] = nums[m-1]
	maxStepValue[m-2] = nums[m-2]

	maxValue := maxStepValue[m-1]
	for i := m - 3; i >= 0; i-- {
		maxStepValue[i] = nums[i] + maxValue
		maxValue = max(maxStepValue[i+1], maxValue)
	}

	return max(maxStepValue[0], maxValue)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
