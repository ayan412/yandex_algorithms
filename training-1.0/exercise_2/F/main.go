package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)
	sequence := make([]int, size)
	reverse := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&sequence[i])
		reverse[size-1-i] = sequence[i]
	}

	cycleSequences(sequence, reverse)

}

func cycleSequences(sequence, reverse []int) {

	var reversePointer *[]int = &reverse // в reversePointer передан обратный слайс

	for i := 0; i < len(sequence); i++ {
		if sequence[i] != (*reversePointer)[i] {
			reversePointer = fillShifted(reversePointer)
			//fmt.Println(reversePointer)
			break
		}
	}

	for i := 1; i < len(sequence); i++ {
		if sequence[i] != (*reversePointer)[i] {
			//fmt.Printf("iteration:%d\n", i)
			reversePointer = fillShifted(reversePointer)
		}

		if sequence[i] == (*reversePointer)[i] && i == len(sequence)-1 {
			result := (*reversePointer)[i+1:]
			fmt.Print(len(result), "\n")
			for i, v := range result {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(v)
			}
			fmt.Println()
		}
	}
}

func fillShifted(reversePointer *[]int) (r *[]int) {
	shifted := make([]int, len(*reversePointer)+1, len(*reversePointer)+1)
	for i := len(*reversePointer); i > 0; i-- {
		shifted[i] = (*reversePointer)[i-1]
	}
	r = &shifted
	return r
}
