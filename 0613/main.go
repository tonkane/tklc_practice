package main

import "sort"

func findMaximumElegance(items [][]int, k int) int64 {
	// 利润从大到小排序
	sort.SliceStable(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] < items[j][1]
		}
		return items[i][0] > items[j][0]
	})

	// 保存类别出现次数
	r := make(map[int]int, 0)

	// k长度队列
	q := make([][]int, k)

	// 最大不同序列数
	maxDis := k
	r2 := make(map[int]int, 0)
	for i := range items {
		r2[items[i][1]]++
	}

	maxDis = min(maxDis, len(r2))

	// 初始化
	total_profit := 0
	for i := 0; i < k; i++ {
		q[i] = items[i]
		total_profit += items[i][0]
		r[items[i][1]]++
	}

	// 如果长度一致直接返回
	if len(r) == k {
		return int64(total_profit) + int64(k)*int64(k)
	}

	ck := len(r)
	res := int64(total_profit) + int64(ck)*int64(ck)

	// 如果长度不一致，尝试从队尾开始替换
	start := k - 1
Loop1:
	for i := k; i < len(items); i++ {
		// 已经有了，不要
		if r[items[i][1]] > 0 {
			continue
		}
		ck := len(r)
		for j := start; j >= 0; j-- {
			// 首个多次出现的数
			if r[q[j][1]] > 1 {
				
				// 无论如何都超过不了
				if q[j][0]-items[i][0] > (maxDis)*(maxDis)-ck*ck {
					break Loop1
				} else {
					// 替换
					r[q[j][1]]--
					r[items[i][1]] = 1
					total_profit -= q[j][0] - items[i][0]
					q[j] = items[i]
					start = j - 1
					res = max(res, int64(total_profit) + int64(ck+1)*int64(ck+1))
					break
				}
			}
		}
	}
	
	return res
}

func main() {
	findMaximumElegance([][]int{{11, 8}, {31, 5}, {30, 1}, {19, 5}, {32,3}, {6,2}, {1,2}, {35,5}, {33,1}}, 6)
	findMaximumElegance([][]int{{2, 5}, {2, 2}, {7, 5}, {2, 4}, {6, 5}}, 2)
	// findMaximumElegance([][]int{{10,1},{10,1},{10,1},{10,1},{10,1},{10,1},{10,1},{10,1},{10,1},{10,1},{3,2},{3,3},{3,4},{3,5},{3,6},{3,7},{3,8},{3,9},{3,10},{3,11}}, 10)
}
