package main

func minimumSteps(s string) int64 {
	lastZero := -1

	step := int64(0)
	for i := range s {
		if s[i] == '0' {
			lastZero++
			if lastZero != i {
				step += int64(i-lastZero)
			}
		}
	}

	return int64(step)
}

func main() {

}