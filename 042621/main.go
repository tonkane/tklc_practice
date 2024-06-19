package main

import "sort"

func letterCombinations(digits string) []string {
	nums := [][]string{{""}, {""}, {"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"},
		{"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}
	var dfs func(idx int)
	leng := len(digits) - 1
	ans := []string{}
	var path string
	dfs = func(idx int) {
		if idx > leng {
			return
		}
		for _, value := range nums[int(digits[idx]-'0')] {
			path += value
			dfs(idx + 1)
			if idx == leng {
				ans = append(ans, path)
				// return
			}
			path = path[:len(path)-1]
		}

	}
	dfs(0)
	return ans
}

func generateParenthesis(n int) []string {
	l := "("
	r := ")"
	ans := []string{}

	total := n * 2
	// lnums := 0

	path := ""
	var dfs func(idx int, lnums int)
	dfs = func(idx int, lnums int) {
		if idx == total {
			ans = append(ans, path)
			return
		}
		// 左括号个数小于 N 可添加左括号
		if lnums < n {
			path += l
			dfs(idx+1, lnums+1)
			path = path[:len(path)-1]
		}
		// 右括号
		if idx-lnums < lnums {
			path += r
			dfs(idx+1, lnums)
			path = path[:len(path)-1]
		}

	}
	dfs(0, 0)
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	// 排序
	sort.Ints(candidates)
	var dfs func(idx int, target int, pre int)
	path := []int{}
	mIdx := len(candidates) - 1
	dfs = func(idx, target, pre int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if target < 0 || idx > mIdx {
			return
		}
		// 选
		if pre != candidates[idx] {
			path = append(path, candidates[idx])
			dfs(idx+1, target-candidates[idx], 0)
			path = path[:len(path)-1]
		}
		// 不选
		dfs(idx+1, target, candidates[idx])
	}
	dfs(0, target, 0)
	return ans
}

func main() {
	// combinationSum2([]int{10,1,2,7,6,1,5}, 8)
	println(2/3)
}