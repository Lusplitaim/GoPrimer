package main

import (
	"fmt"
	dynamicProgramming "main/algorithms/dynamic-programming"
)

func main() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(dynamicProgramming.Rob(arr))
}
