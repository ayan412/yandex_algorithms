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
	// Объем для 3*10^9
	const maxCapacity = 2048*2048
	buf := make([]byte, 0, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()

	gs := strings.Split(scanner.Text(), " ")
	g, s := gs[0], gs[1]
	gNum, _ := strconv.Atoi(g)
	sNum, _ := strconv.Atoi(s)
	//fmt.Println(gNum, sNum)

	// мапа для слова W
	W := make(map[string]int)

	// g сим-лы из которых состоит W
	scanner.Scan()
	gline := scanner.Text()

	// запись в мапу букв W
	for i := 0; i < gNum; i++ {
		W[string(gline[i])]++
	}

	// посл-ть симв-в S
	scanner.Scan()
	Sline := scanner.Text()

	// мапа для окна из послед-ти S
	S := make(map[string]int)

	// границы окна
	start, end := 0, gNum-1

	res := 0

	// запол-е окна длиной gNum-1 т.к. индек-я с 0
	for i := start; i <= end; i++ {
		S[string(Sline[i])]++
	}

	// fmt.Println(check(S, W))

	if check(S, W) == true {
		res++
	}

	// цикл пока не правая граница окна не дойдет до конца послед-ти S
	for end < sNum-1 {
		// проверка на выход за пределы границ S (ouf of range)
		if end == len(Sline) {
			break
		}

		// убрать ЗНАЧ-Е левой границы
		S[string(Sline[start])]--

		// если знач-е по ключу стало меньше или равно 0 (т.е. символ больше не встреч-ся), то удалить 
		if value, ok := S[string(Sline[start])]; ok && value <= 0 {
			delete(S, string(Sline[start]))
		}
		// сдвинуть левую гран-у
		start++
		// сдвинуть правую гран-у
		end++
		// добавить символ в мапу по новой гран-е
		S[string(Sline[end])]++
		// fmt.Println("in FOR:")
		// проверка на совпадение изменной мапы S с W
		if check(S, W) == true {
			res++
			// fmt.Println("res:", res)
		}
	}

	fmt.Println(res)

}

// ф-я для проверки совпадения ключей и их значений в обеих мапах 
func check(S, W map[string]int) bool {
	// fmt.Println(len(S),len(W))
	// fmt.Println(S,W)
	count := 0
	fl := false
	for k, v := range S {
		// если знач-я по ключу в W нет или оно не соот-т знач-ю по ключу в мапе S
		if curVofM, ok := W[k]; !ok || curVofM != v {
			//	fmt.Println("break")
			break
		} else {
			count++
		}
	}

	// своего рода проверка на совпадение ВСЕХ ключей и знач-й
	if count == len(S) {
		fl = true
		// fmt.Println("fl = true:", count)
	}

	return fl
}
