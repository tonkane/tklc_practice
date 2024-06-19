package main

// import "sort"

// func mostCompetitive(nums []int, k int) []int {
//     st := nums[:0] // 把 nums 当作栈
//     for i, x := range nums {
//         for len(st) > 0 && x < st[len(st)-1] && len(st)+len(nums)-i > k {
//             st = st[:len(st)-1]
//         }
//         if len(st) < k {
//             st = append(st, x)
//         }
//     }
//     return st
// }



func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	n := len(nums)
	for i, v := range nums {
		for len(stack) > 0 && v < stack[len(stack)-1] && len(stack)+n-i > k {
			stack =  stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, v)
		}
	}
	return stack
}


func main() {
	mostCompetitive([]int{1, 4, 1, 1, 5, 4, 9, 6}, 4)
	// mostCompetitive([]int{11,52,57,91,47,95,86,46,87,47,70,56,54,61,89,44,3,73,1,7,87,48,17,25,49,54,6,72,97,62,16,11,47,34,68,58,14,36,46,65,2,15}, 18)
}
