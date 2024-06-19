package main

import "sort"

func maximumBeauty2(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	left := 0
	right := 0
	ans := 1
	for right < n && left < n {
		if nums[right]-k <= nums[left]+k {
			ans = max(right-left+1, ans)
			right++
		} else {
			left++
		}
	}

	return ans
}

func maximumBeauty(nums []int, k int) int {
    m := 0
    for _, x := range nums {
        m = max(m, x)
    }
    diff := make([]int, m + 2)
    for _, x := range nums {
        diff[max(x - k, 0)]++
        diff[min(x + k + 1, m + 1)]--
    }
    res, count := 0, 0
    for _, x := range diff {
        count += x
        res = max(res, count)
    }
    return res
}


func main() {
	maximumBeauty([]int{4,6,1,2}, 2)
}