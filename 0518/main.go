package main

import (
	// "strconv"
	"strconv"
	"strings"
	// "strings"
	"math"
)

// import "math"

func maxDivScore(nums []int, divisors []int) int {
	n := len(divisors)
	r := make([]int, n)
	ans := math.MaxInt
	tmax := -1
	// tmin := math.MaxInt
	for i, d := range divisors {
		for _, v := range nums {
			if v%d == 0 {
				r[i]++
			}
		}
		if r[i] > tmax {
			ans = i
			tmax = r[i]
		}
		if r[i] == tmax && divisors[i] < divisors[ans] {
			ans = i
		}
		// tmin = min(tmin, divisors[i])
	}

	// if ans == math.MaxInt {
	// 	return tmin
	// }

	return divisors[ans]
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	l1 := len(v1)
	l2 := len(v2)

	v1i := 0
	v2i := 0

	for v1i < l1 || v2i < l2 {
		num1, num2 := 0, 0 
		if v1i < l1 {
			num1,_ = strconv.Atoi(v1[v1i])
		}  
		if v2i < l2 {
			num2,_ = strconv.Atoi(v2[v2i])
		}
		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
		v1i++
		v2i++
	}
	return 0
}



func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	// dp[i][j] 的定义为 
	// A面对区间i,j 时相对B的净胜分
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for i := n-2; i>=0; i-- {
		for j := i+1; j < n; j++ {
			dp[i][j] = max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1])
		}
	}

	return dp[0][n-1] >= 0
}



func longestPalindrome(s string) int {
	n := len(s)
	m := make(map[byte]int)

	for i := range s {
		m[s[i]]++
	}
	ans := 0
	for _, v := range m {
		ans += (v/2)*2
	}
	if n > ans {
		ans++
	}
	return ans
}



func removeDuplicateLetters(s string) string {
	r := make([]int, 'z'+1)

	for _, v := range s {
		r[v]++
	}

	ans := []rune{}
	inAns := make([]bool, 'z'+1)
	for _, v := range s {
		r[v]--
		if inAns[v] {
			continue
		}
		for len(ans)>0 && ans[len(ans)-1] > v && r[ans[len(ans)-1]] > 0 {
			last := ans[len(ans)-1]
			ans = ans[:len(ans)-1]
			inAns[last] = false
		}
		ans = append(ans, v)
		inAns[v] = true
	}
	return string(ans)
}

func main() {
	predictTheWinner([]int{1,5,2})
	// maxDivScore([]int{31,91,47,6,37,62,72,42,85}, []int{95,10,8,43,21,63,26,45,23,69,16,99,92,5,97,69,33,44,8})
	// maxDivScore([]int{73, 13, 20, 6}, []int{56, 75, 83, 26, 24, 53, 56, 61})
}