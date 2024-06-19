package main

import (
	"math"
	"strings"
)

func garbageCollection(garbage []string, travel []int) int {
	ans := 0

	// 收垃圾需要的时间
	lastM := -1
	lastG := -1
	lastP := -1
	for i, g := range garbage {
		ans += len(g)
		if strings.Contains(g, "M") {
			lastM = i
		}
		if strings.Contains(g, "G") {
			lastG = i
		}
		if strings.Contains(g, "P") {
			lastP = i
		}
	}

	lastM--
	lastG--
	lastP--

	sums := 0 
	for i, t := range travel {
		sums += t
		if lastM == i {
			ans += sums
		}
		if lastG == i {
			ans += sums
		}
		if lastP == i {
			ans += sums
		}
	}

	return ans
}


type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}


func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	
	count := len(q)
	for count > 0 {
		cur := q[0]
		q = q[1:]

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}

		count--

		if count == 0 {
			count = len(q)
			for j := 0; j < count-1; j++ {
				q[j].Next = q[j+1]
			}
		}
	}

	return root
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


type BSTIterator struct {
	nodeVal []int
	next int
	size int
}


func Constructor(root *TreeNode) BSTIterator {
	nodeVal := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		nodeVal = append(nodeVal, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return BSTIterator{
		nodeVal: nodeVal,
		next : -1,
		size : len(nodeVal),
	}
}


func (this *BSTIterator) Next() int {
	this.next++
	return this.nodeVal[this.next]
}


func (this *BSTIterator) HasNext() bool {
	if this.next < this.size-1 {
		return true
	}
	return false
}



func findOrder(numCourses int, prerequisites [][]int) []int {
	v := make([]int, numCourses)


	m := make([][]int, numCourses)

	// for i := range m {
	// 	m[i] = make([]int, 0)
	// }

	for _, p := range prerequisites {
		m[p[0]] = append(m[p[0]], p[1])
	}
	ans := []int{}
	sc := true
	var dfs func(i int)
	dfs = func (i int)  {
		v[i] = 1
		for _, mv := range m[i] {
			if v[mv] == 0 {
				dfs(mv)
				ans = append(ans, mv)
				// v[mv] = 1
			} else if v[mv] == 1 {
				// 成环
				sc = false
			}
		}
		v[i] = 2
	}

	for i, ok := range v {
		// 未访问到
		if ok == 0 {
			// v[i] = 1		
			dfs(i)
			ans = append(ans, i)
		}
		if !sc {
			return []int{}
		}
	}
	return ans
}



func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	dp[0] = []int{triangle[0][0]}

	if n == 1 {
		return triangle[0][0]
	}

	ans := math.MaxInt

	for i := 1; i < n; i++ {
		for j:=0; j<i+1; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
			if i == n-1 {
				ans = min(ans, dp[i][j])
			}
		}
		
	}
	return ans
}





func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]

	for j:=1; j<m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 0; i < n-1; i++ {
		dp[i+1][0] = dp[i][0] + grid[i+1][0]
		for j := 1; j < m; j++ {
			dp[i+1][j] = min(dp[i][j], dp[i+1][j-1]) + grid[i+1][j]
		}
	}

	return dp[n-1][m-1]
}



func main() {
	// minPathSum([][]int{{1,3,1}, {1,5,1}, {4,2,1}})
	minPathSum([][]int{{1,2,3}, {4,5,6}})
}