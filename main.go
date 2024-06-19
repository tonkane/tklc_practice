package main

import (
	// "go/printer"
	// "runtime/trace"
	"fmt"
	"math"
	"sort"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ThroneInheritance struct {
	Name      string
	ChildList map[string][]string
	DeathList map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		Name:      kingName,
		ChildList: make(map[string][]string),
		DeathList: make(map[string]bool),
	}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.ChildList[parentName] = append(t.ChildList[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.DeathList[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	ans := make([]string, 0)
	var dfs func(string)
	dfs = func(name string) {
		if !t.DeathList[name] {
			ans = append(ans, name)
		}
		for _, child := range t.ChildList[name] {
			dfs(child)
		}
	}
	dfs(t.Name)
	return ans
}

func minOperations(nums []int) int {
	sort.Ints(nums)
	len := len(nums)
	minRes := len
	start := 1
	// 去重?
	for i := 1; i < len; i++ {
		if nums[i] != nums[i-1] {
			nums[start] = nums[i]
			start++
		}
	}
	for i := 0; i < start; i++ {
		j := sort.Search(start, func(k int) bool {
			return nums[k] > nums[i]+len-1
		})
		minRes = min(minRes, len-j+i)
	}
	return minRes
}

func minOperations2(nums []int) int {
	sort.Ints(nums)
	len := len(nums)
	minRes := len
	start := 1
	// 去重?
	for i := 1; i < len; i++ {
		if nums[i] != nums[i-1] {
			nums[start] = nums[i]
			start++
		}
	}
	// for i:=0; i<start; i++ {
	// 	j := sort.Search(start, func(k int) bool {
	// 		return nums[k] > nums[i]+len-1
	// 	})
	// 	minRes = min(minRes, len-j+i)
	// }
	j := 0
	for index, leftV := range nums {
		rightV := leftV + len - 1

		for j < start && nums[j] <= rightV {
			minRes = min(minRes, len-(j-index+1))
			j++
		}
	}
	return minRes
}

// 二分查找函数
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // 计算中间索引，防止溢出
		if arr[mid] == target {
			return mid // 找到目标值，返回索引
		} else if arr[mid] < target {
			left = mid + 1 // 如果中间值小于目标值，更新左边界
		} else {
			right = mid - 1 // 如果中间值大于目标值，更新右边界
		}
	}
	return -1 // 未找到目标值，返回-1
}

func maximumCount(nums []int) int {
	// 2分法找0出现的位置
	// 注意可能出现多个0
	ans := 0
	len := len(nums)
	left, right := 0, len-1
	zeroIndex := -1
	// 预先判断一下头尾
	if nums[left] > 0 || nums[right] < 0 {
		return len
	}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == 0 {
			zeroIndex = mid
			break
		} else if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// print(zeroIndex)
	// 如果找不到 zeroIndex
	if zeroIndex == -1 {
		// println("l:", left)
		// println("r:", right)

		ans = max(left, len-right-1)
	} else {
		zeroLeft := zeroIndex
		zeroRight := zeroIndex
		for i := zeroIndex; i >= 0; i-- {
			if nums[i] == 0 {
				zeroLeft = i
			}
		}
		for j := zeroIndex; j < len; j++ {
			if nums[j] == 0 {
				zeroRight = j
			}
		}
		println("left:", zeroLeft)
		println("right:", zeroRight)
		// println()
		ans = max(zeroLeft, (len - zeroRight - 1))
	}

	return ans
}

func maximumBinaryString(binary string) string {
	ans := ""
	len := len(binary)
	s := []rune(binary)
	firstZeroIndex := -1
	zeroNums := 0
	// 不用全遍历
	// 其实只要知道首个0位置 和 后面1的数量就行了
	for i := len - 1; i >= 0; i-- {
		if s[i] == '0' {
			firstZeroIndex = i
			zeroNums++
		}
	}

	if zeroNums <= 1 {
		return binary
	} else {
		// 这里可以用两个repeat 优化？
		ans = strings.Repeat("1", len)
		ans = ans[:firstZeroIndex+zeroNums-1] + "0" + ans[firstZeroIndex+zeroNums:]
	}

	return ans
}

// 预处理，找出 1-50 中 与 i 互质的元素
const maxValue = 51

var clist [maxValue][]int

func init() {
	for i := 1; i < maxValue; i++ {
		for j := 1; j < maxValue; j++ {
			if gcd(i, j) == 1 {
				clist[i] = append(clist[i], j)
			}
		}
	}
}

// 辗转相除法
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	type pair struct {
		depth int
		id    int
	}

	valDepthId := [maxValue]pair{}
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		val := nums[x]
		maxDepth := 0
		for _, j := range clist[val] {
			p := valDepthId[j]
			if p.depth > maxDepth {
				maxDepth = p.depth
				ans[x] = p.id
			}
		}

		tmp := valDepthId[val]

		valDepthId[val] = pair{depth: depth, id: x}
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
			}
		}
		valDepthId[val] = tmp
	}
	dfs(0, -1, 1)
	return ans
}

