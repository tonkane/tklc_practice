package main


func maximumLength(s string) int {
	n := len(s)
	left := 0
	right := 0
	ans := -1

	m := [27][2]int{}
	size := 0
	for right < n {
		for right < n && s[left] == s[right] {
			size = right - left + 1
			right++
		}
		
		// 如果超过3则至少 有N-2满足要求
		// if size >= 3 {
		ans = max(ans, size-2)
		// }
		
		// size 和 过往记录比较
		if size > m[s[left]-'a'][0] {
			if m[s[left]-'a'][0] != 0 {
				ans = max(ans, m[s[left]-'a'][0])
			}
			m[s[left]-'a'][1] = m[s[left]-'a'][0]
			m[s[left]-'a'][0] = size
		} else if size >  m[s[left]-'a'][1] {
			if m[s[left]-'a'][0] - size <= 1 {
				ans = max(ans, m[s[left]-'a'][0]-1)
			}
			m[s[left]-'a'][1] = size
		} else if size == m[s[left]-'a'][1] {
			ans = max(ans, size)
		}

		// right++
		left = right
	}

	if ans == 0 {
		return -1
	}
	return ans
}


func main() {
	maximumLength("cbc")
}