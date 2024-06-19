package main

import (
	"slices"
	"sort"
)

// 树状数组
type treeArr []int

func (ta treeArr) add(i int) {
	// 树状数组，更新 i 及 i + lowbit 位置的 数据
	for ; i < len(ta); i += i & -i {
		ta[i]++
	}
}

// 返回 [1 - i] 的元素和
func (ta treeArr) pre(i int) (res int) {
	// 移除最后一个1 找到下一个2次幂
	for ; i > 0; i &= i - 1 {
		res += ta[i]
	}
	return
}

func resultArray(nums []int) []int {
	ans := []int{}

	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	// 对连续出现的元素去重
	sorted = slices.Compact(sorted)
	m := len(sorted)

	// 排序完成数组
	a := []int{nums[0]}
	b := []int{nums[1]}

	// 维护两个树状数组
	t1 := make(treeArr, m+1)
	t2 := make(treeArr, m+1)

	t1.add(sort.SearchInts(sorted, nums[0]) + 1)
	t2.add(sort.SearchInts(sorted, nums[1]) + 1)

	for _, x := range nums[2:] {
		// 快速找到元素出现的位置
		v := sort.SearchInts(sorted, x) + 1
		// 找大于 v 的元素的数量，即全部 - 小于等于v 的元素数量
		gc1 := len(a) - t1.pre(v) 
		gc2 := len(b) - t2.pre(v)

		if gc1 > gc2 || gc1 == gc2 && len(a) <= len(b) {
			a = append(a, x)
			t1.add(v)
		} else {
			b = append(b, x)
			t2.add(v)
		}
	}

	ans = append(a, b...)
	return ans
}

func main() {
	slices.Compact()
}
