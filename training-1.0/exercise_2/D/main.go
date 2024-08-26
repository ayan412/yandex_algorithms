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

	var numbers, res []int
	

	if scanner.Scan() {
		text := scanner.Text()
		strNumbers := strings.Fields(text)

		for _, strNum := range strNumbers {
			num,err := strconv.Atoi(strNum)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	for i:=1;i<len(numbers)-1;i++ {
		if numbers[i]>numbers[i+1] && numbers[i]>numbers[i-1] {
			res = append(res, numbers[i])
		}
	}
	fmt.Println(len(res))
}
