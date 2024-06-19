package main

import (
	"sort"
	"strings"
)

func average(salary []int) float64 {
	sums := 0
	count := len(salary)
	a, b := 0, 1000000
	for i := 0; i < count; i++ {
		sums += salary[i]
		a = max(a, salary[i])
		b = min(b, salary[i])
	}

	return float64(sums-a-b) / float64(count-2)
}

func rotate(nums []int, k int) {
	// 每K个位置和首位交换
	// 交换 count-1 次
	count := len(nums)
	idx := 0
	start := 0
	for i := 0; i < count; i++ {
		idx = idx + k
		// for idx >= count {
		// 	idx -= count
		// }
		idx %= count
		if start == idx {
			start++
			idx = start
			continue
		}
		nums[start], nums[idx] = nums[idx], nums[start]
	}

	return
}

func maxProfit(prices []int) int {
	count := len(prices)
	// 状态转移存在两种状态
	// 持有股票和没有股票的状态
	dp := make([][2]int, count)
	dp[0][1] = -prices[0]
	for i := 1; i < count; i++ {
		// 没有股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	// 最后一定是没有持有股票
	return dp[count-1][0]
}

func hIndex(citations []int) int {
	ans := 0
	sort.Ints(citations)
	count := len(citations)
	for i := 0; i < count; i++ {
		if citations[i]+i > count {
			ans = max(ans, count-i)
			break
		}
		if citations[i]>0 && citations[citations[i]+i-1] >= citations[i] {
			ans = citations[i]
		}
	}

	return ans
}



func productExceptSelf(nums []int) []int {
	count := len(nums)
	ans := make([]int, count)
	for i := 0; i < count; i++ {
		ans[i] = 1
	}

	qj := 1
	for i := 1; i<count; i++ {
		qj *= nums[i-1]
		ans[i] = qj
	}

	qj = 1
	for i := count-2; i >=0; i-- {
		qj *= nums[i+1]
		ans[i] *= qj
	}

	return ans
}


func canCompleteCircuit2(gas []int, cost []int) int {
	count := len(gas)
	r := make([][2]int, count)

	a := 0
	for i := 0; i < count; i++ {
		r[i][0] = gas[i] - cost[i]
		r[i][1] = i
		a += r[i][0]
	}

	if a<0 {
		return -1
	}
	rc := make([][2]int, count)
	copy(rc, r)
	sort.SliceStable(r, func(i, j int) bool {
		if r[i][0] > r[j][0] {
			return true
		}
		if r[i][0] == r[j][0] {
			return r[i][1] < r[j][1]
		}
		return false
	})
	ans := -1
	for _, v := range r {
		costs := 0
		for i := 0; i < count; i++ {
			idx := (v[1]+i) % count
			costs += rc[idx][0]
			if costs < 0 {
				break
			}
		}
		if costs >= 0 {
			ans = v[1]
			break
		}
	}

	return ans
}

func canCompleteCircuit(gas []int, cost []int) int {
	count := len(gas)

	for i := 0; i < count; {
		sumsGas, sumsCost, idx := 0,0,0

		for idx < count {
			j := (idx+i) % count
			sumsGas += gas[j]
			sumsCost += cost[j]
			if sumsCost > sumsGas {
				break
			}
			idx++
		}

		if idx == count {
			return idx
		} else {
			i += idx+1
		}
	}
	return -1
}


func reverseWords(s string) string {
	r := strings.Split(s, " ")
	count := len(r)
	ans := ""
	for i := count-1; i >= 0; i-- {
		if r[i] == "" {
			continue
		}
		ans += r[i]+" "
	}
	return ans[:len(ans)-1]
}


func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sums := numbers[left] + numbers[right]
		if  sums == target {
			return []int{left+1, right+1}
		}
		if sums > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}



func minSubArrayLen(target int, nums []int) int {
	count := len(nums)
	sums := 0
	left, right := 0,0
	ans := count
	for right < count {
		if nums[right] >= target {
			return 1
		}
		sums += nums[right]
		if sums >= target {
			for left < right {
				if sums-nums[left] >= target {
					sums -= nums[left]
					left++
				} else {
					break
				}
			}
			ans = min(ans, right-left+1)
		}
		right++
	}
	if sums < target {
		return 0
	} else {
		return ans
	}
}


func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])
	mr := make([]bool, m)
	nr := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				mr[i] = true
				nr[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if mr[i] {
			for jj := 0; jj < n; jj++ {
				matrix[i][jj] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if nr[j] {
			for ii := 0; ii < m; ii++ {
				matrix[ii][j] = 0
			}
		}
	}
}


func gameOfLife(board [][]int)  {
	m := len(board)
	n := len(board[0])

	r := make([]int, n)
	pre := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sums := 0
			if i>=1 {
				if j>=1 {
					sums+=r[j-1]
				}
				sums+=r[j]
				if j<n-1 {
					sums+=r[j+1]
				}
			}
			if i<m-1 {
				if j>=1 {
					sums+=board[i+1][j-1]
				}
				sums+=board[i+1][j]
				if j<n-1 {
					sums+=board[i+1][j+1]
				}	
			}
			if j>=1 {
				sums+=pre
			}
			if j<n-1 {
				sums+=board[i][j+1]
			}
			if j >=1 {
				r[j-1] = pre
			}
			if j == n-1 {
				r[j] = board[i][j]
			}
			pre = board[i][j]

			if sums == 3 {
				board[i][j] = 1
			}
			if sums < 2 || sums > 3 {
				board[i][j] = 0
			}
		}
	}
}


func main() {
	minSubArrayLen(7, []int{2,3,1,2,4,3})
}