package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/neuralnetwork"
	"github.com/mac-lawson/gotorch/tensor"
)

func main() {

	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{1.1, 22.2, 3.3},
			{4.7, 20.8, 34.5},
			{13.4, 222.3, 31.2},
		},
	}

	wi := neuralnetwork.Weights{
		W: []float64{1.0},
		B: 4.4,
	}

	result, _ := neuralnetwork.ConvolutionalNeuralNetwork(tn, wi, 10, 3)
	res := result.Y
	fmt.Println(res[2])
}
