package stringAlgos

import "fmt"

func isPalindrome(x int) bool {
	str := fmt.Sprint(x)

	strLen := len(str)
	sIdx := 0
	eIdx := strLen - 1

	fmt.Println(strLen)

	for sIdx < strLen && eIdx >= 0 && sIdx <= eIdx {
		fmt.Printf("str[%d] = %d, str[%d] = %d\n", sIdx, str[sIdx], eIdx, str[eIdx])
		if str[sIdx] != str[eIdx] {
			return false
		}
		sIdx++
		eIdx--
	}

	return true
}
