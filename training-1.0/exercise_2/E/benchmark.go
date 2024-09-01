package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"slices"
)

func bench() {
	sizes := []int{10000, 100000, 1000000, 100000000}

	for _, size := range sizes {
		fmt.Printf("Размер слайса: %d\n", size)
		
		slice := generateRandomSlice(size)
		
		loopTime := benchmarkLoop(slice)
		fmt.Printf("Время цикла: %v\n", loopTime)
		
		sortTime := benchmarkSort(slice)
		fmt.Printf("Время сортировки: %v\n", sortTime)
		
		fmt.Println()
	}
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}

func benchmarkLoop(slice []int) time.Duration {
	start := time.Now()

	max, min := -1, 100001
	for _,v := range slice {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max, min)
	
	return time.Since(start)
}

func benchmarkSort(slice []int) time.Duration {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	
	start := time.Now()
	
	sort.Ints(sliceCopy)

	max := slices.Max(slice)
	min := slices.Min(slice)

	fmt.Println(max, min)
	
	return time.Since(start)
}