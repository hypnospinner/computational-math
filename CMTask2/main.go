package main

import (
	"fmt"
	"math"
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

// LUDecomposition returns 2 matrixes L - lower triangular matrix & U - upper triangular matrix
// which A = LU
func (m *Matrix) LUDecomposition() (Matrix, Matrix) {
	var lum Matrix
	var uum Matrix
	var size int

	if m.cols > m.rows {
		size = m.rows
	} else {
		size = m.cols
	}

	lum = Matrix{
		value: make([][]float64, size),
		rows:  size,
		cols:  size,
	}

	uum = Matrix{
		value: make([][]float64, size),
		rows:  size,
		cols:  size,
	}

	for i := 0; i < size; i++ {
		lum.value[i] = make([]float64, size)
		uum.value[i] = make([]float64, size)

		for j := 0; j < size; j++ {
			uum.value[i][j] = m.value[i][j]
			lum.value[i][j] = 0.0
		}
	}

	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			lum.value[j][i] = uum.value[j][i] / uum.value[i][i]
		}
	}

	for k := 1; k < size; k++ {
		for i := k - 1; i < size; i++ {
			for j := i; j < size; j++ {
				lum.value[j][i] = uum.value[j][i] / uum.value[i][i]
			}
		}

		for i := k; i < size; i++ {
			for j := k - 1; j < size; j++ {
				uum.value[i][j] = uum.value[i][j] - lum.value[i][k-1]*uum.value[k-1][j]
			}
		}
	}

	return lum, uum
}

// Norm is maximum value of summ of element of over each row
func (m *Matrix) Norm() float64 {
	var cur float64
	var norm float64
	norm = 0.0

	for i := 0; i < m.rows; i++ {
		cur = 0.0

		for j := 0; j < m.cols; j++ {
			cur += math.Abs(m.value[i][j])
		}

		if cur > norm {
			norm = cur
		}
	}

	return norm
}

func multiply(a *Matrix, b *Matrix) Matrix {

	if a.cols != b.rows {
		panic("Can't multiply matrixes of this size")
	}

	result := Matrix{
		value: make([][]float64, a.rows),
		rows:  a.rows,
		cols:  b.cols,
	}

	for i := 0; i < result.rows; i++ {
		result.value[i] = make([]float64, result.cols)

		for j := 0; j < result.cols; j++ {
			// element in i rows & j cols is a sum of multiplilcation of
			// elements in i rows of matrix a
			// element in j row of matrix b

			result.value[i][j] = 0.0

			for k := 0; k < a.cols; k++ {
				result.value[i][j] += a.value[i][k] * b.value[k][j]
			}
		}
	}

	return result
}

func main() {
	m1 := Matrix{
		value: [][]float64{
			{3, 5},
			{6, 7},
		},
		rows: 2,
		cols: 2,
	}

	lum, uum := m1.LUDecomposition()

	m1.print()
	fmt.Println()

	lum.print()
	fmt.Println()

	uum.print()
	fmt.Println()
	// r := multiply(&m1, &m2)

	// r.print()
}
