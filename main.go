package main

import (
	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {

	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{1.3, 1.4},
			{1.5, 1.6},
		},
	}
	// w := neuralnetwork.Weights{
	// 	W: []float64{1.3, 1.4},
	// 	B: 2,
	// }
	neuralnetwork.SimpleNeuralNetwork(3, 1, tn)
}
