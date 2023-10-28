package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/neuralnetwork"
	"github.com/mac-lawson/gotorch/tensor"
)

func main() {

	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{1.1, 2.2, 3.3, 4.4},
			{2.2, 3.3, 4.4, 5.5},
		},
	}
	wei := neuralnetwork.Weights{
		W: []float64{1, 1, 1, 1},
		B: 1,
	}
	result, err := neuralnetwork.Perceptron(tn, wei, 3)

	if err != nil {
	} else {
		fmt.Println(result.Y)

	}
}
