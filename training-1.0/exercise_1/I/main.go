package main

import (
	"fmt"
)

// Пузырьковая с-ка - текущий < последующего
func sortBuble(first, second int) (a, b *int) {
	if first < second {
		return &first, &second
	}
	return &second, &first
}

func main() {

	var a,b,c,d,e int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &d)
	fmt.Scanf("%d", &e)

	aa, bb := sortBuble(a, b)
	bb, _ = sortBuble(*bb, c)
	aa, bb = sortBuble(*aa, *bb)
	dd, ee := sortBuble(d, e)

	if *bb <= *ee && *aa <= *dd {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}