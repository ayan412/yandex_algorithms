package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(scanner *bufio.Scanner) map[int]bool {

	m := make(map[int]bool)

	if scanner.Scan() {
		line := scanner.Text()
		num := strings.Fields(line) 

		for _,v := range num {
			n, _ := strconv.Atoi(v)
			m[n] = true
		}
	}
	return m
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	m1 := readLines(scanner)
	m2 := readLines(scanner)

	//fmt.Println(m1, m2)

	intersection := []int{}
	for n := range m1 {
		if m2[n] {
			intersection = append(intersection, n)
		}
	}

	sort.Ints(intersection)

	for _, num := range intersection {
        fmt.Printf("%d ", num)
    }
    fmt.Println()

}