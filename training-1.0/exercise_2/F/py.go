package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var seqN int
	fmt.Scan(&seqN)

	seq := make([]int, seqN)
	for i := 0; i < seqN; i++ {
		fmt.Scan(&seq[i])
	}

	left := 0
	right := seqN - 1
	add := 0

	for left < right {
		if seq[left] == seq[right] {
			left++
			right--
		} else {
			right = seqN - 1
			left++
			add = left
		}
	}

	fmt.Println(add)

	if add > 0 {
		result := make([]string, add)
		for i := 0; i < add; i++ {
			result[i] = strconv.Itoa(seq[add-1-i])
		}
		fmt.Println(strings.Join(result, " "))
	}
}
