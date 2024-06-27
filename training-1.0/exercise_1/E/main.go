package main

import (
	"fmt"
)

func main() {
	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)

	if p2 == 1 && n2 == 1 {
		if k1 <= k2 {
			fmt.Println(1, 1)
		} else if k1 <= k2*m {
			fmt.Println(1, 0)
		} else if m == 1 {
			fmt.Println(0, 1)
		} else {
			fmt.Println(0, 0)
		}
	} else {
		// Adjustments to ensure calculations are correct
		minLvl := (k2 + (p2-1)*m + n2 - 1) / ((p2-1)*m + n2)
		maxLvl := (k2 - 1) / ((p2-1)*m + n2 - 1)
		if minLvl > maxLvl || n2 > m {
			fmt.Println(-1, -1)
		} else {
			maxP1 := (k1 - 1) / (m * maxLvl) + 1
			minP1 := (k1 - 1) / (m * minLvl) + 1
			maxN1 := (k1 - 1) / maxLvl % m + 1
			minN1 := (k1 - 1) / minLvl % m + 1
			if maxP1 != minP1 {
				fmt.Println(0, 0)
			} else {
				// Ensure both outputs are considered
				if maxN1 != minN1 {
					fmt.Println(maxP1, 0)
				} else {
					fmt.Println(maxP1, maxN1)
				}
			}
		}
	}
}



// package main

// import "fmt"

// // Panic handler
// func handlePanic() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Panic handling")
// 	}
// }

// func findPN(K1, M, K2, P2, N2 int) (int, int) {
// 	defer handlePanic()

// 	var P1, N1, ke, kp int

// 	if N2 > M || K2 < N2 {
// 		return -1, -1
// 	}

// 	if N2 == 1 && P2 == 1 {
// 		if K1 > K2 {
// 			if K1 <= M && M != 1 {
// 				return 1, 0
// 			} else if K1 > M && M == 1 {
// 				return 0, 1
// 			}
// 		} else {
// 			return 1, 1
// 		}
// 	}

// 	// для поиска кол-ва квартир на одном этаже в подъезде
// 	div := (P2-1)*M + N2 - 1

// 	// Количество квартир на одном этаже в подъезде
// 	if N2 == 1 && P2 != 1 {
// 		ke = K2
// 	} else {
// 		ke = K2 / div
// 	}
	
// 	//Но for N = 1 особый случай

// 	// Количество квартир в одном подъезде
// 	kp = ke * M

// 	// проверка корректности данных
// 	// k2v := M*ke
// 	// if k2v > K2 {
// 	// 	return -1, -1
// 	// }

// 	P1 = ((K1 - 1) / kp) + 1
// 	N1 = (((K1 - 1) % kp) / ke) + 1
// 	return P1, N1
// }

// func main() {
// 	// var K1, M, K2, P2, N2 int = 3, 2, 2, 2, 1
// 	var K1, M, K2, P2, N2 int
// 	fmt.Scan(&K1, &M, &K2, &P2, &N2)

// 	P1, N1 := findPN(K1, M, K2, P2, N2)
// 	fmt.Println(P1, N1)
// }
