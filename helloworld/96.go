package main

// 不同的二叉搜索树

import "fmt"

func main() {
	fmt.Println("different bst")
	fmt.Println(numTrees(5))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}

	return dp[n]
}
