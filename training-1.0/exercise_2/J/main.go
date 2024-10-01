package main

import (
	"fmt"
	"math"
)

// сравнить каждую ноту с нотой до неё и вывести границы нот "послушав" все знач-я нот.
func main() {
	var prev, cur float64
	var L, R float64 = 30, 4000
	var s string
	var n int

	// кол-во нот
	fmt.Scanf("%d", &n)
	// самая первая "предыдущая" нота
	fmt.Scanf("%f", &prev)

	for i := 0; i < n-1; i++ {
		fmt.Scanf("%f %s", &cur, &s)
		// (cur+prev)/2 - середина отрезка
		if s == "closer" {
			// сравнение каждой текущей ноты с предыдущей
			if cur > prev {
				L = math.Max(L, (cur+prev)/2)
			} else {
				R = math.Min(R, (cur+prev)/2)
			}
		} else {

			if cur > prev {
				R = math.Min(R, (cur+prev)/2)
			} else {
				L = math.Max(L, (cur+prev)/2)
			}
		}
		// текущая теперь как предыдущая
		prev = cur
		//fmt.Println(L,R)

	}
	// точность 6 цифр в дробной части
	fmt.Printf("%.6f %.6f\n", L, R)
}
