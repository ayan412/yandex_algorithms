package main 

import (
	"fmt"
)

func main() {
	var a,b,c int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	if a+b>c && b+c>a && c+a>b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}


