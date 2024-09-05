package main

import (
	"fmt"
)

func GetPrizePlace(resultVas int, nums []int) int {
	if resultVas == 0 {
		return 0
	}
	// Search for the maximum prize place that can be taken with the found number
	maxProbPlace := 0
	for _, num := range nums {
		if num > resultVas {
			maxProbPlace++
		}
	}
	return maxProbPlace + 1
}

func DefineResultVasily(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	findEl := false
	indNeedEl := 0
	indMax := 0
	if nums[0] <= nums[len(nums)-1] {
		indMax = len(nums) - 1
	}
	// Search for the desired number that matches the search criteria
	for i := 1; i < len(nums)-1; i++ {
		currEl := nums[i]
		if currEl > nums[indMax] {
			findEl = false
			indMax = i
		} else if currEl%5 == 0 && currEl%2 != 0 && currEl > nums[i+1] {
			if findEl {
				if currEl > nums[indNeedEl] {
					indNeedEl = i
				}
			} else {
				findEl = true
				indNeedEl = i
			}
		}
	}
	if !findEl || (findEl && indNeedEl <= indMax) {
		return 0
	}
	return nums[indNeedEl]
}

func mainn() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}
	fmt.Println(GetPrizePlace(DefineResultVasily(numbers), numbers))
}
// 1000 995 990 985 980 975 970 965 960 955 950 945 940 935 930 925 920 915 910 905 - c этими данными неверный ответ 2.
