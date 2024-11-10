package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024) // 64KB
	scanner.Buffer(buf, 1024*1024)   // буфер 1MB

	// скан-м кол-во клавиш
	scanner.Scan()
	n,_ := strconv.Atoi(scanner.Text())

	// скан-м допус-е кол-во нажатий
	scanner.Scan()
	slcCint := make([]int, n)
	for i, v := range strings.Fields(scanner.Text()) {
    slcCint[i], _ = strconv.Atoi(v)
	}

	// скан-м общее кол-во нажатий
	scanner.Scan()
	//k := scanner.Text()

	// скан-м послед-ть нажатых клавиш
	scanner.Scan()
    pressCount := make(map[string]int, n)
    for _, key := range strings.Split(scanner.Text(), " ") {
        pressCount[key]++
    }

    // Проверяем превышение допустимого количества нажатий
    for i := 0; i < n; i++ {
        key := strconv.Itoa(i + 1)
        if pressCount[key] > slcCint[i] {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}