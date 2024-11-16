package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// мапа для хранения баланса по фамилии
	// в Г для несущест-го ключа значение по умолч-ю т.о. вып-ся ус-е о "ему заводится счет с нулевым балансом."
	dataMap := make(map[string]int)
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")
		// тек-я операция
		operation := line[0]

		switch operation {

		case "DEPOSIT":
			{
				// фамилия
				surname := line[1]
				// сумма
				price, _ := strconv.Atoi(line[2])
				// заполнение мапы
				dataMap[surname] += price
			}

		case "INCOME":
			{
				// процент
				percent, _ := strconv.Atoi(line[1])
				// общая сумма с учетом процента
				for k, v := range dataMap {
					// только для полож-го баланса
					if v > 0 {
						dataMap[k] = (v * percent / 100) + v
					}
				}
				//fmt.Println(operation, dataMap)
			}

		case "BALANCE":
			{
				surname := line[1]
				// если баланс по фамилии сущест-т, то вывести иначе ошибка
				if value, ok := dataMap[surname]; ok {
					fmt.Println(value)
				} else {
					fmt.Println("ERROR")
				}

			}

		case "TRANSFER":
			{
				// фамилии
				surname1, surname2 := line[1], line[2]
				price, _ := strconv.Atoi(line[3])
				// новые значения мап согласно ключам
				dataMap[surname1] -= price
				dataMap[surname2] += price
			}

		case "WITHDRAW":
			{
				// фамилия
				surname := line[1]
				price, _ := strconv.Atoi(line[2])
				// новое зн-ие
				dataMap[surname] = dataMap[surname] - price
				//fmt.Println(dataMap[surname])
			}
		}
	}
}
