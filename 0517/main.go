package main

import (
	"math"
	"slices"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	d2p := make([][]int, n)
	wn := len(worker)
	for i := range d2p {
		d2p[i] = []int{difficulty[i], profit[i]}
	}

	slices.SortFunc(d2p, func(a, b []int) int {
		return a[0] - b[0]
	})

	slices.Sort(worker)

	ans := 0

	d := 0
	w := 0

	// md := 0
	mp := 0
	for w < wn {
		// 说明能做工作
		if d < n && d2p[d][0] <= worker[w] {
			mp = max(mp, d2p[d][1])
			d++
		} else {
			ans += mp
			w++
		}
	}

	return ans
}



func findPermutationDifference(s string, t string) int {
	r := [26]int{}
	ans := 0
	for i, word := range s {
		r[word-'a'] = i
	}

	for i, word := range t {

		if i >= r[word-'a'] {
			ans += (i-r[word-'a'])
		} else {
			ans += (r[word-'a']-i)
		}

	}

	return ans
}


func maximumEnergy(energy []int, k int) int {
	n := len(energy)

	ans := math.MinInt

	c := k
	for i := n-1; c > 0; {
		tmp := 0
		
		for j :=i; j >= 0;  j-=k {
			tmp += energy[j]
			ans = max(ans, tmp)
		}
		c--
		i--
	}

	return ans
}



func maxScore2(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][][3]int, n)

	for i := range dp {
		dp[i] = make([][3]int, m)
	}

	dp[0][0] = [3]int{0, 0, 0}

	mm := grid[0][0]
	for i := 1; i < n; i++ {
		tmp := grid[i][0] - mm
		mm = min(grid[i][0], mm)
		dp[i][0] = [3]int{0, tmp, tmp}
	}

	mm = grid[0][0]
	for j := 1; j < m; j++ {
		tmp := grid[0][j] - mm
		mm = min(grid[0][j], mm)
		dp[0][j] = [3]int{tmp, 0, tmp}
	}

	ans := math.MinInt
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j][0] = dp[i][j-1][0] + (grid[i][j] - grid[i][j-1])
			dp[i][j][1] = dp[i-1][j][1] + (grid[i][j] - grid[i-1][j])

			dp[i][j][2] = max(dp[i][j][0], dp[i][j][1])

			for ii:=i-1; ii>=0; ii-- {
				dp[i][j][2] = max(dp[ii][j][0] + dp[i][j][1], dp[i][j][2])
			}

			for jj:=j-1; jj>=0; jj-- {
				dp[i][j][2] = max(dp[i][jj][1] + dp[i][j][0], dp[i][j][2])
			}

			

			ans = max(ans, dp[i][j][2])
		}
	}
	return ans
}


func maxScore(grid [][]int) int {
	ans := math.MinInt
	n := len(grid)
	m := len(grid[0])

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for j := range dp[0] {
		dp[0][j] = math.MaxInt
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = math.MaxInt
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mn := min(dp[i+1][j], dp[i][j+1])
			ans = max(ans, grid[i][j]-mn)
			dp[i+1][j+1] = min(mn, grid[i][j])
		}
	}

	return ans
}

func subsets(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	ans = append(ans, []int{})
	path := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if i >= n {
			return
		}
		for k := range nums {
			if k > i {
				path = append(path, nums[k])
				dfs(k)
				ans = append(ans, append([]int{}, path...))
				path = path[:len(path)-1]
			}
		}
	}
	dfs(-1)
	return ans
}


func partition(s string) [][]string {
	n := len(s)
	ans := [][]string{}
	path := []string{}
	var dfs func(s, e int)
	dfs = func(start, end int) {
		if end>n {
			return
		}
		if rv(s[start:end]) {
			path = append(path, s[start:end])
			if end == n {
				ans = append(ans, append([]string{}, path...))
			}
			dfs(end, end+1)
			path = path[:len(path)-1]
		}
		dfs(start, end+1)
	}
	dfs(0, 1)
	return ans
}


func rv(s string) bool {
	n := len(s)
	i,j:=0,n-1
	for i<j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}


func decodeString(s string) string {
	stack := []string{}
	n := len(s)
	ans := ""
	for i := 0; i < n; i++ {
		tmp := s[i]-'0'
		if tmp>=0 && tmp<=9 {
			tmps := string(tmp)
			for j:=i+1;j<n; {
				tmpx := s[j]-'0'
				if tmpx>=0 && tmpx<=9 {
					tmps += string(tmpx)
					i=j
					j++
				} else {
					break
				}
			}
			stack = append(stack, tmps)
		} else if s[i] == '[' {

		} else if s[i] == ']' {

		}else {
			ans += string(s[i])
		}
	}
	return ans
}


func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{}
	stack = append(stack, 0)
	for i := 1; i < n; i++ {
		if temperatures[i] > temperatures[i-1] {
			for len(stack) > 0 {
				last := stack[len(stack)-1]
				if temperatures[last] < temperatures[i] {
					ans[last] = i - last
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, i)
	}

	return ans
}


func main() {
	s := "bbab"
	println(rv(s))
}