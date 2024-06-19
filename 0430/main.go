package main

import (
	"math"
	"sort"
	"strings"
)

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	ans := 0
	for _, value := range hours {
		if value >= target {
			ans++
		}
	}
	return ans
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	idx := m + n - 1
	for j>=0 {
		if i >=0 && nums1[i] >= nums2[j] {
			nums1[idx] = nums1[i]
			nums1[i] = math.MinInt
			i--
		} else {
			nums1[idx] = nums2[j]
			// nums1[j] = math.MinInt
			j--
		}
		idx--
	}
}


func removeElement(nums []int, val int) int {
	count := len(nums)
	idx := count-1
	for i := 0; i <= idx; i++ {
		for  idx >= 0 && nums[idx] == val {
			idx--
		}
		if idx < i {
			break
		}
		if nums[i] == val && idx>=0 {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx--
		}
	}
	return idx+1
}


func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}


func removeDuplicates(nums []int) int {

	count := len(nums)
	idx := 1
	i := 1
	for i < count {
		if idx > 1 && nums[i] == nums[idx-2] {
			i++
			continue
		}
		nums[idx] = nums[i]
		idx++
		i++
	}
	return idx
}



func romanToInt(s string) int {
	count := len(s)
	// r := make([]string, count)
	ans := 0
	s += "I"
	for i := 0; i < count; i++ {
		tmp := string(s[i])
		next := string(s[i+1])
		// 无特殊情况
		if tmp == "M" {
			ans += 1000
		}
		if tmp == "D" {
			ans += 500
		}
		if tmp == "L" {
			ans += 50
		}
		if tmp == "V" {
			ans += 5
		}
		// 有特殊情况
		// hasNext := i+1 < count
		if tmp == "I" {
			if next == "V" {
				ans += 4
				i++
			} else if next == "X" {
				ans += 9
				i++
			} else {
				ans += 1
			}
		}

		if tmp == "X" {
			if next == "L" {
				ans += 40
				i++
			} else if next == "C" {
				ans += 90
				i++
			} else {
				ans += 10
			}
		}

		if tmp == "C" {
			if next == "D" {
				ans += 400
				i++
			} else if next == "M" {
				ans += 900
				i++
			} else {
				ans += 100
			}
		}
	}
	return ans
}



func lengthOfLastWord(s string) int {
	count := len(s)
	ans := 0
	for i := count-1; i >=0; i-- {
		if s[i] == 32 {
			if ans > 0 {
				break
			}
		} else {
			ans++
		}
	}
	return ans
}

func strStr(haystack string, needle string) int {
	return strings.Index(haystack ,needle)
}


func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	count := len(s)
	left := 0
	right := count - 1
	for  left < right {
		if !isWordOrNums(s[left]) {
			left++
			continue
		}
		if !isWordOrNums(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}

func isWordOrNums(r byte) bool {
	if (r>=97 && r<=122) || (r>=48 && r<=57) {
		return true
	} else {
		return false
	}
}



func isSubsequence(s string, t string) bool {
	countS := len(s)
	countT := len(t)
	if countS == 0 {
		return true
	}
	if countS > countT {
		return false
	}
	sIdx := 0
	tIdx := 0
	for tIdx < countT {
		if s[sIdx] == t[tIdx] {
			sIdx++
		}
		tIdx++
		if sIdx == countS {
			return true
		}
	}
	return false
}


func canConstruct(ransomNote string, magazine string) bool {
	r := [26]int{}

	for i := 0; i < len(magazine); i++ {
		r[magazine[i]-'a']++
	}

	for j := 0; j < len(ransomNote); j++ {
		r[ransomNote[j]-'a']--

		if (r[ransomNote[j]-'a']<0) {
			return false
		}
	}

	return true
}


func isIsomorphic(s string, t string) bool {
	st := make(map[byte]byte)
	ts := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		value, ok := st[s[i]]
		value2, ok2 := ts[t[i]]
		if !ok && !ok2 {
			st[s[i]] = t[i]
			ts[t[i]] = s[i]
		} else {
			if ok && value != t[i] {
				return false
			}
			if ok2 && value2 != s[i] {
				return false
			}
		}
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	ps := make(map[byte]string)
	sp := make(map[string]byte)

	l := strings.Split(s, " ")

	if len(pattern) != len(l) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		v1, ok1 := ps[pattern[i]]
		v2, ok2 := sp[l[i]]

		if !ok1 && !ok2 {
			ps[pattern[i]] = l[i]
			sp[l[i]] = pattern[i]
		} else {
			if ok1 && v1 != l[i] {
				return false
			}
			if ok2 && v2 != pattern[i] {
				return false
			}
		}
	}

	return true
}


func isHappy(n int) bool {
	
	for n>=10 {
		tmp := 0
		for n>0 {
			mod := n%10
			mod *= mod
			tmp += mod
			n /= 10
		}
		n = tmp
	}
	if n == 1 || n==7 {
		return true
	}
    return false
}


func main() {
	isHappy(19)
}