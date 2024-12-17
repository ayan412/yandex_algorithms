package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 2048 * 1024
	buf := make([]byte, 0, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	str1 := strings.Fields(scanner.Text())

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	str2 := strings.Fields(scanner.Text())

	slcN := make([]int, N)
	slcM := make([]int, M)

	for i := 0; i < N; i++ {
		slcN[i], _ = strconv.Atoi(str1[i])
	}

	for i := 0; i < M; i++ {
		slcM[i], _ = strconv.Atoi(str2[i])
	}

	// Находим минимальную разницу
	i, j := 0, 0
	minDiff := 2000000
	var resN, resM int

	// Цикл по 2-м указателям в 2-х срезах
	for i < N && j < M {
		// разница между футболкой и штанами
		diff := slcN[i] - slcM[j]
		absDiff := diff
		// Абсолютное зн-ие (без минуса)
		if diff < 0 {
			absDiff = -diff // умнож-е на -1 чтобы получить положит-е
		}

		if absDiff < minDiff {
			minDiff = absDiff
			resN = slcN[i]
			resM = slcM[j]
		}

		// Двигаем указатели на основе diff
		if diff < 0 { // отриц-е, которое будет дальше увел-ся (в абсолютном зн-и) если идти по j++, поэтому след-я итерация i
			i++
		} else if diff > 0 { // положит-е, которое будет умень-ся (в абсолютном зн-и) если идти по i++
			j++
		} else {
			break // елси разница уже равно ноль
		}
	}

	fmt.Printf("%d %d\n", resN, resM)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	//var N, M int

// 	scanner.Scan()
// 	strN := scanner.Text()
// 	//fmt.Println(strN)
// 	N, _ := strconv.Atoi(strN)

// 	scanner.Scan()
// 	str1 := strings.Fields(scanner.Text())

// 	scanner.Scan()
// 	strM := scanner.Text()
// 	M, _ := strconv.Atoi(strM)

// 	scanner.Scan()
// 	str2 := strings.Fields(scanner.Text())

// 	// fmt.Println(N, M)
// 	// fmt.Println(str1, str2)

// 	var slcN []int

// 	for i := 0; i < N; i++ {
// 		n, err := strconv.Atoi(str1[i])
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		slcN = append(slcN, n)
// 	}

// 	//fmt.Println(slcN)

// 	var slcM []int

// 	for i := 0; i < M; i++ {
// 		m, err := strconv.Atoi(str2[i])
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		slcM = append(slcM, m)
// 	}

// 	//fmt.Println(slcM)

// 	mindiff := 100000000
// 	result := [2]int{11,21}

// 	i, j := N-1, M-1

// 	for i >= 0 && j >= 0 {
// 		//fmt.Println(i, j)
// 		curdiff := slcN[i] - slcM[j]

// 		if slcN[i] > slcM[j] {

// 			if curdiff < mindiff {
// 				mindiff = curdiff
// 				result[0] = slcN[i]
// 				result[1] = slcM[j]
// 				// fmt.Println("---")
// 				i--
// 			}
// 		} else if slcN[i] == slcM[j] {
// 			mindiff = curdiff
// 			result[0] = slcN[i]
// 			result[1] = slcM[j]
// 			//	fmt.Println("+++")
// 			i--
// 			j--
// 		} else {
// 			curdiff = -curdiff
// 			if curdiff < mindiff {
// 				mindiff = curdiff
// 				result[0] = slcN[i]
// 				result[1] = slcM[j]
// 			}
// 			j--
// 		}
// 	}

// 	fmt.Printf("%d %d\n", result[0], result[1])
// }
