package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	na := make(map[int]bool)
	nb := make(map[int]bool)

	//slcn, slcm := []int{}, []int{}

	for i:=0;i<n;i++{
		if scanner.Scan() {
			nn, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("type converting problem: ", err)
			}
			na[nn] = true
		}
	}

	for i:=0;i<m;i++{
		if scanner.Scan() {
			mm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("type converting problem: ", err)
			}
			nb[mm] = true
		}
	}

	intersection := []int{}
	for n := range na {
		if nb[n] {
			intersection = append(intersection, n)
		}
	}
	sort.Ints(intersection)

	for _, num := range intersection {
        fmt.Printf("%d ", num)
    }

	nointersectionA := []int{}
	//nointersectionB := []int{}


	// for _, val := range intersection {
	// 	if na[val]
	// }
	
	for n := range na {
		if intersection
	}


	// na[n] = true
	// nb[m] = true
	fmt.Println(na,nb, nointersectionA)
}