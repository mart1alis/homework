package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 5, 9, 10},
		{11, 4, 13, 14, 15},
		{16, 17, 18, 9, 20},
		{21, 22, 23, 0, 25},
	}

	mainDiagonalSum := 0
	for i := 0; i < 5; i++ {
		mainDiagonalSum += matrix[i][i]
	}


	fmt.Println("Исходная матрица:")
	printMatrix(matrix)

	transposedMatrix := transposeMatrix(matrix)
	determinant := calculateDeterminant(matrix)

	fmt.Printf("Определитель матрицы: %d\n", determinant)

	fmt.Println("Транспонированная матрица:")
	printMatrix(transposedMatrix)

	fmt.Printf("Сумма главной диагонали: %d\n", mainDiagonalSum)
}

func transposeMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}
func calculateDeterminant(matrix [][]int) int {

	if len(matrix) == 1 {
		return matrix[0][0]
	}

	if len(matrix) == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0
	for col := 0; col < len(matrix); col++ {

		minor := getMinor(matrix, 0, col)

		minorDet := calculateDeterminant(minor)

		sign := 1
		if col%2 != 0 {
			sign = -1
		}

		det += sign * matrix[0][col] * minorDet
	}

	return det
}

func getMinor(matrix [][]int, row, col int) [][]int {
	// Создаем новую подматрицу
	minor := make([][]int, len(matrix)-1)
	for i := range minor {
		minor[i] = make([]int, len(matrix)-1)
	}

	r := 0
	for i := 0; i < len(matrix); i++ {
		if i == row {
			continue
		}
		c := 0
		for j := 0; j < len(matrix); j++ {
			if j == col {
				continue
			}
			minor[r][c] = matrix[i][j]
			c++
		}
		r++
	}

	return minor
}
