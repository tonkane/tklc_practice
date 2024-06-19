package main

import (
	"slices"
	"sort"
)

func maxIncreasingCells(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	r := map[int][][2]int{}

	// 把相同的数放一起
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r[mat[i][j]] = append(r[mat[i][j]], [2]int{i, j})
		}
	}

	k := make([]int, 0, len(r))
	for i := range r {
		k = append(k, i)
	}
	// 排序
	sort.Ints(k)

	// 行记录
	rows := make([]int, m)
	// 列记录
	cols := make([]int, n)

	for _, x := range k {
		postions := r[x]
		c := make([]int, len(postions))
		for i, pos := range postions {
			// 每个相同数的最大值
			c[i] = max(rows[pos[0]], cols[pos[1]]) + 1
		}

		// 相同数值更新需要暂存避免影响一下个的更新
		for i, pos := range postions {
			rows[pos[0]] = max(rows[pos[0]], c[i])
			cols[pos[1]] = max(cols[pos[1]], c[i])
		}
	}

	return slices.Max(rows)
}

func main() {
	maxIncreasingCells([][]int{{3, 1, 6}, {-9,5,7}})
}
