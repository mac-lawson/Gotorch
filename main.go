package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/neuralnetwork"
	"github.com/mac-lawson/gotorch/tensor"
)

func main() {

	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{1.3, 1.4},
			{1.5},
			{2.3, 5.6, 10},
		},
	}

	result, er := neuralnetwork.ConvolutionalNeuralNetwork(tn, 4, 1)

	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println(result)
	}

}
