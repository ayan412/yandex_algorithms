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

	var lines [2]string
	for i := range lines {
		if scanner.Scan() {
			lines[i] = scanner.Text()
		} else {
			break
		}
	}

	numbers := make([]int, 0)
	for _,numStr := range strings.Split(lines[1], " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка при преобразовании числа:", err)
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Println(lines)

}