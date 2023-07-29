package main

import (
	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {
	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{6.7, 4.5},
			{4.5, 4.5},
		},
	}
	wei := neuralnetwork.Weights{
		W: []float64{4.5, 3.3},
		B: 4.5,
	}
	neuralnetwork.Neuron(tn, &wei, 1)
}
