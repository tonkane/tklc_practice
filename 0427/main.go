package main

import (
	"math"
	// "math/bits"
)

func findColumnWidth(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	// var countNumsLen func(nums int) int
	countNumsLen := func(nums int) int {
		ans := 0
		if nums <= 0 {
			nums *= -1
			ans++
		}
		for nums > 0 {
			nums /= 10
			ans++
		}
		return ans
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		rMin := math.MaxInt
		rMax := math.MinInt
		for j := 0; j < m; j++ {
			rMin = min(rMin, grid[j][i])
			rMax = max(rMax, grid[j][i])
		}

		ans = append(ans, max(countNumsLen(rMin), countNumsLen(rMax)))
	}
	return ans
}

func countNumsLen(nums int) int {
	ans := 0
	if nums <= 0 {
		nums *= -1
		ans++
	}
	for nums > 0 {
		nums /= 10
		ans++
	}
	return ans
}

func jump(nums []int) int {
	leng := len(nums)
	dp := make([]int, leng)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && (i+j) < len(nums); j++ {
			if dp[i+j] == 0 {
				dp[i+j] += (dp[i] + 1)
			}
		}
	}
	return dp[leng-1]
}

// n,m 的和 = m,s - n,s
// 所以任意区间的和只需要 0-len 计算一遍
// 最大和是差值最大的两个数
func maxSubArray(nums []int) int {
	// sums := nums[0]
	ans := nums[0]
	sList := make([]int, len(nums))
	sList[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// sums += nums[i]
		sList[i] = sList[i-1] + nums[i]
	}
	curMin := nums[0]
	for j := 1; j < len(sList); j++ {
		ans = max(max(sList[j]-curMin, ans), sList[j])
		if curMin > sList[j] {
			curMin = sList[j]
		} 
	}
	return ans
}

func canJump(nums []int) bool {
	maxPosition := 0
	leng := len(nums)
	// if leng == 1 {
	// 	return true
	// }
	for i := 0; i < leng; i++ {
		if i > maxPosition {
			break
		}
		maxPosition = max(maxPosition, i+nums[i])
	}
	return maxPosition >= leng-1
}


// 如果偶位是正数，奇位是负数
// 至多30位
func baseNeg2(n int) string {
	ans := ""
	ansList := [32]int{}
	count := 0
	if n == 0 {
		return "0"
	}
	for n > 0 {
		ansList[count] = n &1
		n = n >> 1
		count++
	}
	for i := 0; i < 32; i++ {
		// 奇数
		if i & 1 == 1 {
			if ansList[i] == 1 {
				ansList[i+1] += 1
			}
			if ansList[i] == 2 {
				ansList[i] = 0
				ansList[i+1] += 1
			}
		}
		// 偶数被前面的借了
		if i & 1 == 0 && ansList[i] == 2 {
			ansList[i] = 0
			ansList[i+1] += 1
		}
	}
	firstOne := -1
	for i := 31; i >= 0; i-- {
		if firstOne == -1 && ansList[i] == 1 {
			firstOne = i
		}
		if i <= firstOne {
			if ansList[i] == 1 {
				ans += "1"
			} else {
				ans += "0"
			}
		}
	}
	return ans
}

func main() {
	var a int 
	a = '5'-'0'
	println(a)
}
