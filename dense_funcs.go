package matrix

import (
	"math"
)

// Apply the given function on each element of the matrix
func (A *DenseMatrix) ApplyFunc(f func(inp float64) float64) Matrix {

	B := A.Copy()
	for i := 0; i < B.rows; i++ {
		index := i * B.step
		for j := 0; j < B.cols; j++ {
			B.elements[index] = f(B.elements[index])
			index++
		}
	}

	return B
}

// Apply Tanh to all the elements of the matrix
func (A *DenseMatrix) Tanh() Matrix {
	B := A.ApplyFunc(math.Tanh)
	return B
}

// Apply Sigmoid to all the elements of the matrix
func (A *DenseMatrix) Sigmoid() Matrix {

	f := func(sum float64) float64 {
		return 1.0 / (1.0 + math.Exp(-sum))
	}

	B := A.ApplyFunc(f)
	return B
}
