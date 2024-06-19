package main

// import "google.golang.org/appengine/capability"

// import "fmt"

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ans := 0
	left := 0
	right := len(plants) - 1

	rA := capacityA
	rB := capacityB
	for left < right {
		if rA < plants[left] {
			rA = capacityA
			ans++
		}
		if rB < plants[right] {
			rB = capacityB
			ans++
		}

		rA -= plants[left]
		rB -= plants[right]

		left++
		right--
	}

	if left == right {
		if rA < plants[left] && rB < plants[right] {
			ans++
		}
	}

	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val:-101, Next: head}

	pre := dummy
	cur := dummy.Next
	find := false
	var bigHead *ListNode
	var p *ListNode
	for cur != nil {
		if !find && cur.Val >= x {
			bigHead = cur
			find = true
		}

		if find && cur.Val >= x {
			p = cur
		}

		if find && cur.Val < x {
			p.Next = cur.Next
			cur.Next = nil
			pre.Next = cur
			pre = pre.Next
			cur = p
		}

		if !find {
			pre = cur
		}

		cur = cur.Next
	}

	pre.Next = bigHead

	return dummy.Next
}

// func partition(head *ListNode, x int) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	dummy := &ListNode{Val: -101, Next: head}

// 	pre := dummy
// 	cur := dummy
// 	find := false
// 	var bigHead *ListNode
// 	for cur != nil && cur.Next != nil {
// 		if !find && cur.Val >= x {
// 			bigHead = cur
// 			find = true
// 		}

// 		if find && cur.Next.Val < x {
// 			pre.Next = cur.Next
// 			cur.Next = cur.Next.Next
// 			pre = pre.Next
// 		}

// 		if !find {
// 			pre = cur
// 		}
// 		cur = cur.Next
// 	}

// 	pre.Next = bigHead

// 	return dummy.Next
// }



type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

type LRUCache struct {
	size int
	capability int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}


func initDLinkNode(key , value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkNode(0, 0),
		tail: initDLinkNode(0, 0),
		capability: capacity,
	}
	l.head.next = l.tail
	l.tail.next = l.head
	return l
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.cache[key]
	if !ok {
		node = initDLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capability {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}


func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ne := len(equations)
	v := make([]bool, ne)
	vr := make([]bool, ne)

	nq := len(queries)

	re := make([]float64, nq)

	for rei := range re {
		re[rei] = float64(-1)
	}


	// 给方程组中的每个变量编号
    id := map[string]int{}
    for _, eq := range equations {
        a, b := eq[0], eq[1]
        if _, has := id[a]; !has {
            id[a] = len(id)
        }
        if _, has := id[b]; !has {
            id[b] = len(id)
        }
    }

	path := []int{}
	rpath := []int{}
	find := false
	var dfs func(start, end string) 
	dfs = func(start, end string) {
		if find {
			return
		}

		if start == end && len(path)>0 {
			rpath = append([]int{}, path...)
			find = true
			return
		}

		// if v[idx] {
		// 	return
		// }

		
		for i := 0; i < ne; i++ {
			if !v[i] {
				if start == equations[i][0] {
					v[i] = true
					path = append(path, (i+1))
					dfs(equations[i][1], end)
					path = path[:len(path)-1]
					v[i] = false
				}
			}
		}
		
		for i := 0; i < ne; i++ {
			if !vr[i] {
				if start == equations[i][1] {
					vr[i] = true
					path = append(path, (i+1)*-1)
					dfs(equations[i][0], end)
					path = path[:len(path)-1]
					vr[i] = false
				}
			}
		}
		
		

	}

	for i := 0; i < ne; i++ {
		// for j := range 2 {
			for qi , q := range queries {
				_, hasS := id[q[0]]
        		_, hasE := id[q[1]]

				if !hasS || !hasE {
					continue
				}

				if re[qi] != float64(-1) {
					continue
				}

				if q[0] == equations[i][0] || q[0] == equations[i][1] {
					dfs(q[0], q[1])
				}
				

				if find {
					ans := float64(1)
					for _, p := range rpath {
						if p > 0 {
							ans *= values[p-1]
						} else {
							ans *= 1/values[p*-1-1]
						}
					}
					re[qi] = ans
					find = false
					rpath = []int{}
				}
			}
		// }
	}

	return re
}


func findMin(nums []int) int {
	left := 0
	right := len(nums)
	n := len(nums)
	for left < right {
		mid := (left+right)/2

		if mid>=1 && mid<n-1 && nums[mid] < min(nums[mid-1], nums[mid+1]) {
			return nums[mid]
		}

		if nums[mid] > nums[right-1] {
			left = mid+1
		} else {
			right = mid
		}
	}
	return min(nums[left], min(nums[0] ,nums[len(nums)-1]))
}



func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	li, lj := 0, 0
	ri, rj := n-1, m-1

	for li < ri || (li == ri && lj < rj) {
		total := (ri-li)*m - lj + (rj+1)
		println(total)
		mi := li + ((lj+total/2)/m)
		mj := (lj + (total/2)) % m

		println(mi, ":", mj, ":", matrix[mi][mj])
		if matrix[mi][mj] == target {
			return true
		}

		if matrix[mi][mj] < target {
			li = mi+(mj+1)/m
			lj = (mj+1)%m
		} else {
			ri = mi+(mj+m-1)/m-1
			rj = (mj+m-1)%m
		}


	}

	if li>=n {
		return false
	}

	return matrix[li][lj] == target
}

func main() {
	searchMatrix([][]int{{1}, {3}}, 0)
	// searchMatrix([][]int{{1,1}}, 2)
	searchMatrix([][]int{{1,3,5,7}, {10,11,16,20}, {23,30,34,50}}, 20)
	// findMin([]int{2,3,4,5,1})
	// calcEquation([][]string{{"a", "e"}, {"b", "e"}}, []float64{2.0, 3.0}, [][]string{{"e", "e"}})
}