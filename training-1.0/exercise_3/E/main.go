package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"

    "strings"
)

func main() {

    m := make(map[string]bool)
    k := make(map[string]bool)

    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan() 

        line := strings.Fields(scanner.Text())

        for _, str := range line {
            m[str] = true
        }
    

    

	scanner.Scan()
	d := scanner.Text()
    

    for _,v := range d {
        k[string(v)]=true
    }

    slc := []int{}

    // проход по d
    for v := range k {
        // если цифры из d нет в m, добавить её в слайc
        if m[v] != true {
            c,_ := strconv.Atoi(v)
            slc = append(slc, int(c))
        }
    }
    // длина слайса 
    //fmt.Println(m,k,slc)
    fmt.Printf("%d\n", len(slc))
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