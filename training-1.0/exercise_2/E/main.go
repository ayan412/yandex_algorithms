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

	//line := "10 20 15 10 30 5 1"
	//line := "15 15 10"
	//line := "10 15 20"
	//line := "20 15 15 15 15 1"
	//line := "30 15 30 5 1"
	//line := "30 35 25 1"
	//line := "10 20 15 10 30 5 1"

	//strNumbers := strings.Split(line, " ")

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
			fiveMap[k] = num
		}
		indexMap[k] = num
		numbers = append(numbers, num)
	}

	// fmt.Println(strNumbers, indexMap, fiveMap, max, min, numbers)

	//var noplace int = -1

	// var bestPlace = len(numbers) + 1
	foundPlace := false
	place := 1

	//uniqueNumbers := removeDuplicates(numbers)

	for i, num := range fiveMap {
		for k, v := range indexMap {
			if v == max && i > k && indexMap[i+1] == min {
				if num == min {
					foundPlace = false
					break
				}
				if num == max {
					foundPlace = true
					break
				} else {
					for _, value := range indexMap {
						if value > num {
							foundPlace = true
							place++
						}
					}

				}
				break
			}
		}
	}

	if foundPlace == false {
		fmt.Println(0)
	} else {
		fmt.Println(place)
	}
}

// func removeDuplicates(slice []int) []int {
//     keys := make(map[int]struct{})
//     i := 0
//     for _, entry := range slice {
//         if _, exists := keys[entry]; !exists {
//             keys[entry] = struct{}{}
//             slice[i] = entry
//             i++
//         }
//     }
//     return slice[:i]
// }
