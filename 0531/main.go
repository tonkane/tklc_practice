package main


func findMissingAndRepeatedValues(grid [][]int) []int {
	sums := 0
	n := len(grid[0])

	m := make([]bool, n*n)
	ans := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[grid[i][j]-1] {
				ans = append(ans, grid[i][j])
			} else {
				sums += grid[i][j]
			}
			m[grid[i][j]-1] = true
		}
	}
	ans = append(ans, (1+n*n)*n*n/2-sums)
	return ans
}

func main() {

}