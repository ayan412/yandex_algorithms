package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Tuple struct {
	count int // сколько раз встретился элемент
	idx   int // каким по счету встретился в первый раз
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var n int
	var C, D string
	// fmt.Scan(&n, &C, &D)
	keywords := make(map[string]bool)
	// scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	slcLine := strings.Fields(scanner.Text())
	//fmt.Println(slcLine)

	n, _ = strconv.Atoi(slcLine[0])
	C, D = slcLine[1], slcLine[2]

	caseSensitive := (C == "yes")
	startNum := (D == "yes")

	for i := 0; i < n && scanner.Scan(); i++ {
		word := scanner.Text()
		// если не важен регистр
		if !caseSensitive {
			word = strings.ToLower(word)
		}
		keywords[word] = true
	}

	// Все последовательность из латинских букв, цифр и знаков подчеркивания
	regexpStr := `[a-zA-Z0-9_]*[a-zA-Z_]+[a-zA-Z0-9_]*`
	regexPattern := regexp.MustCompile(regexpStr)

	idCount := make(map[string]Tuple)

	for scanner.Scan() {
		line := scanner.Text()
		// все возможные совпадения без ограничений (-1)
		matches := regexPattern.FindAllString(line, -1)
		//fmt.Println(matches)
		for _, elem := range matches {
			// начало НЕ должно быть с цифры && начало явл-я цифрой т.е. !false = true
			// TRUE && TRUE
			if !startNum && unicode.IsDigit(rune(elem[0])) { // D == no - первой не может быть цифра
				continue // пропуск - след-ая итерация elem
			}
			if !caseSensitive { // C == no - нечувствительны т.е. if !false = TRUE
				elem = strings.ToLower(elem)
			}
			if !keywords[elem] { // не ключевое слово - нет в мапе ключевых слов
				// проверка на сущ-е в мапе - запятая, ок
				value, ok := idCount[elem]
				if !ok { // элемент не найден в мапе т.е. встречается впервые, и в мапе по ключу elem инициал-м структуру
					// fmt.Println(idCount)
					idCount[elem] = Tuple{0, len(idCount)}
					//fmt.Println(len(idCount), idCount)
				} else { // если эл-нт уже встречался
					value.count++         // обновляем ТОЛЬКО счетчик в структуре
					idCount[elem] = value // перезаписываем структуру с обнов-м счетчиком в мапу по ключу
					// fmt.Println(len(idCount), idCount)
				}
			}
		}
	}

	ans := ""
	maxCount := 0
	ansId := len(idCount)
	for key, value := range idCount {
		count := value.count
		id := value.idx
		/*
			идентификатор, который встречается чаще (count > maxCount),
			либо идентификатор, который встречается столько же раз (count == maxCount),
			но появился раньше (id < ansId)
		*/
		if count > maxCount || (count == maxCount && id < ansId) {
			ans, maxCount, ansId = key, count, id
		}
	}
	fmt.Println(ans)
}
