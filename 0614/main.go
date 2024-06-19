package main

func maxScore(nums []int, x int) int64 {
	n := len(nums)

	// 当是偶数时最大和值
	// 当是奇数时最大和值
	o := int64(nums[0])
	j := int64(nums[0])-int64(x)
	if nums[0] % 2 != 0 {
		o, j = j, o 
	}
	
	ans := max(o, j)
	for i := 1; i < n; i++ {
		if nums[i] % 2 == 0 { // 偶数
			o = max(o+int64(nums[i]), j+int64(nums[i])-int64(x))
			ans = max(ans, o)
		} else {
			j = max(j+int64(nums[i]), o+int64(nums[i])-int64(x))
			ans = max(ans, j)
		}
	}

	return ans
}

func main() {

}