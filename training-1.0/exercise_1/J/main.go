package main

import (
	"fmt"
	"math"
)

const epsilon = 1e-7

func main() {
	var a, b, c, d, e, f float64
	fmt.Scan(&a, &b, &c, &d, &e, &f)

	// Применяем метод Гаусса

	// поменять местами зн-я коэфф-в, чтобы не получить деления на ноль
	if math.Abs(a) < math.Abs(c) {
		a, c = c, a
		b, d = d, b
		e, f = f, e
	}

	if math.Abs(a) > epsilon {
		m := c / a //
		c = 0      // приведение к ступенчатому виду под 1-м коэфф-м т.е С
		d -= m * b // Это соответствует (d - (c/a)b)
		f -= m * e // Это соответствует (f - (c/a)e)
	}

	/*
		Начинаем с уравнения:
			cx + dy = f
		Мы вычитаем из него первое уравнение, умноженное на m = c/a:
			cx + dy = f - (c/a)(ax + by = e)
		Раскрывая скобки, получаем:
			cx + dy - (c/a)ax - (c/a)by = f - (c/a)e
		Упрощаем:
			cx - cx + dy - (c/a)by = f - (c/a)e
			cx - cx сокращается до 0, поэтому x действительно исчезает из уравнения:
			dy - (c/a)by = f - (c/a)e
		Приводим подобные члены:
			(d - (c/a)b)y = f - (c/a)e
	*/

	// Анализируем результаты
	
	// случай ур-я 0x+0y=0
	if math.Abs(a) <= epsilon && math.Abs(c) <= epsilon {
        if math.Abs(b) <= epsilon && math.Abs(d) <= epsilon {
            if math.Abs(e) <= epsilon && math.Abs(f) <= epsilon {
                fmt.Println(5) // Любая пара чисел является решением
            } else {
                fmt.Println(0) // Нет решений
            }
        } else if math.Abs(b) > epsilon || math.Abs(d) > epsilon {
            if math.Abs(b) <= epsilon && math.Abs(e) > epsilon {
                fmt.Println(0) // Нет решений
            } else if math.Abs(d) <= epsilon && math.Abs(f) > epsilon {
                fmt.Println(0) // Нет решений
            } else if math.Abs(b*f - d*e) <= epsilon {
                y := 0.0
                if math.Abs(b) > epsilon {
                    y = e / b
                } else {
                    y = f / d
                }
                fmt.Printf("4 %.9f\n", y) // y = const, x - любое
            } else {
                fmt.Println(0) // Нет решений
            }
        } else {
            fmt.Println(0) // Нет решений
        }
    } else {
        if math.Abs(d) <= epsilon {
            if math.Abs(f) <= epsilon {
                if math.Abs(b) <= epsilon {
                    x0 := e / a
                    fmt.Printf("3 %.9f\n", x0) // x = x0, y - любое
                } else {
                    k := -a / b
                    b := e / b
                    fmt.Printf("1 %.9f %.9f\n", k, b) // y = kx + b
                }
            } else {
                fmt.Println(0) // Нет решений
            }
        } else {
            y := f / d
            x := (e - b*y) / a
            fmt.Printf("2 %.9f %.9f\n", x, y) // Единственное решение
        }
    }
}