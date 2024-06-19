package main

import "math"

func cherryPickup(grid [][]int) int {
	// dp[i][j][k]
	// i行 j、k 列 (i,j) (i,k)
	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)

	for i := range dp {
		dp[i] = make([][]int, n+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+2)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt32
			}
		}
	}

	dp[0][1][n] = grid[0][0] + grid[0][n-1]

	ans := dp[0][1][n]
	for i := 1; i < m; i++ {
		// 9种组合？
		for j := 0; j <= min(i, n-1); j++ {
			for k := n - 1; k >= max(n-i-1, 0); k-- {
				// println(i,":",j,":",k)
				dp[i][j+1][k+1] = max(dp[i-1][j][k], dp[i-1][j][k+1], dp[i-1][j][k+2], dp[i-1][j+1][k], dp[i-1][j+1][k+1], dp[i-1][j+1][k+2], dp[i-1][j+2][k], dp[i-1][j+2][k+1], dp[i-1][j+2][k+2]) + grid[i][j] + grid[i][k]
				// println(dp[i][j+1][k])
				if j == k {
					dp[i][j+1][k+1] -= grid[i][j]
				}

				if i == m-1 {
					ans = max(ans, dp[i][j+1][k+1])
				}
			}
		}
	}
	return ans
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])

	a := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// v := make([][]bool, n)

	// for vi := range v {
	// 	v[vi] = make([]bool, m)
	// }

	q := [][2]int{}
	ans := 0

	for ii := 0; ii < n; ii++ {
		for jj := 0; jj < m; jj++ {
			if grid[ii][jj] == '1' {
				q = [][2]int{{ii, jj}}
				ans++
			}

			for len(q) > 0 {
				cur := q[0]
				q = q[1:]

				if grid[cur[0]][cur[1]] == '1' {
					for _, plus := range a {
						i := cur[0] + plus[0]
						j := cur[1] + plus[1]
						if i >= 0 && i < n && j >= 0 && j < m {
							if grid[i][j] == '1' {
								q = append(q, [2]int{i, j})
							}
						}
					}
				}
				// 已访问
				grid[cur[0]][cur[1]] = '0'
			}
		}
	}

	return ans
}

func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])

	// 如果 n 或 m 只有 一行保持原样
	if n==1 || m==1 {
		return
	}

	a := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// 只处理边缘
	i, j := 0, 0
	ai := 0
	q := [][2]int{}
	for ai < 4 {
		i += a[ai][0]
		j += a[ai][1]

		if i >= n || i < 0 || j >= m || j < 0 {
			i -= a[ai][0]
			j -= a[ai][1]
			ai++
		} else {
			// 边缘循环
			// println(i, ":", j)
			if board[i][j] == 'O' {
				q = [][2]int{{i, j}}
			}

			for len(q) > 0{
				cur := q[0]
				q = q[1:]

				if board[cur[0]][cur[1]] == 'O' {
					for _, plus := range a {
						ic := cur[0] + plus[0]
						jc := cur[1] + plus[1]
						if ic >= 0 && ic < n && jc >= 0 && jc < m {
							if board[ic][jc] == 'O' {
								q = append(q, [2]int{ic, jc})
							}
						}
					}
				}
				// 已访问
				board[cur[0]][cur[1]] = 'A'
			}
		}
	}

	// A -> O O -> X
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}

}

type Node struct {
    Val int
    Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	// 空节点的情况
	if node == nil {
		return nil
	}
	// 已访问创建的节点
    m := make(map[int]*Node)

	// 队列
	q := []*Node{node}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		newNode, ok := m[cur.Val]
		if !ok {
			m[cur.Val] = &Node{Val: cur.Val}
			newNode = m[cur.Val]
		}

		if cur.Neighbors != nil {
			for _, nei := range cur.Neighbors {
				neiNode, neiOk := m[nei.Val]
				if !neiOk {
					m[nei.Val] = &Node{Val: nei.Val}
					neiNode = m[nei.Val]
					q = append(q, nei)
				}
				newNode.Neighbors = append(newNode.Neighbors, neiNode) 
			}
		}
	}

	return m[1]
}


// 堆排序
func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	n := len(nums)
	for i := n-1; i >= n-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize) 
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2; i>=0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}


func main() {

	findKthLargest([]int{3,2,1,5,6,4}, 2)
}
