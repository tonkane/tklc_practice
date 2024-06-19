package main

import (
	"container/heap"
	"math"
	"slices"
	"strings"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i] > h[j] {
		return true
	} else {
		return false
	}
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Peek() interface{} {
	if h.Len() == 0 {
		return -1
	}
	return h[0]
}

var _ heap.Interface = &Heap{}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	h := &Heap{}
	n := len(quality)
	// 每个的工资比
	type pair struct{ q, w int }
	pairs := make([]pair, n)
	for i, q := range quality {
		pairs[i] = pair{q, wage[i]}
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		return a.w*b.q - b.w*a.q
	})
	// ans := math.MaxFloat64

	sumsQ := 0

	for _, p := range pairs[:k] {
		heap.Push(h, p.q)
		sumsQ += p.q
	}

	// 可能的最小值
	ans := float64(sumsQ*pairs[k-1].w) / float64(pairs[k-1].q)

	// 后面的工人薪资比更大
	for _, p := range pairs[k:] {
		if p.q < h.Peek().(int) {
			sumsQ -= h.Peek().(int) - p.q
			heap.Pop(h)
			heap.Push(h, p.q)
			// heap.Fix(h, 0)
			ans = min(ans, float64(sumsQ*p.w)/float64(p.q))
		}
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{}
	q = append(q, root)
	ans := []float64{}
	count := len(q)
	sums := 0
	nums := 0
	for count > 0 {
		cur := q[0]
		q = q[1:]
		sums += cur.Val
		count--
		nums++

		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}

		if count == 0 {
			count = len(q)
			ans = append(ans, float64(sums)/float64(nums))
			sums = 0
			nums = 0
		}
	}

	return ans
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	ans := 0
	for {
		cur := q[0]
		q = q[1:]
		ans++

		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}

		if len(q) == 0 {
			break
		}
	}
	return ans
}

func getMinimumDifference2(root *TreeNode) int {
	if root == nil {
		return math.MaxInt
	}

	l := math.MaxInt
	r := math.MaxInt
	if root.Left != nil {
		l = min(abs(root.Val-root.Left.Val), getMinimumDifference2(root.Left))
	}
	if root.Right != nil {
		r = min(abs(root.Val-root.Right.Val), getMinimumDifference2(root.Right))
	}
	return min(l, r)
}

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt
	var dfs func(node *TreeNode)
	pre := math.MaxInt
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 中序遍历
		dfs(node.Left)
		ans = min(ans, abs(node.Val-pre))
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return a * -1
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	mid := length/2
	if length == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(append([]int{}, nums[0:mid]...))
	root.Right = sortedArrayToBST(append([]int{}, nums[mid+1:]...))

	return root
}


func addBinary(a string, b string) string {
	ans := ""
	al := len(a)-1
	bl := len(b)-1

	// 处理成同一个长度好处理
	if al>bl {
		b = strings.Repeat("0", al-bl) + b
	}
	if bl>al {
		a = strings.Repeat("0", bl-al) + a
	}

	l := max(al, bl)

	plus := false 
	for i := l; i >=0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if plus {
				ans = "1" + ans
			} else {
				ans = "0" + ans
			}
			plus = true
		} else if a[i] == '0' && b[i] == '0' {
			if plus {
				ans = "1" + ans
			} else {
				ans = "0" + ans
			}
			plus = false
		} else {
			if plus {
				ans = "0" + ans
			} else {
				ans = "1" + ans
				plus = false
			}
		}
	}

	if plus {
		ans = "1" + ans
	}

	return ans
}


func hammingWeight(n int) int {
	ans := 0
	for n != 0 {
		n = n & (n-1)
		ans++
	}
	return ans
}



func singleNumber(nums []int) int {
	ans := 0

	for _,v := range nums {
		ans ^= v
	}
	
	return ans
}


func reverseBits(num uint32) uint32 {
	count := 32
	ans := uint32(0)
	for count>0 {
		ans |= num & 1 << (count-1)
		num >>= 1
		count--
	}
	return ans
}


func main() {
	nums := []int{1, 2}
	a := nums[0:1]
	b := nums[2:]

	println(a, b)
}
