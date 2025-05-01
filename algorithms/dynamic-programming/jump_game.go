package dynamicProgramming

// Dynamic Programming approach.
// https://leetcode.com/problems/jump-game/
func JumpGameDP(nums []int) bool {
	l := len(nums)

	if l == 0 {
		return false
	}

	goodCell := make([]bool, l)

	for i := l - 1; i >= 0; i-- {
		jumpLen := nums[i]

		// We can jump to the last index.
		if jumpLen+i >= l-1 {
			goodCell[i] = true
			continue
		}

		canReachEnd := false
		for j := i + jumpLen; j > i; j-- {
			if goodCell[j] {
				canReachEnd = true
				break
			}
		}
		goodCell[i] = canReachEnd
	}

	return goodCell[0]
}

// Greedy approach.
// https://leetcode.com/problems/jump-game/
func JumpGameGreedy(nums []int) bool {
	l := len(nums)

	if l == 0 {
		return false
	}

	minGoodIdx := l - 1
	for i := l - 1; i >= 0; i-- {
		jumpLen := nums[i]

		if jumpLen+i >= minGoodIdx {
			// Remember the min index of the element
			// from which we can definitely reach the end.
			minGoodIdx = i
		}
	}

	return minGoodIdx == 0
}
