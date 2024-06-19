package main

import "math"

func distributeCandies(candies int, num_people int) []int {
	// (1+n)*n/2
	// 项数
	n := int(math.Sqrt(float64(candies) * 2))

	if n*n+n < candies*2 {
		n++
	}

	// 层数
	c := n / num_people

	// 余数
	r := n % num_people

	ans := make([]int, num_people)

	sums := 0
	for i := 0; i < num_people; i++ {
		ans[i] = (i+1)*c + (c-1)*c/2*num_people
		if r > i {
			ans[i] += (i + 1) + c*num_people
		}
		sums += ans[i]
	}

	if r == 0 {
		ans[len(ans)-1] -= (sums - candies)
	} else {
		ans[r-1] -= (sums - candies)
	}

	return ans
}

func main() {
	distributeCandies(22, 4)
}
