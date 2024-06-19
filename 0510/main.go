package main

import "slices"

func countTestedDevices(batteryPercentages []int) int {
	ans := 0
	for _, b := range batteryPercentages {
		if b > ans {
			ans++
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	res := [][]*TreeNode{}
	path := []*TreeNode{}
	var dfs func(node *TreeNode, find *TreeNode)
	dfs = func(node *TreeNode, find *TreeNode) {
		if node == nil {
			return
		}

		path = append(path, node)

		if node.Val == find.Val {
			res = append(res, append([]*TreeNode{}, path...))
			return
		}

		dfs(node.Left, find)
		dfs(node.Right, find)
		path = path[:len(path)-1]
	}

	dfs(root, p)
	path = []*TreeNode{}
	dfs(root, q)

	n := min(len(res[0]), len(res[1]))
	i := 0
	for ; i < n; i++ {
		if res[0][i].Val == res[1][i].Val {
			continue
		} else {
			break
		}
	}
	i--
	println(res[0][i].Val)
	return res[0][i]
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	} else {
		return left
	}
}

func maxSubarraySumCircular(nums []int) int {
	ansMin, ansMax := nums[0], nums[0]
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	s := make([]int, n)
	s[0] = nums[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + nums[i]
	}
	// sums := s[len(s)-1]
	// 最大的连续区间
	cMin := nums[0]
	for j := 1; j < n; j++ {
		ansMax = max(ansMax, s[j])
		ansMax = max(ansMax, s[j]-cMin)
		if cMin > s[j] {
			cMin = s[j]
		}
	}

	// 最小的连续区间
	cMax := nums[0]
	for j := 1; j < n-1; j++ {
		ansMin = min(ansMin, s[j])
		ansMin = min(ansMin, s[j]-cMax)
		if cMax < s[j] {
			cMax = s[j]
		}
	}
	return max(ansMax, s[len(s)-1]-ansMin)
}

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	ans := 0
	var dfs func(idx int, pre int, count int)
	dfs = func(idx int, pre int, count int) {
		if idx >= n {
			ans = max(ans, count)
			return
		}

		if nums[idx] > pre {
			dfs(idx+1, nums[idx], count+1)
		}
		dfs(idx+1, pre, count)

	}
	dfs(0, -10001, 0)
	return ans
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
	}
	return slices.Max(dp)
}



func canFinish(numCourses int, prerequisites [][]int) bool {
	n := len(prerequisites)
	v := make([]bool, numCourses)

	mm := make([]map[int]bool, numCourses)

	for i := range mm {
		mm[i] = make(map[int]bool, 0)
	}

	for i := 0; i < n; i++ {
		mm[prerequisites[i][0]][prerequisites[i][1]] = true
	}


	var dfs func(idx int, start int) bool
	dfs = func(idx int, start int) bool {
		if v[idx] && idx == start {
			return false
		}

		res := true
		if !v[idx] {
			v[idx] = true
			for key := range mm[idx] {
				res = res && dfs(key, start)
			}
		}

		return res
	}

	for i := 0; i < numCourses; i++ {
		if len(mm[i]) > 0 {
			ok := dfs(i, i)
			if !ok {
				return false
			}
			v = make([]bool, numCourses)
		}
	}

	return true
}



func main() {
	canFinish(4, [][]int{{2,0}, {1,0}, {3,1}, {3,2}, {1,3}})
	// lengthOfLIS([]int{4, 10, 4, 3, 8, 9})
	// maxSubarraySumCircular([]int{-2})
	// node6 := &TreeNode{Val: 6, }
	// node7 := &TreeNode{Val: 7, }
	// node4 := &TreeNode{Val: 4, }
	// node0 := &TreeNode{Val: 0, }
	// node8 := &TreeNode{Val: 8, }
	// node2 := &TreeNode{Val: 2, Left: node7, Right: node4}
	// node5 := &TreeNode{Val: 5, Left: node6, Right: node2}
	// node1 := &TreeNode{Val: 1, Left: node0, Right: node8}
	// node3 := &TreeNode{Val: 3, Left: node5, Right: node1}
	// lowestCommonAncestor(node3, node5, node1)
}