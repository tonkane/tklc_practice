package main

import (
	"slices"
	"sort"
)

func findMinimumTime(tasks [][]int) int {
	// 按结束时间排序
	slices.SortFunc(tasks, func(a, b []int) int {
		return a[1] - b[1]
	})
	// 开始时间 结束时间 到目前的总时间
	type tuple struct{l, r, s int}
	// 哨兵
	h := []tuple{{-1, -1, 0}}
	for _, task := range tasks {
		start, end, d := task[0], task[1], task[2]

		idx := sort.Search(len(h), func(i int) bool {
			return h[i].l >= start
		}) - 1

		// 当前时间 - 最大不包含区间的时间 = 可重用时间
		d -= h[len(h)-1].s - h[idx].s

		// 最大不包含区间的时间 可能一半不包含一半包含，有可能有可重用时间
		// 当不包含区间结束时间大于当前开始时间说明有可重用时间
		if start <= h[idx].r {
			d -= h[idx].r - start + 1 
		}

		if d <= 0 {
			continue
		}

		// 结束 - 持续 是否会影响最后一个的结束时间
		// 如果影响了就需要更新最后一个的结束时间和持续时间
		for end-h[len(h)-1].r <= d {
			top := h[len(h)-1]
			h = h[:len(h)-1]
			d += top.r - top.l + 1
		}
		
		h = append(h, tuple{end-d+1, end, h[len(h)-1].s + d})
	}

	return h[len(h)-1].s
}

func main() {
	findMinimumTime([][]int{{1,3,2}, {2,5,3}, {3,6,2}})
}