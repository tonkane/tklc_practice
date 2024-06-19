package main

import "sort"

func findWinners(matches [][]int) [][]int {
	r := make(map[int]int)

	for _, v := range matches {
		if _, ok := r[v[0]]; !ok {
			r[v[0]] = 0
		}
		r[v[1]]++
	}

	// ans := [][]int{}
	a0 := []int{}
	a1 := []int{}
	for k, v := range r {
		if v == 0 {
			a0 = append(a0, k)
		}
		if v == 1 {
			a1 = append(a1, k)
		}
	}
	sort.Ints(a0)
	sort.Ints(a1)
	return [][]int{a0, a1}
}