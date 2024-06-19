package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])
	matrixLen := m + n - 1
	matrix := make([][]int, matrixLen)
	ans := make([][]int, m)

	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := getIdx(i, j, n)
			matrix[idx] = append(matrix[idx], mat[i][j])
		}
	}

	for i := 0; i < matrixLen; i++ {
		sort.Ints(matrix[i])
		for idx := 0; idx < len(matrix[i]); idx++ {
			if i < n {
				ans[idx][idx+i] = matrix[i][idx]
			} else {
				ans[i-n+1+idx][idx] = matrix[i][idx]
			}
		}
	}

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		idx := getIdx(i, j, n)
	// 		ans[i][j] = matrix[idx][min(i,j)]
	// 	}
	// }

	return ans
}

func getIdx(i, j, n int) int {
	if j-i >= 0 {
		return j - i
	} else {
		return (i - j) + n - 1
	}
}

type ListNode struct {
    Val int
    Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	var sums int
	var plus int
	var point *ListNode
	point = ans
	for l1 != nil || l2 != nil {
		if l1 == nil {
			sums = l2.Val + plus
		} else if l2 == nil {
			sums = l1.Val + plus
		} else {
			sums = l1.Val + l2.Val + plus
		}
		if sums >= 10 {
			sums -= 10
			plus = 1
		} else {
			plus = 0
		}

		next := &ListNode{Val: sums, Next: nil}
		point.Next = next
		point = next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if plus == 1 {
		next := &ListNode{Val: 1, Next: nil}
		point.Next = next
	}
	return ans.Next
}

// 全排列
// 有重复数字，返回不重复的排列
func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	length := len(nums)
	used := [9][21]bool{}
	used2 := make([]bool, length)
	path := []int{}
	// sort.Ints(nums)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx >= length {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < length; i++ {
			if !used[idx][nums[i]+10] && !used2[i] {
				path = append(path, nums[i])
				used2[i] = true
				used[idx][nums[i]+10] = true
				dfs(idx+1)
				used[idx+1] = [21]bool{}
				used2[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}


// 90 度旋转矩阵
// 4个角  (a,b) (c,d) (e,f) (g,h)
func rotate(matrix [][]int)  {
	n := len(matrix)
	distance := n-1
	for i := 0; i < n/2; i++ {
		a, b := i, i
		c, d := i, n-1-i
		e, f := n-1-i, i
		g, h := n-1-i, n-1-i

		for j := distance; j>0; j--{
			// 交换 ab - cd
			matrix[a][b], matrix[c][d] = matrix[c][d], matrix[a][b]
			// 交换 ab - gh
			matrix[a][b], matrix[g][h] = matrix[g][h], matrix[a][b]
			// 交换 ab - ef
			matrix[a][b], matrix[e][f] = matrix[e][f], matrix[a][b]

			// 移动位置
			b++
			c++
			e--
			h--
		}

		distance -= 2
	}
}


func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	mapList := make(map[string]int)

	ansIndex := 0
	for i := 0; i < len(strs); i++ {
		tmp := sortString(strs[i])
		idx, ok := mapList[tmp]
		if !ok {
			mapList[tmp] = ansIndex
			ansIndex++
			ans = append(ans, []string{strs[i]})
		} else {
			ans[idx] = append(ans[idx], strs[i])
		}
	}

	return ans
}

func sortString(s string) string {
	r := []rune(s)

	sort.SliceStable(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}




func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])

	i , j := 0, 0

	ans := make([]int, m*n)
	idx := 0
	top, bottom, left, right := 0, m-1, 0, n-1
	status := 0
	for j < right || i < bottom || j > left || i > top {
		ans[idx] = matrix[i][j]

		if status == 0 {
			if j < right {
				j++
				idx++
			} else {
				status = 1
				top++
			}
		} else if status == 1 {
			if i < bottom {
				i++
				idx++
			} else {
				status = 2
				right--
			}
		} else if status == 2 {
			if j > left {
				j--
				idx++
			} else {
				status = 3
				bottom--
			}
		} else if status == 3 {
			if i > top {
				i--
				idx++
			} else {
				status = 0
				left++
			}
		}

		if idx >= m*n {
			break
		}
	}

	// if left == top {
	// 	ans[len(ans)-1] = matrix[left][top]
	// }
	if idx < m*n {
		ans[idx] = matrix[i][j]
	}
	return ans
}

func main() {
	// spiralOrder([][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}})

	spiralOrder([][]int{{6,9,7}})
	// spiralOrder([][]int{{1,2,3}, {4,5,6}, {7,8,9}})
}