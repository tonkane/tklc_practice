package main

func maxOperations(nums []int) int {
	n := len(nums)
	var dfs func(i int, j int, pre int) int
	// 记忆化存储
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
	}

	dfs = func(i int, j int, pre int) (count int) {
		if j-i < 2 {
			return
		}

		if mem[i][j] != 0 {
			return mem[i][j]
		}
		
		if pre == nums[i]+nums[i+1] {
			count = max(count, dfs(i+2, j, pre)+1)
		}
		if pre == nums[j-1]+nums[j-2] {
			count = max(count, dfs(i, j-2, pre)+1)
		}
		if pre == nums[i]+nums[j-1] {
			count = max(count, dfs(i+1, j-1, pre)+1)
		}

		mem[i][j] = count

		return
	}


	l := dfs(2, n, nums[0]+nums[1])
	j := dfs(0, n-2, nums[len(nums)-1]+nums[len(nums)-2])
	k := dfs(1, n-1, nums[0]+nums[len(nums)-1])

	return 1+ max(l,j,k)
}


func main() {
	maxOperations([]int{1,9,7,3,2,7,4,12,2,6})
}