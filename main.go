package main

import (
	"fmt"

	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {

	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{1.3, 1.4},
			{1.5},
			{2.3, 5.6, 10},
		},
	}

	result, er := neuralnetwork.SimpleNeuralNetwork(2, 2, tn, true)

	if er != nil {
		fmt.Println(er)
	} else {
		tensor.Readfloat64(fintn)
	}

}
