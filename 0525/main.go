package main

import (
	"math"
	"sort"
	"strconv"
)

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)

	for i := 0; i < n-indexDifference; i++ {
		for j := i + indexDifference; j < n; j++ {
			k := nums[i] - nums[j]
			if k >= valueDifference || k*-1 >= valueDifference {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func maxVowels(s string, k int) int {
	ans := 0
	m := make(map[byte]bool)
	m['a'], m['e'], m['i'], m['o'], m['u'] = true, true, true, true, true

	left := 0
	right := left + k - 1
	n := len(s)
	for i := range right + 1 {
		println(i)
		if m[s[i]] {
			ans++
		}
	}
	c := ans
	for l, r := 1, left+k; r < n; {
		if m[s[r]] {
			c++
		}
		if m[s[l-1]] {
			c--
		}
		ans = max(c, ans)
		l++
		r++
	}

	return ans
}

func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	n := len(s)
	ans := 0
	for i := 0; i <= n-k; i++ {
		t, _ := strconv.Atoi(s[i:i+k])
		if t == 0 {
			continue
		}
		if num % t == 0 {
			ans++
		}
	}

	return ans
}


func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	if k == 1 {
		return 0
	}
	n := len(nums)
	ans := math.MaxInt
	for left := 0; left <= n-k; left++ {
		right := left+k-1
		ans = min(ans, nums[right]-nums[left])
	}
	return ans
}


func duplicateNumbersXOR(nums []int) int {
	ans := 0

	r := [51]int{}

	for _, v := range nums {
		r[v]++
		if r[v] == 2 {
			ans ^= v
		}
	}

	return ans
}


func occurrencesOfElement(nums []int, queries []int, x int) []int {
	r := []int{}

	for i, v := range nums {
		if v == x {
			r = append(r, i)
		}
	}

	nr := len(r)
	ans := []int{}

	for _, v := range queries {
		if v > nr {
			ans = append(ans, -1)
		} else {
			ans = append(ans, r[v-1])
		}
	}

	return ans
}


func queryResults(limit int, queries [][]int) []int {
	// 球 - 色
	s := make(map[int]int)
	// 色 - 球
	c := make(map[int]int)

	ans := []int{}

	for _, v := range queries {
		if qs, ok := s[v[0]]; ok {
			if qs != v[1] {
				c[qs]--
				if c[qs] == 0 {
					delete(c, qs)
				}
				c[v[1]]++
			}
		} else {
			c[v[1]]++
		}
		s[v[0]] = v[1]
		ans = append(ans, len(c))
	}

	return ans
}


func getResults(queries [][]int) []bool {
	// 断位
	a := []int{0}
	// 前缀最大可放区间
	b := []int{}


	for _, v := range queries {
		if v[0]==1 {
			a = append(a, v[1])
			if a[len(a)-1] < a[len(a)-2] {
				sort.Ints(a)
			}
		}
	}
}



func main() {
	maxVowels("leetcode", 3)
}