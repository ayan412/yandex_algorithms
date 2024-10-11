package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	na := make(map[int]bool)
	nb := make(map[int]bool)

	//slcn, slcm := []int{}, []int{}

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			nn, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("type converting problem: ", err)
			}
			na[nn] = true
		}
	}

	for i := 0; i < m; i++ {
		if scanner.Scan() {
			mm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("type converting problem: ", err)
			}
			nb[mm] = true
		}
	}

	intersection := []int{}
	nointersectionA := []int{}
	nointersectionB := []int{}
	for n := range na {
		if nb[n] {
			intersection = append(intersection, n)
		} else {
			nointersectionA = append(nointersectionA, n)
		}
	}

	for n := range nb {
		if na[n] == false {
			nointersectionB = append(nointersectionB, n)
		}
	}

	sort.Ints(intersection)
	sort.Ints(nointersectionA)
	sort.Ints(nointersectionB)

	//fmt.Println("-------")

	fmt.Println(len(intersection))
	for _, num := range intersection {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	fmt.Println(len(nointersectionA))
	for _, num := range nointersectionA {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	fmt.Println(len(nointersectionB))
	for _, num := range nointersectionB {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
