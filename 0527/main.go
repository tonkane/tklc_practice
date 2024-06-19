package main


func missingRolls(rolls []int, mean int, n int) []int {
	sums := 0
	m := len(rolls)
	// t := m+n
	for _, v := range rolls {
		sums += v
	}
	s := mean*(m+n)-sums

	if s > n*6 || s < n*1 {
		return []int{}
	}

	avg := s/n
	r := s % n

	ans := []int{}
	
	for i := 0; i < n; i++ {
		num := avg
		if r > 0 {
			num += 1
		}
		ans = append(ans, num)
		r--
	}
	return ans

}


func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	r := make([]int, n+1)

	m := make(map[byte]bool)
	m['a'],m['e'],m['i'],m['o'],m['u'] = true,true,true,true,true

	for i, v := range words {
		if m[v[0]] && m[v[len(v)-1]] {
			r[i+1] = r[i]+1
		} else {
			r[i+1] = r[i]
		}
	}

	ans := []int{}

	for i := 0; i < len(queries); i++ {
		ans = append(ans, r[queries[i][1]+1] - r[queries[i][0]])
	}

	return ans
}


func numSubarraysWithSum2(nums []int, goal int) int {
	n := len(nums)
	r := make([]int, n+1)

	for i := 0; i < n; i++ {
		r[i+1] = r[i]+nums[i]
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i+1; j < n+1; j++ {
			if r[j] - r[i] == goal {
				ans++
			}
			if r[j] - r[i] > goal {
				break
			}
		}	
	}
	return ans
}


func numSubarraysWithSum(nums []int, goal int) (ans int) {
    left1, left2 := 0, 0
    sum1, sum2 := 0, 0
    for right, num := range nums {
        sum1 += num
        for left1 <= right && sum1 > goal {
            sum1 -= nums[left1]
            left1++
        }
        sum2 += num
        for left2 <= right && sum2 >= goal {
            sum2 -= nums[left2]
            left2++
        }
        ans += left2 - left1
    }
    return
}


func main() {
	numSubarraysWithSum([]int{1,0,1,0,1}, 2)
	// vowelStrings([]string{"aba","bcb","ece","aa","e"}, [][]int{{0,2},{1,4},{1,1}})
}