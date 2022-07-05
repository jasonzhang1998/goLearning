package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(888))
	fmt.Println(isPalindrome2(888))
}

// 反转一半的数字，然后判断数字是否相等
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	if x < 10 {
		return true
	}

	cur := 0
	for cur < x {
		cur = cur*10 + x%10
		x = x / 10
	}

	// 这里cur如果比x大，则cur的最后一位可以舍去
	return x == cur || x == cur/10
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
