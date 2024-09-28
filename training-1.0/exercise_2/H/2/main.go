package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var min1, min2 int64 = math.MaxInt64, math.MaxInt64
	var max1, max2, max3 int64 = math.MinInt64, math.MinInt64, math.MinInt64

	

	for scanner.Scan() {

		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}

		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
	}

	if max1*max2*max3 > min1*min2*max1 {
		fmt.Println(max1, max2, max3)
	} else {
		fmt.Println(min1,min2,max1)
	}
	
}
