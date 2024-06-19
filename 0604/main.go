package main


// type node struct {
// 	val int
// 	child []*node
// }

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges)
	r := make([][]int, n+1)
	rw := make([][]int, n+1)
	f := make([]bool, n+1)
	for i := range rw {
		rw[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		r[edges[i][0]] = append(r[edges[i][0]], edges[i][1])
		r[edges[i][1]] = append(r[edges[i][1]], edges[i][0])
		rw[edges[i][0]][edges[i][1]], rw[edges[i][1]][edges[i][0]] = edges[i][2], edges[i][2]
	}
	ans := []int{}
	path := []int{}
	var dfs func(node, w, l int) int
	dfs = func(node, w, l int) (ct int) {
		if f[node] {
			return
		}
		if w !=0 && w %signalSpeed == 0 {
			ct++
		}
		f[node] = true
		cc := 0
		for i := 0; i < len(r[node]); i++ {
			cc = dfs(r[node][i], w+rw[node][r[node][i]], l+1)
			ct += cc
			if l == 0 {
				path = append(path, cc)	
			}	
		}

		// if len(r[node]) == 1 {
		// 	if cc >0 {
		// 		path = append(path, cc)	
		// 	}
		// }

		return
	}

	for i := 0; i <= n; i++ {
		ct := dfs(i, 0, 0)
		// println(i, ":")
		// for _, v := range path {
		// 	println(v)
		// }
		t := 0
		for _, v := range path {
			t += (ct-v)*v
		}
		path = []int{}
		ans = append(ans, t/2)
		// count = 0
		f = make([]bool, n+1)
	}

	return ans
}

func main() {
	countPairsOfConnectableServers([][]int{{0,6,3},{6,5,3},{0,3,1},{3,2,7},{3,1,6},{3,4,2}}, 3)
}