package main

import (
	"slices"
	"sort"
)

func numberOfChild(n int, k int) int {
	r := k % (n - 1)
	c := k / (n - 1)

	ans := 0
	if c%2 == 0 {
		ans = r
	} else {
		ans = n - r - 1
	}
	return ans
}

func valueAfterKSeconds(n int, k int) int {
	if n == 1 {
		return n
	}
	ans := 1 // 自己是1
	r := make([]int, k+1)
	// 初始化
	for i := range r {
		r[i] = i
	}
	ans += r[k]
	mod := 1000000000 + 7
	for j := 0; j < n-2; j++ {
		for i, cs := 0, 0; i <= k; i++ {
			r[i] += cs
			cs = r[i] % mod
		}
		ans += r[k] % mod
	}
	return ans % mod
}

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	sorted := slices.Compact(rewardValues)
	// n := len(rewardValues)
	n := len(sorted)
	// 0-1 背包，选或者不选
	// dp[i][j] 从前i个数得到总奖励j
	// 不选 dp[i][j] = dp[i-1][j]
	// 选 dp[i][j] = dp[i-1][j-v]    j>=v && j-v<v  =>   v <= j < 2v
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sorted[len(sorted)-1]*2)
	}
	dp[0][0] = 1
	end := sorted[len(sorted)-1]*2
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 0; j < end; j++ {
			dp[i][j] = dp[i-1][j]
			if rewardValues[i-1] <= j && rewardValues[i-1]*2 > j {
				dp[i][j] |= dp[i-1][j-rewardValues[i-1]]
			}
			if dp[i][j] == 1 {
				ans = max(ans, j)
			}
		}
	}
	return ans
}


func maxCoins(nums []int) int {
	n := len(nums)
	// 首尾+1
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// 创建一个 +2 数组 防止越界
	r := make([]int, n+2)
	for i:=1; i <=n; i++ {
		r[i] = nums[i-1]
	}
	r[0], r[n+1] = 1, 1

	// 状态转移
	// dp[i][j] = dp[i][k] + dp[k][j] + nums[k]*nums[i]*nums[j]

	for i := n-1; i >=0; i-- {
		for j := i+2; j <= n+1; j++ {
			for k := i+1; k<j; k++ {
				sums := dp[i][k] + dp[k][j] + r[i]*r[k]*r[j]
				dp[i][j] = max(dp[i][j], sums)
			}
		}
	}

	return dp[0][n+1]
}

func main() {
	maxTotalReward([]int{1, 3})
}
