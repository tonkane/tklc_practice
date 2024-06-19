package main

import (
	"math"
	"sort"
	// "golang.org/x/text/search"
)

func isArraySpecial2(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if nums[i]%2 == 0 && nums[i-1]%2 == 1 {
			continue
		}
		if nums[i]%2 == 1 && nums[i-1]%2 == 0 {
			continue
		}
		return false
	}
	return true
}

func sumDigitDifferences(nums []int) int64 {
	// r := [10]int{}
	n := len(nums)
	l := 0
	t := nums[0]
	for t > 0 {
		t = t / 10
		l++
	}

	r := make([][10]int, l)

	for _, v := range nums {
		w := 0
		for v > 0 {
			r[w][v%10]++
			v = v / 10
			w++
		}
	}

	ans := int64(0)

	for i := 0; i < l; i++ {
		tmp := 0
		for j := 0; j < 10; j++ {
			if r[i][j] != 0 {
				ans += int64((n - r[i][j] - tmp) * r[i][j])
				tmp += r[i][j]
			}
		}
	}
	return ans
}

func helper(nums []int, start int) int {
	// 记录不是奇偶性的点
	// 然后二分查找
	n := len(nums)
	if n == 1 {
		return n
	}
	for i := start + 1; i < n; i++ {
		if nums[i]%2 == 0 && nums[i-1]%2 == 1 {
			continue
		}
		if nums[i]%2 == 1 && nums[i-1]%2 == 0 {
			continue
		}
		return i - 1
	}
	return n
}

func isArraySpecial(nums []int, queries [][]int) []bool {
	r := []int{}

	start, end := 0, len(nums)
	for start < end {
		ns := helper(nums, start)
		if ns < end {
			r = append(r, ns)
		}
		start = ns + 1
	}

	ans := []bool{}
	rn := len(r)

	for i := 0; i < len(queries); i++ {
		s := sort.Search(rn, func(j int) bool {
			return r[j] >= queries[i][0]
		})

		e := sort.Search(rn, func(j int) bool {
			return r[j] >= queries[i][1]
		})


		if e - s > 0 {
			ans = append(ans, false)
		} else {
			ans = append(ans, true)
		}

		println(s)
		// println(e)
	}

	return ans
}

func getWinner(arr []int, k int) int {
	win := 0
	tz := 1
	n := len(arr)
	count := 0
	for tz < n {
		if arr[win] > arr[tz] {
			count++
		} else {
			win = tz
			count = 1
		}
		if count >= k {
			break
		}
		tz++
	}

	return arr[win]
}



func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	dp[1] = 1

	for cur := 2; cur <=n; cur++ {
		for j:=1; j*j <=cur; j++ {
			if cur == j*j {
				dp[cur] = 1
				break
			} else {
				dp[cur] = min(dp[j*j]+dp[cur-j*j], dp[cur])
			}
		}
	}

	return dp[n]
}


func findLength2(nums1 []int, nums2 []int) int {
	ans := 0

	left1 := 0
	left2 := 0

	n1 := len(nums1)
	n2 := len(nums2)

	tans := 0
	for left1 < n1 && left2 < n2 {
		if nums1[left1] == nums2[left2] {
			right1 := left1
			right2 := left2
			for right1<n1 && right2<n2 {
				if nums1[right1] == nums2[right2] {
					right1++
					right2++
					tans++
				} else {
					break
				}
			}
			ans = max(tans, ans)
		} 
		tans = 0
		left2++
		if left1 < n1 && left2 == n2 {
			left2 = 0
			left1++
		}
	}

	return ans
}



func findLength(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	start1, end1 := 0, 0
	start2, end2 := n2-1, n2-1

	ans := 0
	for {
		tans := 0
		s1, s2 := start1, start2
		for s1 <= end1 && s2 <= end2 {
			if nums1[s1] == nums2[s2] {
				tans++
				ans = max(ans, tans)
			} else {
				tans = 0
			}
			s1++
			s2++
		}

		c1 := end1 - start1 + 1
		c2 := end2 - start2 + 1

		e1 := end1
		
		if c1 != n1 || start2 == 0 {
			if end1 < n1-1 {
				end1++
				if c1 >= n2 {
					start1++
				}
			} else if start1 < n1-1 {
				start1++
			}
			
		}

		if c2 != n2 || e1 == n1-1 {
			if start2 > 0 {
				start2--
				if c2 >= n1 {
					end2--
				}
			} else if end2 > 0 {
				end2--
			}
		}

		if start2 == 0 && end2 == 0 {
			break
		}
		
	}
	return ans
}

func main() {
	findLength([]int{1,2,3,2,1}, []int{1,3,2,1,5})
}