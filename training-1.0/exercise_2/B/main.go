package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func determineSequenceType(numbers []int) string {
	if len(numbers) <= 1 {
		return "CONSTANT"
	}

	isConstant := true
	isAscending := true
	isWeaklyAscending := true
	isDescending := true
	isWeaklyDescending := true

	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[i-1] {
			isConstant = false
		}
		if numbers[i] <= numbers[i-1] {
			isAscending = false
		}
		if numbers[i] < numbers[i-1] {
			isWeaklyAscending = false
		}
		if numbers[i] >= numbers[i-1] {
			isDescending = false
		}
		if numbers[i] > numbers[i-1] {
			isWeaklyDescending = false
		}
	}

	if isConstant {
		return "CONSTANT"
	} else if isAscending {
		return "ASCENDING"
	} else if isWeaklyAscending {
		return "WEAKLY ASCENDING"
	} else if isDescending {
		return "DESCENDING"
	} else if isWeaklyDescending {
		return "WEAKLY DESCENDING"
	}

	return "RANDOM"
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var numbers []int

	for scanner.Scan() {
		text := scanner.Text()
		if text == "-2000000000" {
			break
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Ошибка при преобразовании строки в число:", err)
			continue
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		return
	}
	//fmt.Println("Считанные числа:", numbers)

	sequenceType := determineSequenceType(numbers)
	fmt.Println(sequenceType)
}
