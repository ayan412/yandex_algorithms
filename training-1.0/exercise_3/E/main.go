package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	m := make(map[rune]bool)
	for _, ch := range scanner.Text() {
		if ch != ' ' {
			m[ch] = true
		}
	}

	scanner.Scan()
	d := scanner.Text()

	count := 0
	for _, ch := range d {
		if !m[ch] {
			m[ch] = true
			count++
		}
	}

	fmt.Println(count)
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	
	"strings"
)

func main() {

	m := make(map[string]bool)
	k := make(map[string]bool)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := strings.Fields(scanner.Text())

		for _, str := range line {
			m[str] = true
		}
	}

	var d string
	fmt.Scan(&d)

	for _,v := range d {
		k[string(v)]=true
	}

	slc := []string{}

	// проход по d
	for v := range k {
		// если цифры из d нет в m, добавить её в слайc
		if m[v] == false {
			slc = append(slc, v)
		}
	}
	// длина слайса 
	fmt.Print(len(slc))
}
*/