package main

import (
	"fmt"

	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {
	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{5.3, 6.6, 7.8},
			{4.4, 6.6, 7.9},
			{3.5, 2.6, 7.8},
			{7.0, 6.6, 7.7},
			{8.5, 4.6, 7.6},
			{9.3, 6.6, 7.4},
			{5.7, 1.6, 7.8},
		},
	}
	wei := neuralnetwork.Weights{
		W: []float64{4.4, 5.5, 6.6},
		B: 1.0,
	}

	y, err := neuralnetwork.RealNeuron(tn, &wei, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y.Y)
	}

}
