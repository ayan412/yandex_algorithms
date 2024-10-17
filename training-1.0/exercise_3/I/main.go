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

	scanner.Scan()

	N,_ := strconv.Atoi(scanner.Text())

	m := make(map[string]int)

	for scanner.Scan() {

		word := scanner.Text()
		// если слово нельзя сконвертировать в число значит это строка
		if _, err := strconv.Atoi(word); err != nil {
			m[word]++
		}

	}
	
	commonLang := []string{}
	for lang,count := range m {
		if count == N {
			commonLang = append(commonLang, lang)
		}
	}

	fmt.Println(len(commonLang))
	for _,v := range commonLang {
		fmt.Println(v)
	}

	fmt.Println(len(m))
	for k := range m {
		fmt.Println(k)
	}
}
