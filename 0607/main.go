package main

func maxOperations(nums []int) int {
	n := len(nums)
	ans := 1
	pre := nums[0] + nums[1]
	for i := 3; i < n; i+=2 {
		if nums[i] + nums[i-1] == pre {
			ans++
		} else {
			break
		}
	}

	return ans
}

func main() {

}
