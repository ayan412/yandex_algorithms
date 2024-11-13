package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
сделать мапу с ключем по ширине, и знач-м высоты блока. map[ширина]высота
для каждой ширины хранить максимальную высоту и в конце суммировать все высоты в мапе,
в которой будут храниться уже все допустимые по условию
(неявная проверка на: ширина верхнего блока строго меньше ширины нижнего) задачи параметры блоков.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024) // 64KB
	scanner.Buffer(buf, 1024*1024)  //  1MB

	// кол-во пар
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	unitMap := make(map[int]int)

	// скан-м размеры блоков с преобр-м в числа
	for i := 0; i < n; i++ {
		scanner.Scan()
		part := strings.Split(scanner.Text(), " ")
		width, _ := strconv.Atoi(part[0])
		height, _ := strconv.Atoi(part[1])
		//fmt.Printf("%d:%d\n", width, height)

		// проверка уже сохран-го знач-я curHeight в мапе по ключу width - (запятая, ок)
		// если знач-я нет (!ok) ИЛИ оно МЕНЬШЕ чем текущее зн-е, то добав-м в мапу текущее (большее).
		if curHeight, ok := unitMap[width]; !ok || height > curHeight {
			unitMap[width] = height
		}
	}

	// fmt.Println("unitMap", unitMap)

	// складываем значения мапы т.е. разл-е высоты блоков
	sum := 0
	for _, v := range unitMap {
		sum += v
	}

	fmt.Println(sum)

}
