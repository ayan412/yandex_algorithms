package main

import (
	"bufio"
	"fmt"

	"os"
	"strconv"
	"strings"
)

func mainf() {
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
	//line := "1 1 5 5"
	//line := "10 5 1 5 1 5"
	//line := "1000 995 990 985 980 975 970 965 960 955 950 945 940 935 930 925 920 915 910 905"
	//line := "15 25 15 25 15 25 15 25 15 25"
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
		if num%10 == 5 && num%2 != 0 {
			fiveMap[k] = num
		}
		indexMap[k] = num
		numbers = append(numbers, num)
	}
	fmt.Println(getPlace(indexMap, defineRes(max, fiveMap, indexMap, numbers), min))
}

func getPlace(indexMap map[int]int, nums []int, min int) int {
	place := 1

	if len(nums) == 0  {
		return 0
	}

	mx := nums[0]
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}

	if mx == min {
		return 0
	}

	for _, value := range indexMap {
		if value > mx {
			place++
		}
	}

	if len(nums) == 0 || place == 0 {
		return 0
	}
	return place
}

func defineRes(max int, fiveMap, indexMap map[int]int, numbers []int) (nums []int) {
	var possibleNums []int
	for i, num := range fiveMap {
		for k, v := range indexMap {

			if v == max && i > k && indexMap[i+1] != 0 && num > indexMap[i+1]&& max != indexMap[len(numbers)-1] {
				possibleNums = append(possibleNums, num)
			}
		}
	}
	return possibleNums
}

