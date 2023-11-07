package main

import (
	"github.com/mac-lawson/gotorch/matrix"
	"github.com/mac-lawson/gotorch/tensor"
)

func main() {
	m := matrix.Matrix{
		Data: [][]float64{
			{1.0, 2, 3},
			{3.5, 9.0, 8.6},
		},
	}
	m1 := matrix.Scalar(4, &m)
	tn := tensor.Matrix_to_Tensor(*m1)

	tensor.Readfloat64(*tn)
}
