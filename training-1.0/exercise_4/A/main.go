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

	// скан-м и получаем кол-во пар
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	dir := make(map[string]string)

	// скан-м строки со словами
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		words := strings.Split(line, " ")
		a, b := words[0], words[1]
		dir[a] = b
	}

	// скан-м слово которому найти его пару
	scanner.Scan()
	find := scanner.Text()

	for k := range dir {
		if dir[k] == find {
			fmt.Println(k)
		}
		// выводим по самому слову в кач-ве ключа
		if k == find {
			fmt.Println(dir[find])
		}
	}
}
