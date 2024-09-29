package main

import (
	"fmt"
	"strings"
)

func main() {

	var lines, columns, mines int

	fmt.Scanf("%d %d %d", &lines, &columns, &mines)

	// интерфейс для хранения и чисел и строк
	// двумерный слайс создан сначала по зн-ю строк
	field := make([][]interface{}, lines)
	for i := range field {
		// создание строки по зн-ю колонки
		field[i] = make([]interface{}, columns)
		// заполнение всех клеток нулями
		for j := range field[i] {
			field[i][j] = 0 
		}
	}

	// Ф-я для обновления соседних клеток после установки мин
	// аноним-я функция - замыкание closure
	updateField := func(line, column int) {

		coordinates := [][2]int { // двумерный слайс неогранич-й по кол-ву строк, только по кол-ву сивм-в в строке - их 2
			// смещение по 8-ми направлениям относ-но одной точки
			{line - 1, column - 1}, {line - 1, column}, {line - 1, column + 1},
			{line, column - 1}, {line, column + 1}, 
			{line + 1, column - 1}, {line + 1, column}, {line + 1, column + 1},
		}


		// проверка каждой соседней клетки
		for _, coordinate := range coordinates {
			i,j := coordinate[0], coordinate[1]
			// пропуск клеток больше значений поля
			if i < 0 || j < 0 || i > lines - 1 || j > columns - 1 {
				continue
			}
			// пропуск клеток где мина
			if field[i][j] == "*" {
				continue
			}
			// увел-е зн-е на + 1
			if value, ok := field[i][j].(int);ok {
				field[i][j] = value + 1
			}
			// field[i][j] = field[i][j].(int) + 1 // утверждение типа к int	
		}
	}

	// Вывод мин на поле
	for i := 0; i < mines; i++ {
		var line, column int 

		fmt.Scanf("%d %d", &line, &column)
		field[line-1][column-1] = "*"
		updateField(line-1, column-1)
	}

	// Вывод итогового рез-та
	for _, row := range field {
		var strRow []string
		for _,cell := range row {
			strRow = append(strRow, fmt.Sprintf("%v", cell))
		}
		// вывод символов разделенных пробелом
		fmt.Println(strings.Join(strRow, " "))
	}

}