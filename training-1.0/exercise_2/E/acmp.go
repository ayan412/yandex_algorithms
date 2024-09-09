package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Scan(&size)
	numbers := make([]int, size)
	for i:=0;i<size;i++ {
		fmt.Scan(&numbers[i])
	}

	//numbers := []int{10, 20, 15, 10, 30, 5, 1}
	//numbers := []int{15,15,10}
	var winnerPoints = 0
	for i:=0;i<len(numbers);i++ {
		if numbers[i] > winnerPoints {
			winnerPoints = numbers[i]
		}
	}
	//fmt.Println(winnerPoints)
	var winnerCount = 0
	var maxPoints = 0
	for i:=0;i<len(numbers)-1;i++{
		//fmt.Printf("%d", numbers[i])
		if winnerCount >= 1 && numbers[i]%10 == 5 && numbers[i+1] < numbers[i] {
			if numbers[i] > maxPoints {
				maxPoints = numbers[i]
			}
		}
		if numbers[i] == winnerPoints {
			winnerCount++
		}
		
	}
	//fmt.Println(winnerCount, maxPoints)

	if maxPoints == 0 {
		fmt.Println("0")
		return
	} 

	var nBetter int = 0
	for i:=0;i<len(numbers);i++{
		if numbers[i] > maxPoints {
			nBetter++
		}
	}
	fmt.Printf("%d\n", nBetter+1)
		
}
