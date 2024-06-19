package main

func longestAwesome2(s string) int {
	ans := 1

	r := make([]int, len(s))

	r[0] = 0 ^ (1 << (s[0] - '0'))

	for i := 1; i < len(s); i++ {
		r[i] = r[i-1] ^ (1 << (s[i] - '0'))

		if r[i] == 0 || (r[i]&(r[i]-1)) == 0 {
			ans = max(ans, i+1)
		}
	}

	for i := 0; i < len(r); i++ {
		for j := i - ans; j >= 0; j-- {
			t := r[i] ^ r[j]
			if t == 0 || (t&(t-1)) == 0 {
				ans = max(ans, i-j)
			}
		}
	}

	return ans
}

func longestAwesome(s string) (ans int) {
	const D = 10 // s 中的字符种类数
	n := len(s)
	pos := [1 << D]int{}
	for i := range pos {
		pos[i] = n // n 表示没有找到异或前缀和
	}
	pos[0] = -1 // i-j=ans(有前缀) 空前缀 因为 index-0是1前缀，空前缀 i-(-1) = ans
	pre := 0
	for i, c := range s {
		pre ^= 1 << (c - '0')
		for d := 0; d < D; d++ {
			ans = max(ans, i-pos[pre^(1<<d)]) // 奇数
		}
		ans = max(ans, i-pos[pre]) // 偶数
		if pos[pre] == n {         // 首次遇到值为 pre 的前缀异或和，记录其下标 i
			pos[pre] = i
		}
	}
	return
}

func main() {
	longestAwesome("000")
}
