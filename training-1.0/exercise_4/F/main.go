package main

import (
	"bufio"
	"fmt"

	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// иниц-я внешней (осн-й) мапы, где ключ - это покупатель, а знач-я другая НЕСОЗДАННАЯ мапа,
	// поэтому когда впервые добав-я новый покупатель, внутренняя мапа для его товаров ещё не сущ-вует
	dataMap := make(map[string]map[string]int)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		buyer := line[0]
		goods := line[1]
		amount, _ := strconv.Atoi(line[2])
		fmt.Println()
		//fmt.Printf("-%v %v %v", buyer, goods,amount)
		// если по ключу (покупателю) нет данных, то создать для него внутр-ю мапу
		if _, ok := dataMap[buyer]; !ok {
			dataMap[buyer] = make(map[string]int)
		}
		// запись нового товара и суммар-го кол-ва для конкретного ключа
		dataMap[buyer][goods] += amount
	}
	//fmt.Println()
	//fmt.Println(len(dataMap))

	// сортировка сначала покупателя для алфавитного вывода
	keys := make([]string, 0)
	for key := range dataMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		// Вывод покупателей
		fmt.Printf("%s:\n", key)

		// сортировка товаров
		var goods []string

		// item = ключ т.е. товар
		for item := range dataMap[key] { // по соответ-у связанному ключу
			goods = append(goods, item)
		}
		sort.Strings(goods)

		// вывод товара и кол-ва по отсорт-му списку
		for _, item := range goods {
			fmt.Printf("%s %d\n", item, dataMap[key][item])
		}
	}
}
