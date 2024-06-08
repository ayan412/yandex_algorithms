package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, n, m int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

	min1 := a*(n-1) + n
	max1 := a*(n-1) + n + 2*a

	min2 := b*(m-1) + m
	max2 := b*(m-1) + m + 2*b

	fmt.Println(min1,min2,max1,max2)

	minAns := math.Max(float64(min1), float64(min2))
	maxAns := math.Min(float64(max1), float64(max2))

	if maxAns < minAns {
		fmt.Println(-1)
	} else {
		fmt.Printf("%.f %.f\n", minAns, maxAns)
	}
}

// т.е. выводятся зн-я в рамках которых вып-ся усл-е по увиденным поездам.
