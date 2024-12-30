package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Считываем N
	scanner.Scan()
	var N int
	fmt.Sscan(scanner.Text(), &N)

	// Слайс массивов для хранения коор-т
	// points[i] - i это номер строки
	// points[i][0] - x
	// points[i][1] - y
	points := make([][2]int, N)

	// Сканируем строки чисел в слайс
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &points[i][0]) // x
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &points[i][1]) // y
	}

	// Создаем два префиксных массива:
	// leftToRight[i] будет содержать сумму всех подъемов от начала до точки i
	// при движении слева направо
	leftToRight := make([]int, N) // пока все нули

	// rightToLeft[i] будет содержать сумму всех подъемов от конца до точки i
	// при движении справа налево
	rightToLeft := make([]int, N) // пока все нули

	// Заполняем префиксные суммы для движения слева направо
	for i := 1; i < N; i++ {
		// Исполь-м предыдущую сумму
		leftToRight[i] = leftToRight[i-1]
		// Если текущая точка выше предыдущей, прибавляем к уже имею-ся сумме, а если не выше, то в массив повторно пишется знач-е той же суммы
		if points[i][1] > points[i-1][1] {
			leftToRight[i] += (points[i][1] - points[i-1][1]) // т.е. складыв-ся сама разница этих высот
		}
	}

	// Заполняем префиксные суммы для движения справа налево
	for i := N - 2; i >= 0; i-- {
		// тоже самое но наоборот - также исполь-м правый для тек-го, но уже с справа налево
		rightToLeft[i] = rightToLeft[i+1]
		// Если текущая точка выше следующей(справа от неё), добавляем подъем
		if points[i][1] > points[i+1][1] {
			rightToLeft[i] += points[i][1] - points[i+1][1]
		}
	}

	// Считываем M
	scanner.Scan()
	var M int
	fmt.Sscan(scanner.Text(), &M)

	// Обрабатываем, сканируя каждый маршрут
	for i := 0; i < M; i++ {
		scanner.Scan()
		var start int
		fmt.Sscan(scanner.Text(), &start)
		scanner.Scan()
		var end int
		fmt.Sscan(scanner.Text(), &end)

		// 0-based индексация чтобы попасть в массиве в нужные числа
		start--
		end--

		var result int
		if start <= end {
			// Если идем слева направо:
			// Берем разницу префиксных сумм в конечной и начальной точках
			result = leftToRight[end] - leftToRight[start]
		} else {
			// Если идем справа налево:
			// Берем разницу префиксных сумм в начальной и конечной точках
			result = rightToLeft[end] - rightToLeft[start]
		}

		fmt.Println(result)
	}
}
