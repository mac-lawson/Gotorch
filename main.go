package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/matrix"
	"github.com/mac-lawson/gotorch/neuralnetwork"
	"github.com/mac-lawson/gotorch/tensor"
)

func main() {
	n2 := tensor.RandomFloat64(4, 4)
	wi := neuralnetwork.RandomWeights(3)
	neuralnetwork.BinaryClassification(2, 1, *wi, *n2, true)
	m1 := matrix.Matrix{
		Internal: [][]float64{
			{1.0, 2},
			{3.5, 9.0},
		},
	}
	m2 := matrix.Matrix{
		Internal: [][]float64{
			{1.0, 2},
			{3.5, 9.0},
		},
	}

	m3 := matrix.Add(&m1, &m2)
	fmt.Println(m3)

}
