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
		strNumbers := strings.Fields(input)

		fiveMap := make(map[int]int)
		indexMap := make(map[int]int)

		numbers := make([]int, 0)
		max, min := -1, 10001

		for k, v := range strNumbers {
			num, _ := strconv.Atoi(v)
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
			if num%5 == 0 && num%10 != 0 {
				fiveMap[k] = num
			}
			indexMap[k] = num
			numbers = append(numbers, num)
		}

		foundPlace := false
		place := 1

		if len(numbers) < 3 {
			fmt.Println("0")
			return
		}

		for i, num := range fiveMap {
			if foundPlace {
				break // Exit the outer loop if we've found a place
			}
			for k, v := range indexMap {
				if v == max && i > k && i+1 < len(indexMap) && indexMap[i+1] < num {
					if num == min {
						break
					}
					if num == max {
						foundPlace = true
						break
					} else {
						place = 1 // Reset place to 1
						for _, value := range indexMap {
							if value > num {
								place++
							}
						}
						foundPlace = true
						break
					}
				}
			}
		}

		if foundPlace {
			fmt.Println(place)
		} else {
			fmt.Println(0)
		}
	}