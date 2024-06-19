package main

import (
	"strconv"
	"strings"
)

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)

	sums := 0
	if k > 0 {
		for i := 0; i < k; i++ {
			sums += code[i+1]
		}
		ans[0] = sums
	}
	if k < 0 {
		for i := n - 1; i >= n+k; i-- {
			sums += code[i]
		}
		ans[0] = sums
	}

	for i := 1; i < n; i++ {
		if k > 0 {
			sums -= code[i]
			sums += code[(i+k)%n]

		}
		if k < 0 {
			sums += code[i-1]
			sums -= code[(n+i+k-1)%n]
		}
		ans[i] = sums
	}

	return ans
}

type MinStack struct {
	s  []int
	ms []int
}

func Constructor() MinStack {
	r := MinStack{
		[]int{},
		[]int{},
	}
	return r
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if len(this.ms) == 0 {
		this.ms = append(this.ms, val)
	} else {
		if val <= this.ms[len(this.ms)-1] {
			this.ms = append(this.ms, val)
		}
	}
}

func (this *MinStack) Pop() {
	last := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	if last == this.ms[len(this.ms)-1] {
		this.ms = this.ms[:len(this.ms)-1]
	}
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.ms[len(this.ms)-1]
}

func evalRPN(tokens []string) int {
	// n := len(tokens)
	q := []int{}

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			l1 := q[len(q)-1]
			l2 := q[len(q)-2]
			var r int
			if token == "+" {
				r = l1 + l2
			}
			if token == "-" {
				r = l2 - l1
			}
			if token == "*" {
				r = l1 * l2
			}
			if token == "/" {
				r = l2 / l1
			}
			q[len(q)-2] = r
			q = q[:len(q)-1]
		} else {
			intVal, _ := strconv.ParseInt(token, 0, 32)
			q = append(q, int(intVal))
		}
	}

	if len(q) > 0 {
		return q[0]
	} else {
		return 0
	}
}

func simplifyPath(path string) string {
	l := strings.Split(path, "/")
	n := len(l)
	r := []string{}
	for i := 0; i < n; i++ {
		if l[i] == "" || l[i] == "." {
			continue
		} else if l[i] == ".." {
			if len(r) > 0 {
				r = r[:len(r)-1]
			}
		} else {
			r = append(r, l[i])
		}
	}

	// ans := strings.Join(r, "/")

	return "/" + strings.Join(r, "/")
}


type Node struct {
    Val int
    Next *Node
    Random *Node
}


func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	// cur.Next = head.Next

	for cur != nil {
		node := &Node{Val: cur.Val, Next: cur.Next, Random: cur.Random}
		cur.Next = node
		cur = cur.Next.Next
	}

	
	re := head.Next

	cur = head.Next
	for cur != nil {
		if cur.Random != nil {
			cur.Random = cur.Random.Next
		}
		if cur.Next == nil {
			break
		}
		cur = cur.Next.Next
	}

	cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = cur.Next.Next
		
		if next.Next != nil {
			next.Next = next.Next.Next
		}	

		cur = cur.Next
	}

	return re
}

type ListNode struct {
    Val int
    Next *ListNode
}


func reverseBetween(head *ListNode, left int, right int) *ListNode {
	h := &ListNode{Val: 0, Next: head}

	count := 0
	var f *ListNode
	var s *ListNode
	var p *ListNode
	// var next *ListNode
	cur := h
	for cur != nil {
		next := cur.Next
		// 记录反转前的节点
		if count == left-1 {
			f = cur
		}
		if count == left {
			s = cur
			s.Next = nil
		}
		// 需要反转的节点
		if count > left && count<=right {
			cur.Next = p
		}
		if count == right {
			f.Next = cur
		}
		if count == right+1 {
			s.Next = cur
			break
		}
		p = cur
		cur = next
		count++
	}

	return h.Next
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{Val: 0, Next: head}

	fast := h
	slow := h
	// 快节点先出发 n 步
	// 慢节点再出发，快节点到底时，慢节点为需要删除的节点

	count := 0
	for {
		for count <= n {
			fast = fast.Next
			count++
		}
		if fast == nil {
			// 删除节点
			if slow.Next != nil {
				slow.Next = slow.Next.Next
			}
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	return h.Next
}


func deleteDuplicates(head *ListNode) *ListNode {
	h := &ListNode{Val: -101, Next: head}

	cur := h
	var p *ListNode
	for cur !=nil && cur.Next != nil {
		if cur.Next.Val == cur.Val {
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			p.Next = cur.Next
			// p = cur.Next
		} else {
			p = cur
		}
		cur = cur.Next
	}
	return h.Next
}


func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	count := 0
	cur := head
	var p *ListNode
	for cur != nil {
		count++
		if cur.Next == nil {
			cur.Next = head
			break
		}
		cur = cur.Next
	}

	nk := k % count

	cur = head
	for count>nk {
		p = cur
		cur = cur.Next
		nk++
	}
	p.Next = nil
	return cur
}


func main() {
	var node1 *Node
	var node2 *Node
	var node3 *Node
	node3 = &Node{Val: 3, Random: nil, Next: nil}
	node2 = &Node{Val: 2, Random: node1, Next: node3}
	node1 = &Node{Val: 1, Random: nil, Next: node2}

	copyRandomList(node1)
}
