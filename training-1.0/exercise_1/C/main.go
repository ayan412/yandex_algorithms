package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var phoneNumbers []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		phoneNumbers = append(phoneNumbers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var processedNumbers []string

	for _, phone := range phoneNumbers {
		// Filter out non-digit characters
		digits := strings.Map(func(r rune) rune {
			if unicode.IsDigit(r) {
				return r
			}
			return -1
		}, phone)

		// Adjust phone numbers longer than 10 digits
		if len(digits) > 10 {
			digits = digits[1:]
		}

		// Add area code if necessary
		if len(digits) == 7 {
			digits = "495" + digits
		}

		processedNumbers = append(processedNumbers, digits)
	}

	if len(processedNumbers) == 0 {
		fmt.Println("No phone numbers to process.")
		return
	}

	firstNumber := processedNumbers[0]

	for _, number := range processedNumbers[1:] {
		if number == firstNumber {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
