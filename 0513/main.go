package main

import (
	"archive/tar"
	"slices"
)

func orangesRotting(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// 记录
	r := make([][]int, n)
	for i := range r {
		r[i] = make([]int, m)
		for j := range r[i] {
			r[i][j] = 100
		}
	}

	p := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	var dfs func(i, j, minu int)
	dfs = func(i, j, minu int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if r[i][j] <= minu {
			return
		}
		if minu == 0 || grid[i][j] == 1 {
			r[i][j] = min(r[i][j], minu)
			for _, plus := range p {
				a := i + plus[0]
				b := j + plus[1]
				dfs(a, b, minu+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 2 {
				dfs(i, j, 0)
			}
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && r[i][j] == 100 {
				return -1
			}
			if grid[i][j] != 0 {
				ans = max(ans, r[i][j])
			}
		}
	}
	return ans
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, c0, c1 int) *Node {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] {
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &Node{
						true,
						false,
						dfs(rows[:rMid], c0, cMid),
						dfs(rows[:rMid], cMid, c1),
						dfs(rows[rMid:], c0, cMid),
						dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
	}

	return dfs(grid, 0, len(grid))
}

func moveZeroes(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	left := 0
	right := 1

	for right < n && left < n {
		if nums[left] != 0 {
			left++
			continue
		}
		if nums[right] == 0 || right <= left {
			right++
			continue
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// dummy := &ListNode{Next: head}
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return cur
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	curA := headA
	curB := headB

	countA := 0
	countB := 0

	for curA != nil || curB != nil {
		if curA != nil {
			curA = curA.Next
			countA++
		}

		if curB != nil {
			curB = curB.Next
			countB++
		}
	}

	curA = headA
	curB = headB

	// 让 A 长 一点
	if countA < countB {
		curA, curB = curB, curA
		countA, countB = countB, countA
	}

	d := countA - countB

	for curA != nil || curB != nil {
		if curA == curB {
			return curA
		}

		if curA != nil {
			d--
			curA = curA.Next
		}

		if curB != nil && d < 0 {
			curB = curB.Next
		}
	}

	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return ans
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		l := dfs(root.Left) + 1
		r := dfs(root.Right) + 1

		ans = max(ans, l+r)
		return max(l, r)
	}
	dfs(root)
	return ans
}

func findDuplicate(nums []int) int {
	n := len(nums)

	ans := 0
	for i := 0; i < n; {
		if nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				ans = nums[i]
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	return ans
}

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	n := len(nums)
	for i := 0; i < len(nums); {
		if nums[i] == 2 {
			for right >= 0 && nums[right] == 2 {
				right--
			}
			if i >= right {
				break
			}
			nums[i], nums[right] = nums[right], nums[i]
		} else if nums[i] == 0 {
			if i == left {
				i++
				left++
				continue
			}
			for left < n && nums[left] == 0 {
				left++
			}
			if left >= n {
				break
			}
			nums[i], nums[left] = nums[left], nums[i]
		} else {
			i++
		}
	}
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	// ans := 0
	len1 := len(text1)
	len2 := len(text2)
	mem := make([][]int, len1)

	for i := range mem {
		mem[i] = make([]int, len2)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= len1 || j >= len2 {
			return 0
		}

		// C := &mem[i][j]
		// if *C != -1 {
		// 	return *C
		// }

		if mem[i][j] != -1 {
			return mem[i][j]
		}

		var dmax int
		if text1[i] == text2[j] {
			dmax = dfs(i+1, j+1) + 1
		} else {
			// if i < len1 {
			dmax = max(dfs(i+1, j), dmax)
			// }
			// if j < len2 {
			dmax = max(dfs(i, j+1), dmax)
			// }
		}
		mem[i][j] = dmax
		return dmax
	}

	return dfs(0, 0)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[n][m]
}

func uniquePaths2(m int, n int) int {

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// dp[0][0] = 1

	for i := range n {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}

	return dp[m-1][n-1]
}

func uniquePaths(m int, n int) int {

	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// idx := (i+j) % n
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}

func canPartition2(nums []int) bool {
	sums := 0
	for _, v := range nums {
		sums += v
	}
	if sums % 2 == 1 {
		return false
	}

	sums /= 2
	n := len(nums)
	var dfs func(i int, target int) bool 
	dfs = func(i, target int) bool {
		if i >= n {
			return false
		}

		if target == 0 {
			return true
		}

		return dfs(i+1, target-nums[i]) || dfs(i+1, target)
	}
	return dfs(0, sums)
	// return false
}

func canPartition(nums []int) bool {
	sums := 0
	maxV := 0
	for _, v := range nums {
		sums += v
		maxV = max(maxV, v)
	}
	if sums % 2 == 1 {
		return false
	}
	sums /= 2
	if maxV > sums {
		return false
	}

	
	n := len(nums)

	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, sums+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		for j := 1; j <= sums; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][sums]
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}