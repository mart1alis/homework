package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var A, B, C int
	if scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d %d", &A, &B, &C)
	} else if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Вычисление дискриминанта
	D := B*B - 4*A*C

	fmt.Printf("A: %d, B: %d, C: %d\n", A, B, C)
	fmt.Printf("Дискриминант: %d\n", D)

	// Проверка корней уравнения
	if D > 0 {
		x1 := (-float64(B) + math.Sqrt(float64(D))) / (2 * float64(A))
		x2 := (-float64(B) - math.Sqrt(float64(D))) / (2 * float64(A))
		fmt.Printf("Корни уравнения: x1 = %.2f, x2 = %.2f\n", x1, x2)
	} else if D == 0 {
		x := -float64(B) / (2 * float64(A))
		fmt.Printf("Единственный корень уравнения: x = %.2f\n", x)
	} else {
		fmt.Println("Корней нет (D < 0)")
	}
}