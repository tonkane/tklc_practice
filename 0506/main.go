package main

import (
	"math"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][][]int, n*2-1)

	// f[t][j][k]  走t步，走到 (t-j, j) 和 (t-k, k)
	// 由于只能向下或者向右 i+1 或 j+1 所以 i+j=t
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, n+1)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt
			}
		}
	}

	// 边界条件
	f[0][1][1] = grid[0][0]
	for t := 1; t<n*2-1; t++ {
		// i+j=t && 0<=i<=n-1 && 0<=j<=n-1
		// 可走的边界要注意 t 的关系
		for j := max(t-n+1, 0); j<=min(t, n-1); j++ {
			// 不可达
			if grid[t-j][j] < 0 {
				continue
			}

			for k :=j; k<=min(t, n-1); k++ {
				if grid[t-k][k] < 0 {
					continue
				}
				// 状态转移
				f[t][j+1][k+1] = max(f[t-1][j+1][k+1], f[t-1][j+1][k], f[t-1][j][k+1], f[t-1][j][k]) + grid[t-j][j]
				// k == j 时意味着到同一个点
				// 不需要要重复加樱桃数量
				if k != j {
					f[t][j+1][k+1] += grid[t-k][k] 
				}
			}

		}
	}

	return max(f[n*2-2][n][n], 0)
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree2(preorder []int, inorder []int) *TreeNode {

	var dfs func(p []int, i[]int) *TreeNode
	dfs = func(p, i []int) *TreeNode {
		if len(p) == 0 {
			return nil
		}
		node := &TreeNode{Val: p[0]}
		idx := 0
		for idx<len(i) {
			if i[idx] == p[0] {
				break
			}
			idx++
		}
		node.Left = dfs(p[1:idx+1], i[0:idx])
		node.Right = dfs(p[idx+1:], i[idx+1:])

		return node
	}

	return dfs(preorder, inorder)
}



func buildTree(inorder []int, postorder []int) *TreeNode {

	// 反转
	for i, j := 0, len(postorder)-1; i < j; i, j = i+1, j-1 {
		postorder[i], postorder[j] = postorder[j], postorder[i]
	}

	var dfs func(p []int, i[]int) *TreeNode
	dfs = func(p, i []int) *TreeNode {
		if len(p) == 0 {
			return nil
		}
		node := &TreeNode{Val: p[0]}
		idx := 0
		n := len(i)
		for idx<n {
			if i[idx] == p[0] {
				break
			}
			idx++
		}

		node.Left = dfs(p[n-idx:], i[:idx])
		node.Right = dfs(p[1:n-idx], i[idx+1:])

		return node
	}

	return dfs(postorder, inorder)
}



func maxPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt
	}

	l := maxPathSum(root.Left)
	r := maxPathSum(root.Right)

	imax := max(root.Val+l+r, root.Val+l)
	imax = max(root.Val+r, imax)

    return imax
}


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}

	q := []*TreeNode{root}

	count := len(q)
	list := []int{}
	for count > 0 {
		cur := q[0]
		q = q[1:]
		list = append(list, cur.Val)

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}

		count--

		if count == 0 {
			ans = append(ans, list)
			count = len(q)
			list = []int{}
		}
	}

	return ans
}




func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := [][]int{}

	q := []*TreeNode{root}

	count := len(q)
	list := []int{}
	isLeft := true
	for count > 0 {
		cur := q[0]
		q = q[1:]

		if isLeft {
			list = append(list, cur.Val)
		} else {
			list = append([]int{cur.Val}, list...)
		}

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}

		count--

		if count == 0 {
			ans = append(ans, list)
			count = len(q)
			list = []int{}
			isLeft = !isLeft
		}
	}

	return ans
}


func kthSmallest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		k--
		if k == 0 {
			ans =  node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}


func singleNumber(nums []int) int {
	r := [64]int{}

	for _, num := range nums {
		for i:=0; i<64; i++ {
			if num & 1 == 1 {
				r[i]++
			}
			num = num>>1
		}
	}

	ans := 0
	for i:=0; i<64; i++ {
		if r[i] %3 == 1 {
			ans |= (1<<i)
		}
	}
	return ans
}



func rangeBitwiseAnd(left int, right int) int {
	// 把最后一个1变成0 看范围是否大于left
	for right > left {
		right = right & (right-1)
	}
	return right
}


func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}


func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y*y
	}
	return y*y*x
}


func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}

	return dp[n]
}


func wordBreak(s string, wordDict []string) bool {
	n := len(wordDict)
	m := make(map[string]bool, n)
	for _, word := range wordDict {
		m[word] = true
	}

	ls := len(s)
	dp := make([]int, ls+1)
	dp[0] = 1
	for i := 0; i < ls; i++ {
		for j:=i; j>=max(i-20, 0); j-- {
			_, ok := m[s[j:i+1]]
			println(s[j:i+1])
			if ok && dp[j] == 1 {
				dp[i+1] = 1
				break
			}
		}
	}

	return dp[ls] == 1
}


func coinChange(coins []int, amount int) int {
	// 每个数据记录几步能到达
	dp := make([]int, amount+1)
	c := len(coins)

	for i := 0; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < c; j++ {
			if i>=coins[j] {
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}


func main() {
	coinChange([]int{1,2,5}, 11)
}