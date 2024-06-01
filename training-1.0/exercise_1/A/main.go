package main

import (
	"fmt"
)

func main() {
	var tRoom, tCond int
	var Mode string
	//fmt.Println("Enter temperature:")
	fmt.Scanf("%d %d", &tRoom, &tCond)
	//fmt.Println("Enter mode:")
	fmt.Scanf("%s", &Mode)

	t := showT(tRoom, tCond, Mode)
	fmt.Println(t)
	//fmt.Printf("%d %d %s", tRoom, tCond, Mode)


}

func showT(tR int, tC int, m string) int {
	// in case of tR=tC then:
	var t int = tR

	switch m {
	case "heat":
		if tR < tC {
			t = tC	
		}
		if tR > tC {
			t = tR
		}
	case "freeze":
		if tR > tC {
			t = tC
		}
		if tR < tC {
			t = tR
		}
	case "auto":
		if tR >= tC || tR <= tC {
			t = tC
		}
	case "fan": {
		t = tR
	}
	}
	return t
}
