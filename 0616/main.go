package main

import "sort"

func countCompleteDayPairs2(hours []int) int {
	n := len(hours)

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (hours[i]+hours[j])%24 == 0 {
				ans++
			}
		}
	}

	return ans
}

func countCompleteDayPairs(hours []int) int64 {
	n := len(hours)
	r := [24]int64{}
	ans := int64(0)

	for i := 0; i < n; i++ {
		r[hours[i]%24]++
	}

	// 0 == C 2 X
	// 12 == C 2 X
	// 1 == 23
	ans += r[0] * (r[0] - 1) / 2
	ans += r[12] * (r[12] - 1) / 2
	for i := 1; i < 12; i++ {
		ans += r[i] * r[24-i]
	}

	return ans
}

func maximumTotalDamage(power []int) int64 {
	sort.Ints(power)
	n := len(power)
	

	r := [][2]int{}
	r = append(r, [2]int{power[0], 1})
	for i := 1; i < n; i++ {
		if power[i] == r[len(r)-1][0] {
			r[len(r)-1][1]++
		} else {
			r = append(r, [2]int{power[i], 1})
		}
	}

	nr := len(r)

	dp := make([]int64, nr)
	ans := int64(0)
	for i := 0; i < nr; i++ {
		dp[i] = int64(r[i][0]) * int64(r[i][1])
		j := i-1
		count := 0
		for j>=0 {
			c := r[i][0] - r[j][0]
			if c >2 {
				dp[i] = max(dp[j] + int64(r[i][0]) * int64(r[i][1]), dp[i])
				count++
			}
			if count >=3 {
				break
			}
			j--
		}  
		ans = max(dp[i], ans)
	}

	return ans
}

// 线段树节点结构体
type SegmentTreeNode struct {
	Start, End   int // 区间的起始和结束位置
	Sum          int // 当前区间的和
	Left, Right *SegmentTreeNode // 左右子节点
}

// 创建线段树的辅助函数
func buildSegmentTree(nums []int, start, end int) *SegmentTreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &SegmentTreeNode{Start: start, End: end, Sum: nums[start]}
	}
	mid := (start + end) / 2
	node := &SegmentTreeNode{
		Start: start,
		End:   end,
		Left:  buildSegmentTree(nums, start, mid),
		Right: buildSegmentTree(nums, mid+1, end),
	}
	node.Sum = node.Left.Sum + node.Right.Sum
	return node
}

// 线段树查询区间和
func (node *SegmentTreeNode) Query(l, r int) int {
	if node == nil || l > node.End || r < node.Start {
		return 0
	}
	if l <= node.Start && r >= node.End {
		return node.Sum
	}
	return node.Left.Query(l, r) + node.Right.Query(l, r)
}

// 线段树更新节点值
func (node *SegmentTreeNode) Update(index int, val int) {
	if node == nil {
		return
	}
	if node.Start == index && node.End == index {
		node.Sum = val
	} else {
		// mid := (node.Start + node.End) / 2
		node.Left.Update(index, val)
		node.Right.Update(index, val)
		node.Sum = node.Left.Sum + node.Right.Sum
	}
}

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)
	r := make([]int, n)

	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			r[i] = 1
		}
	}

	sgTree := buildSegmentTree(r, 0, n-1)
	ans := []int{}
	for i := 0; i < len(queries); i++ {
		// 查询操作
		if queries[i][0] == 1 {
			// 不包含区间
			sgTree.Query(queries[i][1]+1, queries[i][2]-1)
		}
		// 更新操作
		if queries[i][0] == 2 {
			// 检查3个元素
			nums[queries[i][1]] = queries[i][2]

			if queries[i][1] == 0 {
				if nums[1] > nums[0] && nums[1] > nums[2] {
					r[1] = 1
					sgTree.Update(1, 1)
				}
			} else if queries[i][1] == n-1 {
				if nums[n-2] > nums[n-1] && nums[n-2] > nums[n-3] {
					r[n-2] = 1
					sgTree.Update(n-2, 1)
				}
			} else {

			}

			// sgTree.Update()
		}
	}
	return ans
}


func findLUSlength(a string, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}

func main() {
	maximumTotalDamage([]int{2,9,10})
}
