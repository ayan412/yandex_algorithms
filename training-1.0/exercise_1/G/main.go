package main

import (
	"fmt"
)

func main() {
	var N, K, M int

	// Ввод исходных данных
	fmt.Scan(&N, &K, &M)

	if M > K {
		// Если масса детали больше массы заготовки, то невозможно выточить ни одной детали
		fmt.Println(0)
		return
	}

	totalDetails := 0

	for N >= K {
		// Сколько заготовок можно сделать
		blanks := N / K
		// Остаток сплава после изготовления заготовок
		remainingAlloy := N % K

		// Сколько деталей можно выточить из всех заготовок
		detailsFromBlanks := (K / M) * blanks
		totalDetails += detailsFromBlanks

		// Остаток от заготовок после вытачивания деталей
		remainingFromBlanks := blanks * (K % M)
		// Новый сплав: остаток сплава + остаток от заготовок
		N = remainingAlloy + remainingFromBlanks
	}

	fmt.Println(totalDetails)
}

