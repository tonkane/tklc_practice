package main

// import "sort"

func numberOfWeeks(milestones []int) int64 {
	sums := int64(0)
	imax := 0
	for _, v := range milestones {
		sums += int64(v)
		imax = max(imax, v)
	}
	n2 := (sums+1)/2

	if int64(imax) > n2 {
		return (sums-int64(imax))*2+1
	}

	return sums

}


func maxProduct2(nums []int) int {
	n := len(nums)
	// dp := make([]int, n)

	p := make([]int, n)
	
	p[0] = nums[0]
	ans := p[0]
	if p[0] == 0 {
		p[0] = 1
	}
	for i := 1; i < n; i++ {
		if nums[i] != 0 {
			p[i] = p[i-1]*nums[i]
			ans = max(ans, p[i])
		} else {
			p[i] = 1
		}
		ans = max(ans, nums[i])
	}

	for i := 0; i < n; i++ {
		for j := i-1; j >= 0; j-- {
			if nums[j] == 0 {
				break
			}
			ans = max(ans, p[i]/p[j])
		}
	}

	return ans
}


func maxProduct(nums []int) int {
	imax, imin, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mx, mi := imax, imin
		imax = max(mx*nums[i], nums[i], mi*nums[i])
		imin = min(mi*nums[i], nums[i], mx*nums[i])

		ans = max(imax, ans)
	}

	return ans
}


func main() {
	maxProduct([]int{1,0,-1,2,3,-5,-2})
}