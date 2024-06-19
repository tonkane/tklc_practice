package main


func minimumRounds(tasks []int) int {
	m := make(map[int]int)
	for _,v := range tasks {
		m[v]++
	}
	ans := 0
	for _, value := range m {
		if value == 1 {
			return -1
		}
		ans += value/3
		if value % 3 != 0 {
			ans++
		}
	}
	return ans
}



func findAnagrams(s string, p string) []int {
	a := len(s)
	b := len(p)
	// sums := 0

	if a < b {
		return nil
	}

	m := make(map[byte]int)

	for i := 0; i < b; i++ {
		m[p[i]]++
	}

	for j := 0; j < b; j++ {
		m[s[j]]--
		if m[s[j]] == 0 {
			delete(m, s[j])
		}
	}
	ans := []int{}
	left := 0
	right := b-1

	for {
		if len(m) == 0 {
			ans = append(ans, left)
		}
		m[s[left]]++
		if m[s[left]] == 0 {
			delete(m, s[left])
		}
		left++
		right++

		if right >= a {
			break
		}

		m[s[right]]--
		if m[s[right]] == 0 {
			delete(m, s[right])
		}
	}

	return ans
}


func subarraySum2(nums []int, k int) int {
	n := len(nums)
	s := make([]int, n)

	s[0] = nums[0]
	ans := 0
	if s[0] == k {
		ans++
	}
	for i := 1; i < n; i++ {
		s[i] += s[i-1]+nums[i]
		if s[i] == k {
			ans++
		}
	}


	for i := 0; i < n; i++ {
		for j := i-1; j >=0; j-- {
			if s[i] - s[j] == k {
				ans++
			}
		}
	}

	return ans
}


func subarraySum(nums []int, k int) int {
	n := len(nums)
	m := make(map[int]int)

	m[0] = 1

	sums := 0
	ans := 0
	for i := 0; i < n; i++ {
		sums += nums[i]
		if value, ok := m[sums-k]; ok {
			ans += value
		}
		m[sums]++
	}

	return ans
}


type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}


	dummy := &ListNode{Next: head}
	cur := dummy.Next
	pre := dummy
	for cur != nil && cur.Next != nil {
		cnn := cur.Next.Next
		cur.Next.Next = cur
		pre.Next = cur.Next
		cur.Next = cnn
		pre = cur
		cur = cnn
	}

	return dummy.Next
}


func isPalindrome(head *ListNode) bool {

	tmp := []int{}

	cur := head
	for cur != nil {
		tmp = append(tmp, cur.Val)
		cur = cur.Next
	}

	left := 0
	right := len(tmp)-1

	for left < right {
		if tmp[left] != tmp[right] {
			return false
		}
		left++
		right--
	}

	return true
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func pathSum(root *TreeNode, targetSum int) int {
	// 1. dfs 对每个节点暴力一遍
	// 2. 前缀和
	// curSums - 前缀 = target
	// curSums - target = 前缀
	ans := 0
	pre := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, cur int) {
		if tn == nil {
			return
		}
		cur += tn.Val
		ans += pre[cur-targetSum]
		pre[cur]++
		dfs(tn.Left, cur)
		dfs(tn.Right, cur)
		pre[cur]--
		return
	}
	dfs(root, 0)
	return ans
}


func searchMatrix(matrix [][]int, target int) bool {
	// 右上角
	// 二叉搜索树
	n := len(matrix)
	m := len(matrix[0])
	i, j := 0, m-1
	for i<n && j>= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}


func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for {
		slow = slow.Next
		// 无环
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 重新出发一个指针
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}



func main() {
	var node2 *ListNode
	// node4 := &ListNode{Val: -4, Next:  node2}
	// node3 := &ListNode{Val: 0, Next:  node4}
	node2 = &ListNode{Val: 2, Next:  nil}
	node1 := &ListNode{Val: 3, Next:  node2}

	detectCycle(node1)
}