// https://contest.yandex.ru/contest/27794/problems/F/

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

	// кол-во классов
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// кондеры с мин мощ-ю
	conditioner := make([]int, 1001)
	for i := 0; i < n; i++ {
		scanner.Scan()
		power, _ := strconv.Atoi(scanner.Text())
		conditioner[power]++
	}

	// кол-во кондеров
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// модели кондеров с мощ-ю и ценой
	models := make([]int, 1001)
	for i := 0; i < m; i++ {
		scanner.Scan()
		power, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		price, _ := strconv.Atoi(scanner.Text())
		// записываем макс зн-е мощн-и для каждой цены
		if power > models[price] {
			models[price] = power
		}
	}

	// fmt.Println(conditioner, models)

	sum := 0
	// отд-ый указ-ль для кондера и для модели
	ptrCond, ptrModel := 0, 0

	// 1-ый цикл по мин мощн-м кондеров

	/*
		Поскольку индекс в массиве models является ценой, и мы идем по нему последовательно от меньших значений к большим (ptrModel++),
		то первая подходящая мощность (models[ptrModel] >= needPower) будет иметь минимальную цену.
	*/
	for ptrCond < len(conditioner) {
		if conditioner[ptrCond] == 0 {
			ptrCond++
		} else {
			// опред-м какую мощн-ть кондера для какого конкр-о класса ищем среди моделей
			needPower, needNumber := ptrCond, conditioner[ptrCond] // needNumber - количество кондиционеров одинаковой мощности
			// 2-ой цикл по моделям кондеров
			for ptrModel < len(models) {
				// если мощн-ть в моделях < нужной (миним допустимой), идем по списку моделей дальше
				if models[ptrModel] < needPower {
					ptrModel++
				} else {
					// иначе
					price := ptrModel
					sum += price * needNumber // цена на количество кондиционеров одинаковой мощности
					// след-ий кондер
					ptrCond++
					// выход из 2-го цикла
					break
				}
			}
		}

	}
	fmt.Println(sum)
}
