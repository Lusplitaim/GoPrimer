package dynamicProgramming

import (
	"fmt"
	"math"
)

func DPChange(money int, coins []int) int {
	minNumCoins := make([]int, money)
	minNumCoins[0] = 0

	for i := 1; i < len(minNumCoins); i++ {
		minNumCoins[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			coin := coins[j]
			if i >= coin {
				numCoins := minNumCoins[i-coin] + 1
				if numCoins < minNumCoins[i] {
					minNumCoins[i] = numCoins
				}
				fmt.Println("")
			}
		}
	}

	return minNumCoins[money-1]
}

// https://leetcode.com/problems/min-cost-climbing-stairs/
func GetMinCost(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}

	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	dp := make([]int, len(cost))

	for i := len(cost) - 1; i >= 0; i-- {
		dp[i] = cost[i] + min(getCost(dp, i+1), getCost(dp, i+2))
	}

	return min(dp[0], dp[1])
}

func getCost(arr []int, idx int) int {
	if idx >= len(arr) {
		return 0
	} else {
		return arr[idx]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
