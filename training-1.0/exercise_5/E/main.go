// https://contest.yandex.ru/contest/27794/problems/E/
// https://coderun.yandex.ru/problem/beauty-above-all

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Считываем N
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// K
	scanner.Scan()
	K, _ := strconv.Atoi(scanner.Text())

	trees := make([]int, 0, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		trees = append(trees, num)
	}

	// fmt.Println(N, K, trees)

	// K+1 - упростить работу с индексами, uniqueTrees[0] не используется
	// например, для чисел 1 2 3 2 массив uniqueTrees представ-т собой [0 1 2 1] т.е. индекс = кол-во ПОВТОР-Й числа
	uniqueTrees := make([]int, K+1)

	left, right := 0, 1
	minLeft, minRight, minLen := 0, N-1, N // левая, правая граница и минимал длина отрезка

	// Добавляем первое (left) в списке дерево в подсчет
	uniqueTrees[trees[0]]++
	uniqueCount := 1 // Счетчик уникальных видов деревьев в текущем отрезке

	for right < N {

		leftElem, rightElem := trees[left], trees[right]

		if leftElem == rightElem {
			left++
			fmt.Println("left:", left)
			// удаляем лишние повторения чтобы сдвинуть левый указ-ль
			// 1 2 2 2 3 1 2 для теста
			for leftElem = trees[left]; uniqueTrees[leftElem] > 1; {
				//	fmt.Println("uniqueTrees[leftElem]",uniqueTrees[leftElem])
				uniqueTrees[leftElem]-- // умень-м
				left++                  // сдвигаем
				leftElem = trees[left]  // находим зн-е по новому указ-ю
			}
		} else {
			if uniqueTrees[rightElem] == 0 { // т.е. первое попадание числа в списке
				//fmt.Printf("uniqueTrees[rightElem]==%v\n", uniqueTrees[rightElem])
				uniqueCount++ // счетчик уникал-и
			} // иначе когда число УЖЕ встречалось увелич-м его знач-е в массиве с 1 до 2 и т.д.
			uniqueTrees[rightElem]++ // uniqueCount++ (счетчик уникал-и) не трогаем
		}
		/*
			Когда мы находим отрезок, содержащий K различных видов деревьев, мы проверяем,
			не короче ли он текущего минимального - сравнить его длину с minLen.
			Если этот отрезок короче, то его границы и длина станут новыми минимальными значениями.
		*/
		if uniqueCount == K && right-left+1 < minLen {
			minLeft, minRight, minLen = left, right, right-left+1
			//	fmt.Println(minLeft, minRight, minLen)
		}

		// всегда двигаем правый указ-ль
		right++
	}

	// +1 везде из-за индексации нач-я с нуля
	fmt.Printf("%d %d\n", minLeft+1, minRight+1)
}
