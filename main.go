package main

import (
	"fmt"

	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {
	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{5.0, 6.6, 7.8},
			{4.4, 5.6, 8.9},
		},
	}
	wei := neuralnetwork.Weights{
		W: []float64{4.4, 5.5, 6.6},
		B: 1.0,
	}

	err := neuralnetwork.Neuron(tn, &wei, 1)
	if err != nil {
		fmt.Println(err)
	}
}
