package main

func numRescueBoats(people []int, limit int) int {
	ans := 0

	r := make([]int, limit)

	for _, v := range people {
		ans += v/limit
		r[v%limit]++
	}

	for i := limit-1; i >= 1; i-- {
		for r[i] > 0 {
			tmp := limit
			tmp -= i
			r[i]--
			j := i
			// 只能两个人
			for tmp > 0 && j>0 {
				if tmp >= j && r[j] > 0 {
					tmp -= j
					r[j]--
					break
				} else {
					j--
				}
			}
			ans++
		}
	}

	return ans
}

func main() {
	numRescueBoats([]int{3,2,3,2,2}, 6)
}