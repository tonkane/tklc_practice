package main

func longestEqualSubarray(nums []int, k int) int {
    n := len(nums)
    ans := 0
    cnt := make(map[int]int)
    for i, j := 0, 0; j < n; j++ {
        cnt[nums[j]]++
        /*当前区间中，无法以 nums[i] 为等值元素构成合法等值数组*/
        for j - i + 1 - cnt[nums[i]] > k {
            cnt[nums[i]]--
            i++
        }
        if cnt[nums[j]] > ans {
            ans = cnt[nums[j]]
        }
    }
    return ans
}

func main() {
	longestEqualSubarray([]int{1,3,2,3,1,3}, 3)
}
