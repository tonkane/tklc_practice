package main


func findPeaks(mountain []int) []int {
	ans := []int{}
	n := len(mountain)
	for i := 1; i < n-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return ans
}


func maxSum(nums []int, m int, k int) int64 {
	s := map[int]int{}
	ss := map[int]int{}

	ans := int64(0)
	sums := int64(0)
	for i := range k {
		s[nums[i]]++

		sums += int64(nums[i])
	}

	for j := range s {
		if s[j] > 1 {
			ss[j] = s[j]
			delete(s, j)
		}
	}
	if len(s)+len(ss) >= m {
		ans = sums
	}

	for i := 1; i <= len(nums)-k; i++ {
		sums += int64(nums[i+k-1])
		sums -= int64(nums[i-1])

		if _,ok := ss[nums[i+k-1]]; ok {
			ss[nums[i+k-1]]++
		} else {
			s[nums[i+k-1]]++
			if s[nums[i+k-1]] > 1 {
				ss[nums[i+k-1]] = s[nums[i+k-1]]
				delete(s, nums[i+k-1])
			}
		}

		if _,ok := s[nums[i-1]]; ok {
			s[nums[i-1]]--
			if s[nums[i-1]] == 0 {
				delete(s, nums[i-1])
			}
		} else {
			ss[nums[i-1]]--
			if ss[nums[i-1]] == 1 {
				delete(ss, nums[i-1])
				s[nums[i-1]] = 1
			}
		}

		if len(s)+len(ss) >= m {
			ans = max(ans, sums)
		}
	}

	return ans
}



func maximumSubarraySum(nums []int, k int) int64 {
	ans := int64(0)
	sums := int64(0)
	n := len(nums)
	cnt := map[int]int{}

	left := 0
	// right := left+k
	for left <= n-k {
		right := left+k-1
		if len(cnt) == 0 {
			for i := left; i <= right; i++ {
				if v, ok := cnt[nums[i]]; ok {
					sums = 0
					left = v
					cnt = map[int]int{}
					break
				} else {
					cnt[nums[i]] = i
				}
				sums += int64(nums[i])
			}
		} else {
			if v, ok := cnt[nums[right]]; ok {
				sums = 0
				left = v
				cnt = map[int]int{}
			} else {
				sums += (int64(nums[right]) - int64(nums[left-1]))
				cnt[nums[right]] = right
				delete(cnt, nums[left])
			}
		}
		left++
		ans = max(sums, ans)
	}


	return ans
}



func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	// 使k最大即是使 n-k 个和最小
	f := n-k
	sums := 0
	
	for i := range f {
		sums += cardPoints[i]
	}
	total := sums
	ans := sums

	left := 1
	for left <= k {
		right := left+f-1
		total += cardPoints[right]
		sums += cardPoints[right]
		sums -= cardPoints[left-1]
		ans = min(ans, sums)
		left++
	}

	return total-ans
}


func longestSemiRepetitiveSubstring(s string) int {
	// 记录重复出现的字符串的下标
	// 再次重复后，将left移动至出现下标位置处
	n := len(s)

	left := 0
	right := 1
	ans := 1
	cf := 0
	for right < n {
		if s[right] == s[right-1] {
			if cf != 0 {
				left = cf
			}
			cf = right
		}
		ans = max(ans, right-left+1)
		right++
	}

	return ans
}

func main() {
	maxScore([]int{1,79,80,1,1,1,200,1},3)
}