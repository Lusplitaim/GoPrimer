package main

import (
	"fmt"
	sortingPrimer "main/algorithms/sort"
)

func main() {
	arr := []int{5, 2, -1, 3, 0, 9}

	sortedArr := sortingPrimer.MergeSort(arr)
	fmt.Println(sortedArr)
}
