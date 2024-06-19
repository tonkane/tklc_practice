package main


func minDays2(n int) int {
	dp := make([]int, 1001)
	dp[0] = 0

	for i := 1; i < 1001; i++ {
		// 一个橘子
		dp[i] = dp[i-1] + 1

		// 找 2 
		if i%2 == 0 && i/2>= 1 {
			dp[i] = min(dp[i], dp[i/2] + 1)
		}

		// 找 3
		if i/3 >= 1 {
			if i%3 == 0 {
				dp[i] = min(dp[i], dp[i/3] + 1)
			} else {
				y := i%3
				dp[i] = min(dp[i], dp[i-y] + y)
			}
		}
	}


	if n > 1000 {
		ans := 0
		// n -= ans
		for n > 1000 {
			// 2次2 和 1次3哪个下降快？
			r3 := n % 3
			r2 := n % 2

			if r3 <= r2 {
				// 3下降快
				ans += r3
				n = (n-r3)/3
				ans++
			} else {
				// 2下降快
				ans += r2
				n = (n-r2)/2
				ans++ 
			}
			
		}
		ans += dp[n]
		return ans
	}

	return dp[n]
}



func minDays(n int) int {
	m := map[int]int{}

	var dfs func(int) int
	dfs = func(i int) int {
		if i<=1 {
			return i
		}
		if v,ok := m[i]; ok {
			return v
		}
		res := min(dfs(i/2)+i%2, dfs(i/3)+i%3) + 1
		m[i] = res
		return res
	}
	return dfs(n)
}


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp[1][1] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if i==1 && j==1 {
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}



func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	ans := 0

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {

			if matrix[i-1][j-1] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			ans = max(ans, dp[i][j])
		}
	}

	return ans*ans
}


func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	memo := make([][]int, n)

	for i := range memo {
		memo[i] = make([]int, m)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i<0 {
			return j+1
		}
		if j<0 {
			return i+1
		}

		if memo[i][j] != 0 {
			return memo[i][j]
		}

		if word1[i] == word2[j] {
			res := dfs(i-1, j-1)
			memo[i][j] = res
			return res
		}
		res := min(dfs(i-1, j), dfs(i, j-1), dfs(i-1, j-1)) + 1
		memo[i][j] = res
		return res
	}

	return dfs(n-1, m-1)
}


func isInterleave2(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	o := len(s3)

	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if k == o && i == n && j == m {
			return true
		} else if k == o {
			return false
		}

		// if i>=n || j>=m {
		// 	return false
		// }

		res := false

		if i<n && s1[i] == s3[k] {
			res = res || dfs(i+1, j, k+1)
		}
		if j<m && s2[j] == s3[k] {
			res = res || dfs(i, j+1, k+1)
		}
		return res
	}
	return dfs(0, 0, 0)
}


func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	o := len(s3)

	if m+n != o {
		return false
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			k := i+j-1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[k])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[k])
			}
		}
	}

	return dp[n][m]
}


func main() {
	isInterleave("aabcc", "dbbca", "aadbbcbcac")
}