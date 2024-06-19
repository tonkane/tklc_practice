package main

// import "sort"

func nextPermutation(nums []int)  {
	length := len(nums)
	// 从后往前查找第一个较小数 nums[i] < nums[i+1]
	idx := -1
	for i := length-1; i >=1; i-- {
		if nums[i] > nums[i-1] {
			idx = i-1
			break
		}
	}

	idx2 := -1
	if idx >= 0 {
		// 从后往前找大于较小数的数
		for i := length-1; i >=0; i-- {
			if nums[i] > nums[idx] {
				idx2 = i
				break
			}
		}
		nums[idx], nums[idx2] = nums[idx2], nums[idx]
	}

	r(nums[idx+1:])
}

// 反转
func r(nums []int) {
	for i, n :=0, len(nums); i<n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func isValidSudoku(board [][]byte) bool {
	var rows [10][10]int
	var cols [10][10]int
	var matrix [10][10]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			value := board[i][j] - '0'
			if value>=1 && value<=9 {
				rows[i][value]++
				if rows[i][value] > 1 {
					return false
				}
				cols[j][value]++
				if cols[j][value] > 1 {
					return false
				}
				mIdx := i/3 + j/3*3
				matrix[mIdx][value]++
				if matrix[mIdx][value] > 1 {
					return false
				}
			}
		}
	}

	return true
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	var dfs func(idx int, visit [21]int)
	path := []int{}
	length := len(nums)
	dfs = func(idx int, visit [21]int) {
		if idx >= length {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < length; i++ {
			if visit[nums[i]+10] == 0 {
				path = append(path, nums[i])
				visit[nums[i]+10]++
				dfs(idx+1, visit)
				visit[nums[i]+10]--
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, [21]int{})
	return ans
}


func main() {
	nextPermutation([]int{2,3,1})
}