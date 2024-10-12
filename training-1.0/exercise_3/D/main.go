package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 10000000), 1000000)

	m := make(map[string]bool)

	for scanner.Scan() {
		str := strings.Fields(scanner.Text())

		for _,v := range str {
			m[v] = true
		}
	}

	fmt.Println(len(m))
}

/*

package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
	

	m := make(map[string]bool)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		m[scanner.Text()] = true
	}

	fmt.Println(len(m))
}

*/
