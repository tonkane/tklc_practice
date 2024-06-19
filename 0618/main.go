package main

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	list :=  strings.Split(sentence, " ")
	n := len(list)
	for i := 0; i < n; i++ {
		if list[i][0] == '$' {
			if len(list[i]) > 1 {
				if nums, err := strconv.Atoi(list[i][1:]); err == nil {
					newnums := float64(nums) * (100 - float64(discount)) / 100
					list[i] = fmt.Sprintf("$%.2f", newnums)
				}
			}
		}
	}
	return strings.Join(list, " ")
}

func main() {
	str := discountPrices("there are $1 $2 and 5$ candies in the shop", 50)
	println(str)
}
