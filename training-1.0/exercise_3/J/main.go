package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// коор-ты перес-я с осью Y и X
type CoordArea struct {
	bMax, bMin, bMax_, bMin_ int
}

// координаты
type Coord struct {
	x, y int
}

// Расшир-е прямоугол-а сначала с учетом погрешности d, а затем с учетом t
func extend(initial CoordArea, extension int) CoordArea {
	return CoordArea{
		bMax:  initial.bMax + extension,
		bMin:  initial.bMin - extension,
		bMax_: initial.bMax_ + extension,
		bMin_: initial.bMin_ - extension,
	}
}

// пересеничение 2-х прямоуг-в дают примерную обл-ть нужных для ответа координат с которыми нужно работать дальше проверяя на четность для уже точного ответа
func intersection(lhs, rhs CoordArea) CoordArea {
	return CoordArea{
		bMax:  min(lhs.bMax, rhs.bMax),
		bMin:  max(lhs.bMin, rhs.bMin),
		bMax_: min(lhs.bMax_, rhs.bMax_),
		bMin_: max(lhs.bMin_, rhs.bMin_),
	}
}

/*
проверка на четность исходя из СЛАУ:
b = y - x
b_ = y + x

сложим между собой а потом выразив y, подставим в любое урав-е:
b + b_ = y + x + y - x = 2y ---> y = (b + b_)/2

подставим найденный у во 2-е ур-е:
x = b_ - y = b_ - (b+b_)*1/2 = (b_ - b)/2

тоже самое будет если подставить в 1-е ур-е x = y - b
*/
func getCoords(area CoordArea) []Coord {
	var result []Coord
	// ходим циклом в пределах этих 4-х коор-т для перебора нужных точек в их пределах
	for b := area.bMin; b <= area.bMax; b++ {
		for b_ := area.bMin_; b_ <= area.bMax_; b_++ {
			// проверка на четность: если коорд-а четна в коор-х b и b_ то она даст целую коор-у и в x и y
			// Если (b + b') чётное, то (b' - b) тоже будет чётным
			if (b+b_)%2 == 0 {
				// своего рода возврат в исход-е коорд-ы x и y
				y := (b + b_) / 2
				x := (b_ - b) / 2

				result = append(result, Coord{x, y})
			}
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// сканирование отдельных чисел сразу с преобраз-м
	readInt := func() int {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		return num
	}

	// иниц-я и присвоение
	t, d, n := readInt(), readInt(), readInt()
	possibleLocations := CoordArea{0, 0, 0, 0}

	//разобрать как считываются коор-ы possibleLocations possibleGPSLocations possibleMoveLocations

	for i := 0; i < n; i++ {
		x, y := readInt(), readInt()
		// коор-ы с учетом d
		possibleGPSLocations := extend(CoordArea{y - x, y - x, y + x, y + x}, d)
		// коор-ы с учетом t
		possibleMoveLocations := extend(possibleLocations, t)
		// коор-ы пересечения
		possibleLocations = intersection(possibleMoveLocations, possibleGPSLocations)
	}
	// проверка на четность
	answer := getCoords(possibleLocations)
	fmt.Println(len(answer))
	for _, coord := range answer {
		fmt.Println(coord.x, coord.y)
	}
}

/*
   	 * решение с перебором всех точек при поиске пересечения не проходит по времени
     * нужно решение, которое определяет пересечение возможных локаций за О(1)
     * чтобы описать прямоугольник в декартовой системе координат со сторонами, параллельными осям координат,
     * достаточно 4 чисел: xMin, xMax, yMin, yMax
     * искать пересечение двух прямоугольников, заданных таким образом за О(1) достаточно легко
     *
     * в нашем случае возможные локации - это такой же прямоугольник, но повёрнутый относительно осей координат
     * множество точек, удалённых на манхэттенское расстояние от некоторой начальной точки, на координатной оси
     * образуют квадрат, повёрнутый на угол 45 градусов относительно координатных осей
     * пересечение таких квадратов в общем случае образует прямоугольник
     * для описания такого прямоугольника нужно знать характеристики четырёх его диагоналей
     * уравнение прямой y = k*x + b при наклоне прямоугольника 45 градусов вырождается в y = x + b т.к. tg45 tg135 равен 1 и -1 соотв-но
	 * и где b описывает точку пересечения с осью y и х.
     * прямоугольник ограничен четырьмя диагоналями:
     * y = x + bMax (левая верхняя)
     * y = x + bMin (правая нижняя)
     * y = -x + bMax_ (правая верхняя)
     * y = -x + bMin_ (левая нижняя)
     *
     * в итоге для описания манхэттенских прямоугольников достаточно знать 4 числа (свободные члены в уравнениях
     * ограничивающих прямоугольник прямых) и не перепутать порядок хранения этих чисел
     *
*/

/*
 *  левая верхняя       правая верхняя
 *         . . . . * . . . .
 *         . . . * . * . . .
 *         . . * . . . * . .
 *         . * . . . . . * .
 *         . . * . . . * . .
 *         . . . * . * . . .
 *         . . . . * . . . .
 *  левая нижняя        правая нижняя
 */
