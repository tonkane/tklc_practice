package main

import "strings"

func wateringPlants(plants []int, capacity int) int {
	ans := 0
	n := len(plants)
	r := capacity
	for i := 0; i < n; i++ {
		// 移动到下一个位置
		ans += 1

		if r < plants[i] {
			ans += i * 2
			r = capacity
		}

		r -= plants[i]
	}
	return ans
}

type Trie struct {
	// val int
	words  [26]*Trie
	isWord bool
	word   string
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	cur := this
	for i := 0; i < n; i++ {
		idx := int(word[i] - 'a')
		if cur.words[idx] == nil {
			cur.words[idx] = &Trie{}
		}
		cur = cur.words[idx]
	}
	cur.isWord = true
	cur.word = word
}

func (this *Trie) Search(word string) bool {
	n := len(word)
	cur := this
	for i := 0; i < n; i++ {
		idx := int(word[i] - 'a')
		if cur.words[idx] == nil {
			return false
		}
		cur = cur.words[idx]
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	n := len(prefix)
	cur := this
	for i := 0; i < n; i++ {
		idx := int(prefix[i] - 'a')
		if cur.words[idx] == nil {
			return false
		}
		cur = cur.words[idx]
	}
	return true
}

// type WordDictionary struct {
// 	words [26]*WordDictionary
// 	isWord bool
// }

// func Constructor() WordDictionary {
// 	return WordDictionary{}
// }

// func (this *WordDictionary) AddWord(word string)  {
// 	n := len(word)
// 	cur := this
// 	for i := 0; i < n; i++ {
// 		idx := int(word[i]-'a')
// 		if cur.words[idx] == nil {
// 			cur.words[idx] = &WordDictionary{}
// 		}
// 		cur = cur.words[idx]
// 	}
// 	cur.isWord = true
// }

// func (this *WordDictionary) Search(word string) bool {
// 	n := len(word)

// 	var dfs func(int, *WordDictionary) bool
// 	dfs = func(idx int, node *WordDictionary) bool {
// 		if idx == n {
// 			return node.isWord
// 		}
// 		if word[idx] != '.' {
// 			word := node.words[word[idx]-'a']
// 			if word != nil && dfs(idx+1, word) {
// 				return true
// 			}
// 		} else {
// 			for i := range node.words {
// 				word := node.words[i]
// 				if word != nil && dfs(idx+1, word) {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	}

// 	return dfs(0, this)
// }

func findWords(board [][]byte, words []string) []string {
	n := len(board)
	m := len(board[0])

	v := make([][]bool, n)

	for i := range v {
		v[i] = make([]bool, m)
	}

	plus := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	tire := Constructor()

	for _, word := range words {
		tire.Insert(word)
	}

	find := map[string]bool{}

	var dfs func(i, j int, node *Trie)
	dfs = func(i, j int, node *Trie) {
		if i < 0 || j < 0 || i >= n || j >= m {
			return
		}

		bd := board[i][j]
		node = node.words[bd-'a']

		if node == nil {
			return
		}

		if v[i][j] {
			return
		}

		if node.word != "" {
			find[node.word] = true
		}

		v[i][j] = true

		for _, p := range plus {
			a := i + p[0]
			b := j + p[1]
			dfs(a, b, node)
		}
		v[i][j] = false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j, &tire)
		}
	}

	ans := []string{}

	for word := range find {
		ans = append(ans, word)
	}

	return ans
}

func combine(n int, k int) [][]int {
	ans := [][]int{}

	path := []int{}
	visited := make([]bool, n+1)
	var dfs func(step int, pre int)
	dfs = func(step int, pre int) {
		if step == k {
			ans = append(ans, append([]int{}, path...))
			return
		}

		if step > k {
			return
		}

		for i := range n {
			if !visited[i+1] && i >= pre {
				visited[i+1] = true
				path = append(path, i+1)
				dfs(step+1, i+1)
				path = path[:len(path)-1]
				visited[i+1] = false
			}
		}
	}

	dfs(0, 0)
	return ans
}

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	v := make([][]bool, n)

	for i := range v {
		v[i] = make([]bool, m)
	}

	puls := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	ans := false
	var dfs func(i, j int, cur string)
	dfs = func(i, j int, cur string) {
		if cur == "" {
			ans = true
			return
		}

		if ans {
			return
		}

		if 0 > i || 0 > j || i >= n || j >= m {
			return
		}

		if v[i][j] {
			return
		}

		v[i][j] = true
		for _, p := range puls {
			a := i + p[0]
			b := j + p[1]
			if board[i][j] == cur[0] {
				dfs(a, b, cur[1:])
			}
		}
		v[i][j] = false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ans {
				break
			}
			dfs(i, j, word)
		}
	}
	return ans
}


type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	head1 := head
	head2 := splitList(head)

	head1 = sortList(head1)
	head2 = sortList(head2)

	return mergeList(head1, head2)
}


func splitList(list *ListNode) *ListNode {
	slow := list
	fast := list.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	splitList2 := slow.Next
	slow.Next = nil
	return splitList2
}

func mergeList(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}
	return dummy.Next
}


func findPeakElement(nums []int) int {
	left := 0
	n := len(nums)
	right := len(nums)

	for left < right {
		mid := (left+right)/2

		if mid>0 && nums[mid]> nums[mid-1] && mid<n-1 && nums[mid] > nums[mid+1] {
			return mid
		} else {
			if mid<n-1 && nums[mid] < nums[mid+1] {
				left = mid +1
			} else {
				right = mid
			}
		}
	}

	if n >=2 {
		if nums[0]>nums[1] {
			return 0
		}
		if nums[len(nums)-1]>nums[len(nums)-2] {
			return len(nums)-1
		}
	}

	return 0
}


func main() {
	combine(4, 2)
}
