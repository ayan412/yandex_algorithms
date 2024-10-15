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
	scanner.Buffer(make([]byte, 10000),10000)

	scanner.Scan()
	N,_ := strconv.Atoi(scanner.Text())
	
	
	count := make(map[int]bool)

	for scanner.Scan() {
		
		str := strings.Fields(strings.TrimSpace(scanner.Text()))
		if len(str) == 0 {
			break
		}
		a,_ := strconv.Atoi(str[0])
		b,_ := strconv.Atoi(str[1])
		// count[a] != true - на случай если есть >1 одинаковой пары a и b, истина только 1 пара
		if a >= 0 && b >= 0 && a+b==N-1 && count[a] != true {
			count[a]=true
		}
	}

	fmt.Println(len(count))

}
