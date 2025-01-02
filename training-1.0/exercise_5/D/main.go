// https://contest.yandex.ru/contest/27794/problems/D/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Считываем n
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// и затем r
	scanner.Scan()
	r, _ := strconv.Atoi(scanner.Text())

	// расстония до улицы
	distance := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		distance = append(distance, num)
	}

	// два указателя
	left, right := 0, 1
	counter := 0

	// пока не дойдем до правого крайнего зн-я
	for right < n {
		diff := distance[right] - distance[left]
		if diff > r {
			counter += n - right
			left++ // двигаем левый когда нашли зн-е > r, иначе r++
		} else {
			right++
		}
	}

	fmt.Println(counter)

}
