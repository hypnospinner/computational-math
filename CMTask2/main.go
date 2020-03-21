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

func (m *Matrix) LUM() Matrix {

	var lum Matrix
	var size int

	if m.rows > m.cols {
		size = m.cols
	} else {
		size = m.rows
	}

	value := make([][]float64, size)

	for i := 0; i < size; i++ {
		value[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			value[i][j] = rand.Float64() * 100.0
		}
	}

	lum.value = value
	lum.cols = size
	lum.rows = size

	for i := 0; i < lum.rows-1; i++ {
		for l := i + 1; l < lum.rows; l++ {
			k := lum.value[l][i] / lum.value[i][i]

			for j := 0; j < lum.cols; j++ {
				lum.value[l][j] -= k * lum.value[i][j]
			}
		}
	}

	return lum
}

func main() {
	matrix := generateRandomMatrix(5, 5)

	matrix.print()

	fmt.Println()

	lum := matrix.LUM()

	lum.print()
}
