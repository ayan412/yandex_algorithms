package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// мапа для [word]:count
	mapIndex := make(map[string]int)

	for scanner.Scan() {
		w := scanner.Text()
		mapIndex[w]++
	}
	maxFr := 0
	finalWord := ""

	for k,v := range mapIndex {
		// если частота слова больше максимума
		if v > maxFr {
			maxFr = v
			finalWord = k
			// если частота у слова уже максимальна, то сравниваем с предыдущим ответом
		} else if v == maxFr && k < finalWord {
			finalWord = k
		}
	}

	fmt.Println(finalWord)

}
