package main

import (
	"fmt"
	"os"
)

func main() {
	const MAXN = 1000000
	var vp, vn, p1, n1, k1, m, k2, p2, n2 int64

	// Открываем файлы для чтения и записи
	in, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Считываем входные данные
	fmt.Fscanf(in, "%d %d %d %d %d", &k1, &m, &k2, &p2, &n2)

	// Инициализация переменных
	vp = 0
	vn = 0
	p1 = -1
	n1 = -1

	// Основной расчет
	for i := int64(1); i <= MAXN; i++ { // Пусть в каждом подъезде i квартир
		inp := m * i                   // Тогда всего квартир в подъезде
		tp := ((k2 - 1) / inp) + 1     // Предполагаемый номер подъезда p1
		tk := k2 - (tp-1)*inp          // Предполагаемый номер квартиры k1
		tn := ((tk - 1) / i) + 1       // Предполагаемый номер этажа n1

		if tp == p2 && tn == n2 { // Если совпали этаж и подъезд для второй квартиры...
			tp = ((k1 - 1) / inp) + 1
			tk = k1 - (tp-1)*inp
			tn = ((tk - 1) / i) + 1

			if n1 == -1 { // Если совпадение первое
				n1 = tn // Номер квартиры мы уже знаем
				vn = 1
			}
			if p1 == -1 { // Аналогично - номер подъезда
				p1 = tp
				vp = 1
			}
			if vp > 0 && tp != p1 {
				vp++
			}
			if vn > 0 && tn != n1 {
				vn++
			}
		}
	}

	// Сохраняем результат
	if vp == 0 {
		fmt.Fprint(out, "-1 ")
	}
	if vp > 1 {
		fmt.Fprint(out, "0 ")
	}
	if vp == 1 {
		fmt.Fprint(out, p1, " ")
	}
	if vn == 0 {
		fmt.Fprint(out, "-1")
	}
	if vn > 1 {
		fmt.Fprint(out, "0")
	}
	if vn == 1 {
		fmt.Fprint(out, n1)
	}
}


// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var k1, m, k2, p2, n2 int
// 	fmt.Scan(&k1, &m, &k2, &p2, &n2)

// 	if p2 == 1 && n2 == 1 {
// 		if k1 <= k2 {
// 			fmt.Println(1, 1)
// 		} else if k1 <= k2*m {
// 			fmt.Println(1, 0)
// 		} else if m == 1 {
// 			fmt.Println(0, 1)
// 		} else {
// 			fmt.Println(0, 0)
// 		}
// 	} else {
// 		// Adjustments to ensure calculations are correct
// 		minLvl := (k2 + (p2-1)*m + n2 - 1) / ((p2-1)*m + n2)
// 		maxLvl := (k2 - 1) / ((p2-1)*m + n2 - 1)
// 		if minLvl > maxLvl || n2 > m {
// 			fmt.Println(-1, -1)
// 		} else {
// 			maxP1 := (k1 - 1) / (m * maxLvl) + 1
// 			minP1 := (k1 - 1) / (m * minLvl) + 1
// 			maxN1 := (k1 - 1) / maxLvl % m + 1
// 			minN1 := (k1 - 1) / minLvl % m + 1
// 			if maxP1 != minP1 {
// 				fmt.Println(0, 0)
// 			} else {
// 				// Ensure both outputs are considered
// 				if maxN1 != minN1 {
// 					fmt.Println(maxP1, 0)
// 				} else {
// 					fmt.Println(maxP1, maxN1)
// 				}
// 			}
// 		}
// 	}
// }



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
