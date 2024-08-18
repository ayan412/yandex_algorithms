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

	var numbers []int

	if scanner.Scan() {
		input := scanner.Text()
		strNumbers := strings.Fields(input)

		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	for i := 1; i < len(numbers); i++ {
		if numbers[i] <= numbers[i-1] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
