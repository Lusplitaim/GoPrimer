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

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}

	elForDeletion := head
	prevEl := head
	iteration := 1

	for el := head; el != nil; el, iteration = el.Next, iteration+1 {
		if iteration > n {
			prevEl = elForDeletion
			elForDeletion = elForDeletion.Next
		}
	}

	if prevEl == elForDeletion {
		head = prevEl.Next
	} else if prevEl.Next == nil {
		prevEl = nil
	} else {
		prevEl.Next = elForDeletion.Next
	}

	return head
}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
func MaxStockProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	result := 0

	localMin := prices[0]
	localMax := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] >= localMax {
			localMax = prices[i]
		} else {
			result += localMax - localMin
			localMin = prices[i]
			localMax = prices[i]
		}
	}

	result += localMax - localMin

	return result
}
