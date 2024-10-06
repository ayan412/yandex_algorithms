package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strconv"
)

func main() {

	// чтение по числам
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// мапы как реализация структуры данных "множество" - т.е. хранит только уник-е зн-я
	m := make(map[int]struct{})

	for scanner.Scan() {

		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		// 
		m[num] = struct{}{}		
		
	}
	fmt.Println(len(m))
	//fmt.Println(m)
}
