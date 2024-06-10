package main

import (
	"fmt"
	//"math"
)

func main() {
	var n1, m1, n2, m2 int

	_, err := fmt.Scanf("%d %d %d %d", &n1, &m1, &n2, &m2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Первый вариант
	tm1 := max(n1, m1) + max(n2, m2)
	tn1 := max(min(n1, m1), min(n2, m2))
	res1 := tn1 * tm1

	// Второй вариант
	tm2 := min(n1, m1) + min(n2, m2)
	tn2 := max(max(n1, m1), max(n2, m2))
	res2 := tn2 * tm2

	// Третий вариант
	tm3 := max(n1, m1) + min(n2, m2)
	tn3 := max(min(n1, m1), max(n2, m2))
	res3 := tn3 * tm3

	// Четвертый вариант
	tm4 := min(n1, m1) + max(n2, m2)
	tn4 := max(max(n1, m1), min(n2, m2))
	res4 := tn4 * tm4

	// Находим минимальную площадь
	minArea := res1
	minTN, minTM := tn1, tm1

	if res2 < minArea {
		minArea = res2
		minTN, minTM = tn2, tm2
	}

	if res3 < minArea {
		minArea = res3
		minTN, minTM = tn3, tm3
	}

	if res4 < minArea {
		minArea = res4
		minTN, minTM = tn4, tm4
	}

	fmt.Printf("%d %d\n", minTN, minTM)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
