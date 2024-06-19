package main

import (
	"strconv"

	"github.com/golang/mock/mockgen/model"
	// "github.com/go-playground/locales/lg"
)

func theMaximumAchievableX(num int, t int) int {
	return num + 2*t
}

func cntOfTrees(n int) int {
	if n == 0 {
		return 1
	}
	// 初始化结果为0
	result := 0
	// 计算卡特兰数的公式
	for i := 0; i < n; i++ {
		result += cntOfTrees(i) * cntOfTrees(n-1-i)
	}
	return result
}

func kawaiiStrings(n int) int {
	if n == 0 {
		return 1
	}
	// 初始化结果为0
	result := 0
	// 计算卡特兰数的公式
	for i := 0; i < n; i++ {
		result += kawaiiStrings(i) * kawaiiStrings(n-1-i)
	}
	return result
}

func perfectPair(arr []int) int {
	// write code here
	ans := 0
	n := len(arr)

	// 预处理 arr 检查 5 的数量和 2的数量
	r := make([][3]int, n)

	for i := 0; i < n; i++ {
		t5 := arr[i]
		r[i][2] = arr[i]
		for {
			if t5%5 == 0 {
				r[i][0]++
				r[i][2] /= 5
				t5 /= 5
			} else {
				break
			}
		}
		t2 := arr[i]
		for {
			if t2%2 == 0 {
				r[i][1]++
				r[i][2] /= 2
				t2 /= 2
			} else {
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 把两边5和2的数量加起来
			c5 := r[i][0] + r[j][0]
			c2 := r[i][1] + r[j][1]

			if r[i][2] > 10 || r[j][2] > 10 {
				continue
			}
			p := r[i][2] * r[j][2]
			if p > 10 {
				continue
			}
			if c5 == c2 && p <= 9 {
				ans++
			}
			if c5 == c2+1 && p <= 2 {
				ans++
			}
			if c2 > c5 && (c2-c5) <= 3 && 2<<(c2-c5-1)*p <= 9 {
				ans++
			}

		}
	}

	return ans
}

func fun(n int) int {
	ans := 0
	s := strconv.Itoa(n)
	lg := len(s)
	for i := lg - 1; i >= 0; i-- {
		if s[i] == '0' || s[i] == '5' {
			break
		}
		ans++
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 阶乘
func factorialRecursive(n int, m int) int {
	if n == 0 {
		return 1
	}
	return (n * factorialRecursive(n-1, m)) % m
}

func cntOfMethods(trees []*TreeNode) int {
	// 
	mod_num := 1000000007
	// 计算每个树最后一层的节点数目
	r := make([]int, len(trees))
	for i, tree := range trees {
		q := []*TreeNode{tree}
		r[i] = 2
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
				count = len(q)
				if count != 0 {
					r[i] = count*2
				}
			}
		}
	}

	sums := 1
	for _, num := range r {
		sums *= num
		// sums %= mod_num
	}
	
	ans := 0
	for _, num := range r {
		ans += sums/num
		// ans %= mod_num
	}
	return (factorialRecursive(len(trees)-1, mod_num) * ans) % mod_num
}

func main() {
	fun(52)
}
