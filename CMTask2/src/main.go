package main

import (
	"fmt"
	"math/rand"
)

// Matrix is some abstraction over table of row x col size with float64 values
type Matrix struct {
	value [][]float64
	rows  int
	cols  int
}

// NewSingular stands for E matrix (0 everywhere except main diagonal)
func NewSingular(rows int, cols int) Matrix {
	value := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		value[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			if i == j {
				value[i][j] = 1
			} else {
				value[i][j] = 0
			}
		}
	}

	return Matrix{
		value: value,
		rows:  rows,
		cols:  cols,
	}
}

// NewEmpty creates new Matrix filled with 0
func NewEmpty(rows int, cols int) Matrix {
	value := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		value[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			value[i][j] = 0
		}
	}

	return Matrix{
		value: value,
		rows:  rows,
		cols:  cols,
	}
}

// NewRandom creates new Matrix with random values
func NewRandom(rows int, cols int) Matrix {
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

// Copy creates another instance of matrix
func (m *Matrix) Copy() Matrix {
	result := NewEmpty(m.rows, m.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.value[i][j] = m.value[i][j]
		}
	}

	return result
}

// Print outputs a matrix in a simple table view
func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			f := fmt.Sprintf("%2.2f ", m.value[i][j])

			fmt.Print(f)
		}
		fmt.Println()
	}
}

// SwapRows changes value in 2 passed rows of the Matrix
func (m *Matrix) SwapRows(from int, to int) {
	for i := 0; i < m.cols; i++ {
		temp := m.value[from][i]
		m.value[from][i] = m.value[to][i]
		m.value[to][i] = temp
	}
}

// SwapCols changes value in 2 passed cols of the Matrix
func (m *Matrix) SwapCols(from int, to int) {
	for i := 0; i < m.rows; i++ {
		temp := m.value[i][from]
		m.value[i][from] = m.value[i][to]
		m.value[i][to] = temp
	}
}

// DotProduct returns dot product of passed Matrices
func DotProduct(left Matrix, right Matrix) Matrix {
	// TODO need check for matrixes being of appropriate size

	prod := NewEmpty(left.rows, right.cols)

	for i := 0; i < prod.rows; i++ {
		for j := 0; j < prod.cols; j++ {
			// for each element of resulting matrix (i,j)
			for k := 0; k < left.cols; k++ {
				prod.value[i][j] += left.value[i][k] * right.value[k][j]
			}
		}
	}

	return prod
}

// Abs returns absolute value of passed number
func Abs(number float64) float64 {
	if number < 0 {
		return -number
	}
	return number
}

// Add value of one matrix to another
func Add(a Matrix, b Matrix) Matrix {
	value := make([][]float64, a.rows)

	for r := 0; r < a.rows; r++ {
		value[r] = make([]float64, a.cols)

		for c := 0; c < a.cols; c++ {
			value[r][c] = a.value[r][c] + b.value[r][c]
		}
	}

	return Matrix{
		rows:  a.rows,
		cols:  a.cols,
		value: value,
	}
}

// LUDecomposition2 is some function 2
func LUDecomposition2(a Matrix) (Matrix, Matrix, Matrix, Matrix) {
	m := a.Copy()
	
	L := NewEmpty(m.rows, m.cols)
	U := NewEmpty(m.rows, m.cols)
	P, Q := NewSingular(m.rows, m.cols), NewSingular(m.rows, m.cols)

	// initialize first row in U and first column in L
	for i := 0; i < m.rows; i++ {
		max, row, col := m.value[i][i], i, i

		for r := i; r < m.rows; r++ {
			for c := i; c < m.cols; c++ {
				abs := Abs(m.value[r][c])

				if abs > max {
					max, row, col = abs, r, c
				}
			}
		}

		m.SwapRows(i, row)
		m.SwapCols(i, col)

		P.SwapRows(i, row)
		Q.SwapCols(i, col)
	}

	fmt.Println("Swapped A:")
	m.Print()

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			U.value[0][i] = m.value[0][i]
			L.value[i][0] = m.value[i][0] / U.value[0][0]

			U.value[i][j] = m.value[i][j]

			for k := 0; k < i; k++ {
				U.value[i][j] -= L.value[i][k] * U.value[k][j]
			}

			if i > j {
				L.value[j][i] = 0
			} else {
				L.value[j][i] = m.value[j][i]

				for k := 0; k < i; k++ {
					L.value[j][i] -= L.value[j][k] * U.value[k][i]
				}

				L.value[j][i] /= U.value[i][i]
			}
		}
	}

	return L, U, P, Q
}

func main() {
	A := Matrix{
		rows: 3,
		cols: 3,
		value: [][]float64{
			{5.0, 8.0, 1.0},
			{3.0, -2.0, 6.0},
			{2.0, 1.0, -1.0},
		},
	}

	L, U, P, Q := LUDecomposition2(A)
	fmt.Println("L:")
	L.Print()
	fmt.Println("U:")
	U.Print()
	fmt.Println("P:")
	P.Print()
	fmt.Println("Q:")
	Q.Print()

	PAQ := DotProduct(DotProduct(P, A), Q)

	fmt.Println("PAQ:")

	PAQ.Print()

	LU := DotProduct(L, U)
	fmt.Println("LU:")

	LU.Print()
	fmt.Println("A:")

	A.Print()
}
