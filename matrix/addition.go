package matrix

import "fmt"
// test
func Add(matrixA *Matrix, matrixB *Matrix) *Matrix {
	if len(matrixA.Internal) != len(matrixB.Internal) || len(matrixA.Internal[0]) != len(matrixB.Internal[0]) {
		fmt.Println("Matrices must have the same dimensions for addition.")
		return nil
	}

	result := &Matrix{
		Internal: make([][]float64, len(matrixA.Internal)),
	}

	for i := 0; i < len(matrixA.Internal); i++ {
		result.Internal[i] = make([]float64, len(matrixA.Internal[i]))
		for j := 0; j < len(matrixA.Internal[i]); j++ {
			result.Internal[i][j] = matrixA.Internal[i][j] + matrixB.Internal[i][j]
		}
	}

	return result
}
