package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	const maxCapacity = 1024 * 2048 // 2МБ
	buf := make([]byte, 0, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// слайс т.к. у одного слова в словаре м.б.> 1 ударения
	dictMap := make(map[string][]int)

	for i := 0; i < n; i++ {
		// если n=0 - пустой словарь или проверка на ноль.
		if !scanner.Scan() {
			break
		}
		word := scanner.Text()
		wordLow := strings.ToLower(word)
		for key, value := range word {
			if unicode.IsUpper(value) {
				// заполняем индексами ударений слова из словаря
				dictMap[wordLow] = append(dictMap[wordLow], key)
			}
		}
	}

	scanner.Scan()
	line := strings.Fields(scanner.Text()) // во входных данных не всегда ровно один пробел

	result := 0
	// разбор каждого слова из строки
	for _, oneWord := range line {
		// если слово есть в словаре - запятая,ок
		if curWord, ok := dictMap[strings.ToLower(oneWord)]; ok {

			var uppers []int
			for key, value := range oneWord {
				// индексы ударений в слове
				if unicode.IsUpper(value) {
					uppers = append(uppers, key)
				}
			}

			// больше одного ударения или вообще нет
			if len(uppers) != 1 {
				result++
				// если ударение = 1, то совпадают ли индексы
			} else {
				tempCount := 0
				for _, Stress := range curWord {
					// ударение слова из строки = хотя одному удар-ю из словаря
					if Stress == uppers[0] {
						tempCount++
						break
					}
				}
				// если индексы не совпали
				if tempCount == 0 {
					result++
				}
			}

		} else { // если слова нет в словаре
			var uppers []int
			for key, value := range oneWord {
				if unicode.IsUpper(value) {
					uppers = append(uppers, key)
				}
			}
			// если не равно 1
			if len(uppers) != 1 {
				result++
			}
		}
	}

	fmt.Println(result)

}
