package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	//"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var min1, min2 int64 = math.MaxInt64, math.MaxInt64
	var max1, max2 int64 = math.MinInt64, math.MinInt64

	//var numbers []int64

	for scanner.Scan() {
	
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}

		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}

	}
	
	if min1*min2 > max1*max2 {
		fmt.Printf("%d %d\n", min1, min2)
	} else {
		fmt.Printf("%d %d\n", max2, max1)
	}
}
