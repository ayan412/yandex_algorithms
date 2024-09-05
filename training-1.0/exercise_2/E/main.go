package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 1000 995 990 985 980 975 970 965 960 955 950 945 940 935 930 925 920 915 910 905 - c этими данными код дает 0, но код прошедший все тесты выдает 2.
// код не проходит 9-й тест, видимо где-то ошибка
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scanner.Scan()

	input := scanner.Text()

	//strNumbers := strings.Split(line, " ")

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
		if num%10 == 5 && num%10 != 0 {
			fiveMap[k] = num
		}
		indexMap[k] = num
		numbers = append(numbers, num)
	}

	// fmt.Println(strNumbers, indexMap, fiveMap, max, min, numbers)

	//var noplace int = -1

	// var bestPlace = len(numbers) + 1
	foundPlace := false
	place := 0

	//uniqueNumbers := removeDuplicates(numbers)

	if len(numbers) < 3 {
		fmt.Println("0")
		return
	}

	for i, num := range fiveMap {
		for k, v := range indexMap {
			if v == max && i > k && indexMap[i+1] == min && num>indexMap[i+1]{
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
							place++
						}
						foundPlace = true
					}
				}
				break
			}
		}
	}

	if foundPlace == false {
		fmt.Println(0)
	} else {
		fmt.Println(place+1)
	}
}

