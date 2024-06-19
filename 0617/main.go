package main


func findLUSlength(strs []string) int {
	r := make(map[int]map[string]int)

	count := len(strs)
	for i := 0; i < count; i++ {
		if _, ok := r[len(strs[i])]; !ok {
			r[len(strs[i])] = make(map[string]int)
		}
		r[len(strs[i])][strs[i]]++
	}

	pre := 11
	for i := 10; i >= 1; i-- {
		if _, ok := r[i]; ok{
			for k0, v := range r[i] {
				dup := false
				if len(r[pre]) > 0 {
					for k := range r[pre] {
						if check(k, k0) {
							dup = true
						}
					}
				}
				if v == 1 && !dup {
					return i
				}
			}
			pre = i
		}
	}
	return -1
}

func check(a string, b string) bool {
	na := len(a)
	nb := len(b)
	i:=0
	j:=0
	for i < nb && j < na {
		if b[i] == a[j] {
			i++
		}
		j++
	}
	if i == nb {
		return true
	} else {
		return false
	}
}

func main() {
	println(findLUSlength([]string{"abaa","abaa","eaec","eaec","eae","z"}))
}