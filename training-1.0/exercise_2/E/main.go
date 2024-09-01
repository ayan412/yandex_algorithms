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

	scanner.Scan()
	scanner.Scan()

	input := scanner.Text()

	//line := []string{"10 20 15 10 30 5 1"}

	strNumbers := strings.Fields(input)

	fiveMap := make(map[int]int)
	indexMap := make(map[int]int)

	numbers := make([]int, 0)
	max, min := -1, 100001

	for k, v := range strNumbers {
		num, _ := strconv.Atoi(v)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		if num%10 == 5 {
			fiveMap[num] = k
		}
		indexMap[k] = num
		numbers = append(numbers, num)
	}

	fmt.Println(strNumbers, indexMap, fiveMap, max, min, numbers)

	//var noplace int = -1

	var bestPlace = len(numbers) + 1
	foundPlace  := false

	for num, i := range fiveMap {
		for k, v := range indexMap {
			place := 1
			if v == max && i > k && indexMap[i+1] == min {
				
				for _, value := range indexMap {
					if value > num {
						place++
					}
				}
			}
			if place < bestPlace {
				bestPlace = place
				foundPlace = true
			}
		}
	}
	if foundPlace == false {
		fmt.Println(0)
	} else {
		fmt.Println(bestPlace)
	}
}
