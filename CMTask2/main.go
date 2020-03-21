package main

import (
	"fmt"
	"math/rand"
)

// Matrix is a data structure that incapsulates data over 2D array
type Matrix struct {
	value [][]float64
	rows  int
	cols  int
}

func generateRandomMatrix(rows int, cols int) Matrix {
	value := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		value[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			value[i][j] = rand.Float64() * 100.0
		}
	}

	return Matrix{
		value: value,
		rows:  rows,
		cols:  cols,
	}
}

func (m *Matrix) print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			f := fmt.Sprintf("%2.2f ", m.value[i][j])

			fmt.Print(f)
		}
		fmt.Println()
	}
}

func main() {
	matrix := generateRandomMatrix(5, 5)

	(&matrix).print()
}