func findChampion(grid [][]int) int {
	n := len(grid)
	ans := 0
	for i := 0; i < n; i++ {
		ans = i
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && i != j {
				break
			}
			if j == n-1 {
				return ans
			}
		}
	}
	return ans
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		dfs(target, index+1)

		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index)
			comb = comb[:len(comb)-1] // 回退最后一个元素
		}
	}
	dfs(target, 0)
	return

}

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	var dfs func(target, index int)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	comb := []int{}
	dfs = func(target, index int) {
		if index >= 10 {
			return
		}

		if target == 0 && len(comb) == k {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		dfs(target, index+1)

		if target-nums[index] >= 0 {
			comb = append(comb, nums[index])
			dfs(target-nums[index], index+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(n, 0)
	return ans
}


func combinationSum5(nums []int, target int) int {
	ans := 0
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func (target, idx int)  {

		if target == 0 {
			for _,value := range comb {
				print(value, ",")
			}
			println("")
			println("")
			ans++
			return
		}

		if idx >= len(nums) {
			return
		}

		dfs(target, idx+1)

		if target - nums[idx] >= 0 {
			comb = append(comb, nums[idx])
			dfs(target - nums[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}

	dfs (target, 0)
	return ans
}

func combinationSum4(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _,x := range nums {
			if x <= i {
				f[i] += f[i-x]
			}
		}
	}
	return f[target]
}

// 爬楼梯
// 状态转移 f(n) = f(n-1) + f(n-2)
// 边际 f(1) = 1 f(2) = 2
// 初始值 0 
func climbStairs(n int) int {
	dp := make([]int, n+2)
	dp[1] = 1
	dp[2] = 2
	for i :=3; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 杨辉三角
// 状态转移 f(n) = [f(n-1)[0] .... f(n-1)[n-1]]
// 边际 f(1) = [1] f(2) = [1,1]
// 
func generate(numRows int) [][]int {
	dp := make([][]int, numRows+1)
	dp[1] = make([]int, 1)
	dp[1][0] = 1
	for i := 2; i <= numRows; i++ {
		dp[i] = make([]int, i)
		dp[i][0] = 1
		dp[i][i-1] = 1
		for j := 1; j < i-1; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp[1:]
}


func getRow(rowIndex int) []int {
	dp1 := make([]int, 0)
	dp2 := make([]int, 0)
	dp1 = append(dp1, 1)
	// dp2 = append(dp2, []int{1,1}...)
	if rowIndex == 0 {
		return dp1
	}
	for i:=1; i<=rowIndex; i++ {
		dp2 = make([]int, i+1)
		dp2[0] = 1
		dp2[i] = 1
		for j := 1; j < i; j++ {
			dp2[j] = dp1[j-1] + dp1[j]
		}
		dp1 = append([]int{}, dp2...)
	}

	return dp2
}

// 最大利润
// 状态转移 最大利润 = 当前最大 - 前一个最小
// 边际 f(max) = f(0) f(min) = f(0)

// 其实只要用第I天价格和前面最小价格相减，再取和前个利润相比的最大最即可
// 最后再更新最小价格
func maxProfit(prices []int) int {
	maxProfit := 0
	curMax := prices[0]
	preMin := prices[0]
	curMin := prices[0]
	for i := 1; i < len(prices); i++ {
		if preMin > curMin {
			if prices[i] - curMin >= maxProfit {
				preMin = curMin
				curMax = prices[i]
				maxProfit = curMax - preMin
			}
		}

		if prices[i] < curMin {
			curMin = prices[i]
		}

		if prices[i] > curMax {
			curMax = prices[i]
			maxProfit = curMax - preMin
		}
		
	}
	return maxProfit
}

func countBits(n int) []int {
	max := 17
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = 0
		tmp := i
		for j := 0; j < max; j++ {
			if tmp & 1 == 1 {
				ans[i]++
			}
			tmp = tmp >> 1
		}
	}

	return ans
}


func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	// 计算固定满意的人数
	len := len(customers)
	maxNums := 0
	for i := 0; i < len; i++ {
		if (grumpy[i] == 0) {
			maxNums += customers[i]
		}
	}
	sNums := maxNums
	// 窗口值为 minutes
	// 初始
	tmp := 0
	for j := 0; j < minutes; j++ {
		// 新增的人
		if (grumpy[j] == 1) {
			tmp += customers[j]
		}
	}
	maxNums = max(maxNums, sNums+tmp)
	// 滑动
	for i := 1; (i+minutes-1) < len; i++ {
		if (grumpy[i+minutes-1] == 1) {
			tmp += customers[i+minutes-1]
		}
		if (grumpy[i-1] == 1) {
			tmp -= customers[i-1]
		}
		maxNums = max(maxNums, sNums+tmp)
	}
	return maxNums

}


func isValidBST(root *TreeNode) bool {
	ans := true
	var dfs func(node *TreeNode)
	list := make([]int, 0)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		list = append(list, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(list); i++ {
		if list[i-1] >= list[i] {
			ans = false
			break
		} 
	}
	return ans
}


func recoverTree(root *TreeNode)  {
	list := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return 
		}
		dfs(node.Left)
		list = append(list, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	sort.Ints(list)
	var dfs2 func(node *TreeNode)
	idx :=0
	dfs2 = func (node *TreeNode)  {
		if node == nil {
			return 
		}
		dfs2(node.Left)
		node.Val = list[idx]
		idx++
		dfs2(node.Right)
	}
	dfs2(root)
}


func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var dfs func(node *TreeNode, sum int)
	tmp := make([]int, 0)
	count := 0
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		tmp = append(tmp, node.Val)
		count+=node.Val
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		if count == sum && node.Left == nil && node.Right == nil && len(tmp)>0{
			ans = append(ans, append([]int{}, tmp...))
		}
		if len(tmp)>0 {
			count -= tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(root, targetSum)
	return ans
}

func flatten(root *TreeNode)  {
	var dfs func(node *TreeNode)
	list := make([]*TreeNode, 0)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(list); i++ {
		list[i-1].Left = nil
		list[i-1].Right = list[i]
	}
}

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}


func connect(root *Node) *Node {
	var queue1 = make([]*Node, 0)
	var queue2 = make([]*Node, 0)
	if root == nil {
		return root
	}
	queue1 = append(queue1, root)
	for {
		if queue1[0].Left != nil {
			queue2 = append(queue2, queue1[0].Left)
		}
		if queue1[0].Right != nil {
			queue2 = append(queue2, queue1[0].Right)
		}

		// 指向
		if len(queue1) > 1 {
			queue1[0].Next = queue1[1]
		}

		// 移除一个
		queue1 = queue1[1:]

		if len(queue2) == 0 && len(queue1) == 0{
			break
		}

		// 如果 queue1 都空了，queue2 赋值给 1 2置为空
		if len(queue1) == 0 {
			queue1 = append([]*Node{}, queue2...)
			queue2 = []*Node{}
		}
	}
	return root
}



func amountOfTime1(root *TreeNode, start int) int {
	ans := 0
	queue1, queue2 := []*TreeNode{}, []*TreeNode{}
	leftDepth := 0
	rightDepth := 0
	lr := 0
	startDepth := 0
	if root.Left != nil {
		queue1 = append([]*TreeNode{}, root.Left)
		for len(queue1) > 0 {
			
			if queue1[0].Left != nil {
				queue2 = append(queue2, queue1[0].Left)
			}
			if queue1[0].Right != nil {
				queue2 = append(queue2, queue1[0].Right)
			}

			if queue1[0].Val == start {
				startDepth = leftDepth+1
			}

			queue1 = queue1[1:]

			if len(queue1) == 0 {
				leftDepth++
				if len(queue2) > 0 {
					queue1 = append([]*TreeNode{}, queue2...)
					queue2 = []*TreeNode{}
				}
			}
		}
	}
	
	if root.Right != nil {
		queue1 = append([]*TreeNode{}, root.Right)
		for len(queue1) > 0 {
			
			if queue1[0].Left != nil {
				queue2 = append(queue2, queue1[0].Left)
			}
			if queue1[0].Right != nil {
				queue2 = append(queue2, queue1[0].Right)
			}

			if queue1[0].Val == start {
				startDepth = rightDepth+1
				lr = 1
			}

			queue1 = queue1[1:]

			if len(queue1) == 0 {
				rightDepth++
				if len(queue2) > 0 {
					queue1 = append([]*TreeNode{}, queue2...)
					queue2 = []*TreeNode{}
				}
			}
		}
	}

	// 在左边
	if lr == 0 {
		ans = max(max(startDepth, leftDepth-startDepth), rightDepth+startDepth)
	} else {
		ans = max(leftDepth+startDepth, max(startDepth, rightDepth-startDepth))
	}
	fmt.Println("lr", lr)
	fmt.Println("startDepth", startDepth)
	fmt.Println("leftDepth", leftDepth)
	fmt.Println("rightDepth", rightDepth)
	return ans
}


func amountOfTime(root *TreeNode, start int) int {
	ans := 0
	var dfs func(node *TreeNode, parent *TreeNode)
	list := make(map[int][]int, 0)
	dfs = func(node *TreeNode, parent *TreeNode) {
		if node == nil {
			return
		}

		if parent != nil {
			list[node.Val] = append(list[node.Val], parent.Val)
		}

		if node.Left != nil {
			list[node.Val] = append(list[node.Val], node.Left.Val)
		}

		if node.Right != nil {
			list[node.Val] = append(list[node.Val], node.Right.Val)
		}
		
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	visit := [100001]int{}
	
	var dfs2 func(idx int, depth int)
	dfs2 = func(idx int, depth int) {
		if visit[idx] == 1 {
			return
		}
		visit[idx] = 1
		for _, value := range list[idx] {
			dfs2(value, depth+1)
			ans = max(ans, depth)
		}
	}
	dfs2(start, 0)
	return ans
}



func search(nums []int, target int) int {
	left := 0
	right := len(nums)-1
	var mid int
	ans := -1
	for left <= right {
		// if nums[left] == target {
		// 	return left
		// }
		// if nums[right] == target {
		// 	return right
		// }
		if target > right && target < left {
			break
		}
		mid = (left + right) / 2
		if nums[mid] == target {
			ans = mid
			break
		}
		if (nums[mid] > target && nums[left] < target) || 
		(nums[mid] < nums[right] && nums[right] < target ||
		(nums[mid] > target && nums[right]>nums[mid])){
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}


func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums)
	var mid int
	start, end := -1, -1
	for left < right {
		mid = (left + right) / 2

		if nums[mid] == target {
			for i := mid; i >= 0; i-- {
				if nums[i] == target {
					start = i
				}
			}
			for j := mid; j < len(nums); j++ {
				if nums[j] == target {
					end = j
				}
			}
			break
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return []int{start, end}
}


func distanceTraveled(mainTank int, additionalTank int) int {
	ans := 0
	for i := 0; i <= mainTank; {
		if i+5 <= mainTank {
			ans+=50
			if additionalTank > 0 {
				additionalTank--
				mainTank++
			}
		} else {
			ans+=(mainTank-i)*10
		}
		i+=5
	}
	return ans
}

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	// 首先给数组排个序
	sort.Ints(nums)
	len := len(nums)
	i := 0
	pre := 100001
	for ;i < len-2; i++ {
		if nums[i] == pre {
			continue
		}
		// 最小的数都大于0了，后续的数肯定大于0
		if nums[i] > 0 {
			break
		}
		j := i+1
		k := len-1

		for j < k {
			if nums[i] + nums[j] + nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
			if nums[j] + nums[k] > nums[i]*-1 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
		pre = nums[i]
	}
	return ans
}



func threeSumClosest(nums []int, target int) int {
	ans := target
	// 首先给数组排个序
	sort.Ints(nums)
	len := len(nums)
	i := 0
	pre := 100001
	absValue := 100001

	var abs func(x int) int
	abs = func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}

	for ;i < len-2; i++ {
		if nums[i] == pre {
			continue
		}

		j := i+1
		k := len-1

		for j < k {
			// 找到就直接退出
			sums := nums[i] + nums[j] + nums[k]
			if sums == target {
				return target
			}
			absCur := abs(sums-target)
			if  absCur < absValue {
				ans = sums
				absValue = absCur
			}
			// 如果和大于目标值
			if sums > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
		pre = nums[i]
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	// 首先给数组排个序
	sort.Ints(nums)
	len := len(nums)
	if len < 4 {
		return ans
	}
	i := 0
	pre := math.MaxInt
	for ;i < len-3; i++ {
		pre2 := math.MaxInt
		if nums[i] == pre {
			continue
		}
		// 如果 i i+1 i+2 i+3 的和都大于 target 则返回  
		if nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target {
			break
		}
		y := i+1
		for ;y < len-2; y++ {
			j := y+1
			k := len-1
			if nums[y] == pre2 {
				continue
			}
			for j < k {
				sums := nums[i] + nums[j] + nums[k] + nums[y]
				if sums == target {
					ans = append(ans, []int{nums[i], nums[y], nums[j], nums[k]})
					// for j < k && nums[k] == nums[k-1] {
					// 	k--
					// }
					// for j < k && nums[j] == nums[j+1] {
					// 	j++
					// }
				}
				if sums > target {
					k--
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else {
					j++
					for j < k && nums[j] == nums[j-1] {
						j++
					}
				}
			}
			pre2 = nums[y]
		}
		pre = nums[i]
	}
	return ans
}



func main() {
	// res := minOperations([]int{1,2,3,4,5})
	// res := maximumCount([]int{0})
	neg := sort.SearchInts([]int{2,3,2,4,5}, 2)
	// print(neg)
	// print(res)
	println((3+4)/2)
	search([]int{3,5,1}, 3)
	fourSum([]int{1,0,-1,0,-2,2}, 0)
}
