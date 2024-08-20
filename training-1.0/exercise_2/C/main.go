package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // Считываем три строки
    var lines [3]string
    for i := range lines {
        if scanner.Scan() {
            lines[i] = scanner.Text()
        } else {
            break
        }
    }

    // Парсим вторую строку в массив чисел
    numbers := make([]int, 0)
    for _, numStr := range strings.Split(lines[1], " ") {
        num, err := strconv.Atoi(numStr)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Ошибка при преобразовании числа:", err)
            return
        }
        numbers = append(numbers, num)
    }

    // Находим значение, ближайшее к числу в третьей строке
    targetNum, err := strconv.Atoi(lines[2])
    if err != nil {
        fmt.Fprintln(os.Stderr, "Ошибка при преобразовании числа:", err)
        return
    }

    closestNum := numbers[0]
    closestDiff := math.Abs(float64(numbers[0] - targetNum))
    for _, num := range numbers[1:] {
        diff := math.Abs(float64(num - targetNum))
        if diff < closestDiff {
            closestNum = num
            closestDiff = diff
        }
    }

    fmt.Println(closestNum)

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Ошибка при чтении ввода:", err)
    }
}