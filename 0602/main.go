package main

import (
	"container/heap"
	"sort"
	"strings"
)

func minimumChairs(s string) int {
	ans := 0

	tmp := 0
	for i := range s {
		if s[i] == 'E' {
			tmp++
		} else {
			tmp--
		}
		ans = max(ans, tmp)
	}

	return ans
}

func countDays(days int, meetings [][]int) int {
	sort.SliceStable(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		} 
		return meetings[i][0] < meetings[j][0]
	})

	n := len(meetings)

	s := meetings[0][0]
	e := meetings[0][1]
	ans := 0

	if n == 1 {
		return days-(e-s+1)
	}

	for i := 1; i < n; i++ {
		if e >= meetings[i][0] && e < meetings[i][1]{
			e = meetings[i][1]
		} else if meetings[i][1] > e {
			ans += (e-s)+1
			s = meetings[i][0]
			e = meetings[i][1]
		}
		if i == n-1 {
			ans += (e-s)+1
		}
	}
	// ans += (e-s)+1

	return days-ans
}

type Heap [][2]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i][0] == h[j][0] {
		return h[i][1] > h[j][1]
	} else {
		return h[i][0] < h[j][0]
	}
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Peek() interface{} {
	if h.Len() == 0 {
		return -1
	}
	return h[0]
}

var _ heap.Interface = &Heap{}

func clearStars(s string) string {
	h := &Heap{}

	r := []byte{}

	for i := range s {
		if s[i] == '*' &&  len(r)>0 {
			x := h.Peek().([2]int)
			r[x[1]] = ' '
			heap.Pop(h)
		} else {
			r = append(r, s[i])
			heap.Push(h, [2]int{int(s[i]-'a'), len(r)-1})
		}
	}

	return strings.ReplaceAll(string(r), " ", "")
}



func distributeCandies(candyType []int) int {
	t := make(map[int]bool)

	for i := range candyType {
		t[candyType[i]] = true
	}

	return min(len(t), len(candyType)/2)
}

func main() {
	countDays(6, [][]int{{1,6}})
	// countDays(10, [][]int{{5,7},{1,3},{9,10}})
	// countDays(57, [][]int{{1,48},{23,52},{21,56},{26,55}})
	// countDays(10, [][]int{{3,49},{23,44},{21,56},{26,55},{23,52},{2,9},{1,48},{3,31}})
}