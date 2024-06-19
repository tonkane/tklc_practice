package main

import (
	// "sort"
	"sort"
	"strconv"
)

func numberOfPairs2(nums1 []int, nums2 []int, k int) int {
	n := len(nums1)
	m := len(nums2)

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i]%(nums2[j]*k) == 0 {
				ans++
			}
		}
	}
	return ans
}

func compressedString(word string) string {
	ans := ""

	pre := ' '
	count := 0
	for _, v := range word {
		if v == pre && count != 9 {
			count++
		} else {
			if pre != ' ' {
				ans += strconv.Itoa(count) + string(pre)
			}
			pre = v
			count = 1
		}
	}

	ans += strconv.Itoa(count) + string(pre)

	return ans
}


func numberOfPairs(nums1, nums2 []int, k int) (ans int64) {
	cnt := map[int]int{}
	for _, x := range nums1 {
		if x%k > 0 {
			continue
		}
		x /= k
		// 因数分解 x  = q*q
		// 那么只需要检查到 q 即可
		// 大于q时，必然有一个小于q的数，如果整除，将 x / q 加入统计即可
		for d := 1; d*d <= x; d++ {
			if x%d == 0 {
				cnt[d]++
				if d*d < x {
					cnt[x/d]++
				}
			}
		}
	}
	for _, x := range nums2 {
		ans += int64(cnt[x])
	}
	return
}


func kthLargestValue(matrix [][]int, k int) int {

	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	list := []int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j+1] = dp[i][j+1] ^ dp[i+1][j] ^ dp[i][j] ^ matrix[i][j]
			list = append(list, dp[i+1][j+1])
		}
	}

	sort.Ints(list)

	return list[m*n-k]
}


func main() {
	kthLargestValue([][]int{{5,2},{1,6}}, 4)
}