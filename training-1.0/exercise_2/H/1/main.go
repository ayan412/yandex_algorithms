package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//numbers := []int{30, 0, 50, 1, -100}
	//numbers := []int{3, 5, 1, 7, 9, 0, 9, -3, 10}
	// numbers := []int{-5, -30000, -12}
	// numbers := []int{1,2,3}
	//numbers := []int{9, 5, 1, 4, 7, 9, 7, 8, 2}

	var numbers []int64
	if scanner.Scan() {
	input := scanner.Text()
	strNum := strings.Fields(input)

	for _, v := range strNum {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
		fmt.Println(err)
		}
		numbers = append(numbers, num)
	}
	}

	newSort(numbers,0,int64(len(numbers)),int64(len(numbers))-1)
	// if err != nil {
	// 	// Обработка ошибки
	// 	fmt.Println("Error in 1:", err)
	// 	return
	// }
	// fmt.Println("res1", numbers)
	newSort(numbers,0,int64(len(numbers)-1),int64(len(numbers)-2))
	// if err != nil {
	// 	// Обработка ошибки
	// 	fmt.Println("Error in 2:", err)
	// 	return
	// }
	// fmt.Println("res2", numbers)
	newSort(numbers,0,int64(len(numbers)-3),2)
	// if err != nil {
	// 	// Обработка ошибки
	// 	fmt.Println("Error in 3:", err)
	// 	return
	// }
	// time.Sleep(1000 * time.Millisecond)
	// fmt.Println(numbers)

	res1 := numbers[0] * numbers[1] * numbers[len(numbers)-1]
	res2 := numbers[len(numbers)-1] * numbers[len(numbers)-2] * numbers[len(numbers)-3]

	if res2 >= res1 {
		fmt.Println(numbers[len(numbers)-1],numbers[len(numbers)-2], numbers[len(numbers)-3])
	} else {
		fmt.Println(numbers[0], numbers[1], numbers[len(numbers)-1])
	}
}

func newSort(seq []int64,left,right, k int64) {
	left = 0
	right = int64(len(seq))-1

	//fmt.Println(left, right)

	for left < right {
		x := seq[(left+right)/2]
		eqxfirst := left
		gtxfirst := left
		//fmt.Println(eqxfirst, gtxfirst, x, seq)

		for i := left; i <= right; i++ {
			now := seq[i]
			//fmt.Println(now)
			if now == x {
				seq[i] = seq[gtxfirst]
				seq[gtxfirst] = now
				gtxfirst++
				//fmt.Printf("%d %d=%d %d\n", i, seq[i], x, seq)
			} else if now < x {
				seq[i] = seq[gtxfirst]
				seq[gtxfirst] = seq[eqxfirst]
				seq[eqxfirst] = now
				gtxfirst++
				eqxfirst++
				//fmt.Printf("%d %d<%d %d\n", i, seq[i], x, seq)
			}
		}

		if k < eqxfirst {
			right = eqxfirst - 1
			//fmt.Println(right)
		} else if k >= gtxfirst {
			left = gtxfirst
			//fmt.Println(left)
		} else {
			return
		}
	}
}


// func newSort(seq []int64, left, right, k int64) error {
// 		if len(seq) == 0 {
// 			return fmt.Errorf("empty slice")
// 		}
	
// 		left = 0
// 		right = int64(len(seq)) - 1

// 		if k < 0 || k >= int64(len(seq)) {
// 			return fmt.Errorf("k is out of range")
// 		}

// 		for left < right {
// 			if left < 0 || right >= int64(len(seq)) {
// 				return fmt.Errorf("index out of range: left=%d, right=%d", left, right)
// 			}

// 			mid := (left + right) / 2
// 			x := seq[mid]
// 			eqxfirst := left
// 			gtxfirst := left

// 			for i := left; i <= right; i++ {
// 				if i < 0 || i >= int64(len(seq)) {
// 					return fmt.Errorf("index out of range: i=%d", i)
// 				}

// 				now := seq[i]
// 				if now == x {
// 					if gtxfirst < 0 || gtxfirst >= int64(len(seq)) {
// 						return fmt.Errorf("index out of range: gtxfirst=%d", gtxfirst)
// 					}
// 					seq[i], seq[gtxfirst] = seq[gtxfirst], now
// 					gtxfirst++
// 				} else if now < x {
// 					if gtxfirst < 0 || gtxfirst >= int64(len(seq)) || eqxfirst < 0 || eqxfirst >= int64(len(seq)) {
// 						return fmt.Errorf("index out of range: gtxfirst=%d, eqxfirst=%d", gtxfirst, eqxfirst)
// 					}
// 					seq[i], seq[gtxfirst], seq[eqxfirst] = seq[gtxfirst], seq[eqxfirst], now
// 					gtxfirst++
// 					eqxfirst++
// 				}
// 			}

// 			if k < eqxfirst {
// 				right = eqxfirst - 1
// 			} else if k >= gtxfirst {
// 				left = gtxfirst
// 			} else {
// 				return nil
// 			}
// 		}

// 		return nil
// 	}