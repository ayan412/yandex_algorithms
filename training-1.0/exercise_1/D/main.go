package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	Result(a, b, c)
}

func Result(a, b, c int) {
	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}

	c2 := c * c

	if a == 0 {
		if b == c2 {
			fmt.Println("MANY SOLUTIONS")
		} else {
			fmt.Println("NO SOLUTION")
		}
		return
	}

	// (c^2 - b) should be divisible by a
	if (c2 - b) % a != 0 {
		fmt.Println("NO SOLUTION")
		return
	}

	// Calculate x and check if it's an integer
	x := (c2 - b) / a
	fmt.Println(x)
}
