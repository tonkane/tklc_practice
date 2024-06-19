package main


func distributeCandies(n int, limit int) int {
	if n > limit*3 {
		return 0
	}
	ans := 0

	var dfs func(count, r int)
	dfs = func (count, r int)  {
		if r == 0 && count == 0 {
			ans++
		}

		if count == 0 {
			return
		}

		for i := 0; i <= limit; i++ {
			if r-i >= 0 {
				// println(count-1, ":", r-i)
				dfs(count-1, r-i)
			}
		}
	}
	dfs(3, n)
 	return ans
}


func main() {
	distributeCandies(3,3)
}