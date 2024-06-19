package main

import (
	"fmt"
	"math"
)

func qs(n int, k int) int {
	ans := n/(k-1)
	return ans
}


func qc(n int, a []int) [501]bool {
	ans := [501]bool{}

	for i := 0; i < len(a); i++ {
		ans[a[i]] = true
	}

	return ans
}

func s2t(a string) int {
	ans := 0
	n := len(a)

	r := map[byte]int{}
	r['0'],r['1'],r['2'],r['3'],r['4'],r['5'],r['6'],r['7'],r['8'],r['9']=0,1,2,3,4,5,6,7,8,9
	r['A'],r['B'],r['C'],r['D'],r['E'],r['F']=10,11,12,13,14,15

	count:=0
	for i := n-1; i >=2; i-- {
		ans += r[a[i]] * int(math.Pow(float64(16), float64(count)))
		count++
	}

	return ans
}


func main() {
    a := 0
    // b := 0
	first := true
	qcn := 0
	qca := []int{}
    for {
        n, _ := fmt.Scan(&a)
        if n == 0 {
            break
        } else {
			if first {
				qcn = a
				first = false
			} else {
				qca = append(qca, a)
			}
            // fmt.Printf("%d\n", a + b)
        }
    }
	ans := qc(qcn, qca)
	for i , v := range ans {
		if v  {
			fmt.Printf("%d\n", i)
		}
	}

}