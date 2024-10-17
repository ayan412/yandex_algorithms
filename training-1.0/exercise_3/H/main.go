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

	scanner.Scan()

	m := make(map[int]struct{})

	for scanner.Scan() {
		str := strings.Fields(scanner.Text())
		if len(str) == 0 {
			break
		}
		a, _ := strconv.Atoi(str[0])
		m[a] = struct{}{}

	}
	fmt.Println(len(m))
}
