package main

import (
	"slices"
	"sort"
)

func jobScheduling2(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	ss := make([][2]int, n)
	rr := make([][2]int, n)
	for i := 0; i < n; i++ {
		ss[i][0] = startTime[i]
		ss[i][1] = i

		rr[i][0] = endTime[i]
		rr[i][1] = i
	}

	sort.SliceStable(ss, func(i, j int) bool {
		if ss[i][0] == ss[j][0] {
			return ss[i][1] < ss[j][0]
		}
		return ss[i][0] < ss[j][0]
	})

	sort.SliceStable(rr, func(i, j int) bool {
		if rr[i][0] == rr[j][0] {
			return rr[i][1] < rr[j][0]
		}
		return rr[i][0] < rr[j][0]
	})

	dp := make([]int, n)
	dpE := make([]int, n)
	// dp[0] = profit[ss[0][1]]
	ans := dp[0]
	for i := 0; i < n; i++ {
		// preEnd := sort.Search(n, func(i int) bool {
		// 	return rr[i][0] >= (ss[i][0]-1)
		// })
		preEnd := binarySearch(rr, ss[i][0]-1)
		if rr[preEnd][0] > ss[i][0] {
			dp[i] = max(dp[i], profit[ss[i][1]])
			dpE[preEnd] = max(dpE[preEnd], dp[i])
		} else {
			for  {
				dp[i] = max(dp[i], dpE[preEnd] + profit[ss[i][1]])
				dpE[preEnd] = max(dpE[preEnd], dp[i])
				if preEnd+1 < n && rr[preEnd][0] == rr[preEnd+1][0] {
					preEnd++
				} else {
					break
				}
			}
		}


		// for {
		// 	dp[i] = max(dp[i], dp[preEnd]+ profit[ss[i][1]])
		// 	if preEnd+1 < n && rr[preEnd][0] == rr[preEnd+1][0] {
		// 		preEnd++
		// 	} else {
		// 		break
		// 	}
		// }
		ans = max(dp[i], ans)
	}

	return ans
}

func jobScheduling(startTime, endTime, profit []int) int {
	n := len(startTime)
	type job struct {s, e, p int}
	jobs := make([]job, n)
	for i, start := range startTime {
		jobs[i] = job{start, endTime[i], profit[i]}
	}

	// 按结束时间排序
	slices.SortFunc(jobs, func (a, b job) int {
		return a.e - b.e
	})

	dp := make([]int, n+1)
	for i, job := range jobs {
		j := sort.Search(i, func(j int) bool {
			return jobs[j].e > job.s
		})

		// 第i+1的最大收益
		// 1. 不选 i+1，那么就是 i 的收益
		// 2. 如果选 i+1，那么就是 要找到 j 结束时间小于 i 开始时间的 最大收益 + 当前收益
		dp[i+1] = max(dp[i], dp[j]+job.p)
	}

	return dp[n]
}


func binarySearch(rr [][2]int, target int) int {
	left, right := 0, len(rr)-1
	for left <= right {
		mid := left + (right-left)/2
		if rr[mid][0] < target {
			left = mid + 1
		} else if rr[mid][0] > target {
			right = mid - 1
		} else {
			return mid // 找到目标值，返回索引
		}
	}
	return left // 未找到目标值，返回-1
}



func merge(intervals [][]int) [][]int {
	n := len(intervals)
	ans := [][]int{}

	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < n; i++ {
		if intervals[i][0] <= ans[len(ans)-1][1] {
			if intervals[i][1] > ans[len(ans)-1][1] {
				ans[len(ans)-1][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return ans
}


func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	t := make([]int, n*2)

	if n == 0 {
		return [][]int{newInterval}
	}

	for i := 0; i < n; i++ {
		t[i*2] = intervals[i][0]
		t[i*2+1] = intervals[i][1]
	}

	l := sort.Search(n*2, func(i int) bool {
		return t[i] >= newInterval[0]
	})

	r := sort.Search(n*2, func(i int) bool {
		return t[i] >= newInterval[1]
	})

	// println(l ,r)

	ans := [][]int{}
	for i := 0; i < n*2; i+=2 {
		lt := t[i]
		rt := t[i+1]
		// 中间插入
		if l == r && l == i && newInterval[1]!=t[i] {
			ans = append(ans, []int{newInterval[0], newInterval[1]})
		}
		if r == 0 && newInterval[1] == t[i] {
			lt = newInterval[0]
		}

		if l != r {
			if l == i {
				lt = newInterval[0]
			}
			if l == i || l == i+1 {
				if r >= n*2 {
					rt = newInterval[1]
					i = r
				} else if r&1==0 {
					if t[r] == newInterval[1] {
						rt = t[r+1]
						i = r
					} else {
						rt = newInterval[1]
						i = r-2
					}
				} else {
					rt = t[r]
					i = r-1
				}
			}
		}

		// 跳过数
		ans = append(ans, []int{lt, rt})
	}

	if l >= n*2 {
		ans = append(ans, newInterval)
	}

	return ans
}


func findMinArrowShots2(points [][]int) int {
	n := len(points)

	sort.SliceStable(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	ans := 0
	for i := 0; i < n; {
		ans++
		// lmax := points[i][0]
		rmin := points[i][1]
		for j := i+1; j<n; {
			if points[j][0] <= rmin {
				rmin = min(rmin, points[j][1])
				i = j
				j++
			} else {
				break
			}
		}
		i++
	}

	return ans
}



func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
    maxRight := points[0][1]
    ans := 1
    for _, p := range points {
        if p[0] > maxRight {
            maxRight = p[1]
            ans++
        }
    }
    return ans
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sums := 0
	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}

		v := s*10 + node.Val
		dfs(node.Left, v)
		dfs(node.Right, v)

		if node.Left == nil && node.Right == nil {
			sums += v
		}
	}
	dfs(root, 0)
	return sums
}


func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	q := []*TreeNode{root}
	ans := []int{}
	count := len(q)
	for count > 0 {
		cur := q[0]
		q = q[1:]
		if cur.Left != nil {
			q = append(q, cur.Left)
		}

		if cur.Right != nil {
			q = append(q, cur.Right)
		}

		count--

		if count == 0 {
			ans = append(ans, cur.Val)
			count = len(q)
		}
	}
	return ans
}


func main() {
	// l = 1 r = 2
	findMinArrowShots([][]int{{10,16}, {2,8}, {1,6}, {7,12}})
}