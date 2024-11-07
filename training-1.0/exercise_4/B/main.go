package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// мапа для [слово}:индекс по порядку
	mapIndex := make(map[string][]int)

	indexCount := 0
	for scanner.Scan() {
		w := scanner.Text()
		mapIndex[w] = append(mapIndex[w], indexCount)
		indexCount++
	}

	res := make([]int, indexCount)

	for _, slc := range mapIndex {
		for index, value := range slc {
			res[value] = index
		}
	}

	for _,v := range res {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
