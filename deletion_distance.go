package main

import "fmt"

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func deleteDistance(str1 string, str2 string) int {
	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minimum(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Printf("Default print (%%v): %v\n", dp)
	return dp[m][n]
}
