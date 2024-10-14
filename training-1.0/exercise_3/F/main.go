package main

import (
	"bufio"
	"fmt"
	"os"

)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 10000000), 1000000)
	
	scanner.Scan()
	f := scanner.Text()
	scanner.Scan()
	s := scanner.Text()

	mf := make(map[string]int) 
	
	for i:=0;i<len(f)-1;i++{
		fa := string(f[i])+string(f[i+1])
		mf[fa]++
	}
	
	ms := make(map[string]int)

	for i:=0;i<len(s)-1;i++{
		sa := string(s[i])+string(s[i+1])
		ms[sa]++
	}

	count := 0

	for v := range ms {
		if mf[v] != 0 {
			count += mf[v]
		} 
	}
	fmt.Println(count)
}
