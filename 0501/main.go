package main

import (
	"container/heap"
	"math"
	"sort"
)

type Heap [][]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] < h[j][1]
	}
	return h[i][0] < h[j][0]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	h := &Heap{}
	left, right := candidates-1, n-candidates

	if left+1 < right {
		for i := 0; i <= left; i++ {
			heap.Push(h, []int{costs[i], i})
		}
		for i := right; i < n; i++ {
			heap.Push(h, []int{costs[i], i})
		}
	} else {
		for i := 0; i < n; i++ {
			heap.Push(h, []int{costs[i], i})
		}
	}

	ans := int64(0)

	for i := 0; i < k; i++ {
		p := heap.Pop(h).([]int)
		cost, id := p[0], p[1]
		ans += int64(cost)
		if right - left > 1 {
			if id <= left {
				left++
				heap.Push(h, []int{costs[left], left})
			} else {
				right--
				heap.Push(h, []int{costs[right], right})
			}
		}
	}

	return ans
}


func containsNearbyDuplicate(nums []int, k int) bool {
	ni := make(map[int]int)
	count := len(nums)
	for i := 0; i < count; i++ {
		idx, ok := ni[nums[i]]
		if !ok {
			ni[nums[i]] = i
		} else {
			if i-idx <= k {
				return true
			} else {
				ni[nums[i]] = i
			}
		}
	}

	return false
}


func isAnagram(s string, t string) bool {
	r := [26]int{}
	sc := len(s)
	tc := len(t)
	if sc != tc {
		return false
	}
 	for i := 0; i < sc; i++ {
		r[s[i]-'a']++
	}
	for i := 0; i < tc; i++ {
		r[t[i]-'a']--
		if r[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}


func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	count := len(nums)
	for i := 0; i < count; i++ {
		m[nums[i]] = true
	}
	ans := 0
	for key, value := range m {
		if value {
			m[key] = false
			tmp := 1
			left := key-1
			right := key+1
			for m[left] {
				tmp++
				m[left] = false
				left--
				
			}
			for m[right] {
				tmp++
				m[right] = false
				right++
			}
			ans = max(ans, tmp)
		}
	}
	return ans
}

type ListNode struct {
    Val int
    Next *ListNode
}


func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}
	low := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		if low == fast {
			return true
		}
		low = low.Next
		fast = fast.Next.Next
	}
	return false
}

func plusOne(digits []int) []int {
	count := len(digits)
	plus := 0
	digits[len(digits)-1]++
	if digits[len(digits)-1] < 10 {
		return digits
	}
	for i := count-1; i >=0; i-- {
		digits[i] += plus
		if digits[i] == 10 {
			digits[i] = 0
			plus = 1
		} else {
			plus = 0
		}
	}
	if plus == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}


func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	down := 1
	count := 0
	xc := x
	for xc > 1 {
		xc /= 2
		count++
	}
	count /= 2
	for count > 0{
		down*=2 
		count--
	}
	// 上届
	up := down*2
	var mid int
	for down < up {
		mid = (down+up)/2

		pow := mid*mid
		pow2 := (mid+1)*(mid+1)

		if pow <= x && pow2 > x {
			return mid
		}

		if pow < x {
			down = mid+1
		} else {
			up = mid
		}
	}

	return down
}



func searchInsert(nums []int, target int) int {
	count := len(nums)
	left, right := 0, count
	var mid int
	for left < right {
		mid = (left+right)/2

		if nums[mid] < target {
			left = mid+1
		} else {
			right = mid
		}
	}
	return left
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		lmax := dfs(node.Left)+1
		rmax := dfs(node.Right)+1

		return max(lmax, rmax)
	}

	return dfs(root)+1
}


func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}


func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := invertTree(root.Left)
	r := invertTree(root.Right)
	
	root.Left = r
	root.Right = l

	return root
}


func isSymmetric(root *TreeNode) bool {
	var dfs func(l, r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
	}
	return dfs(root.Left, root.Right)
}


func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var dfs func(root *TreeNode, targetSum int, isLeaf bool) bool
	dfs = func(root *TreeNode, targetSum int, isLeaf bool) bool {
		if root == nil {
			if targetSum == 0 && isLeaf {
				return true
			} else {
				return false
			}
		}
		isLeaf = root.Left == nil && root.Right == nil
		return dfs(root.Left, targetSum-root.Val, isLeaf) || dfs(root.Right, targetSum-root.Val, isLeaf)
	}
	return dfs(root, targetSum, root.Left == nil && root.Right == nil)
}

func main() {
	mySqrt(100)
	println(math.Sqrt2)
}
