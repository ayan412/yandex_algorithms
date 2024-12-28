// https://contest.yandex.ru/contest/27794/problems/B/
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

	const maxCapacity = 1024 * 1024
	buf := make([]byte, 0, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	nk := strings.Fields(scanner.Text())

	N, _ := strconv.Atoi(nk[0])

	K, _ := strconv.Atoi(nk[1])

	var slcN []int

	scanner.Scan()

	str1 := strings.Fields(scanner.Text())

	for i := 0; i < N; i++ {
		n, _ := strconv.Atoi(str1[i])
		slcN = append(slcN, n)
	}

	var sum, count int
	left, right := 0, 0

	for right < N {

		sum = sum + slcN[right]

		// двигаем левый указ-ль до правого указ-ля, чтобы найти на окне K
		for sum > K && left <= right { // левый не должен обгонять правого чтобы окно не схлопнулось
			sum = sum - slcN[left] // сумма без тек-го указ-ля
			left++                 // сдвиг указ-ля
		}

		if sum == K {
			count++
		}
		
		// значит sum<K и просто двигаем правый указ-ль
		right++
	}

	// когда правый указ-ль дойдет до конца, то двигаем только левый
	for left < N {
		sum = sum - slcN[left]
		left++
		if sum == K {
			count++
		}

	}

	fmt.Println(count)

}
