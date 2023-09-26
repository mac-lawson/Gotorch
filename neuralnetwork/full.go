package neuralnetwork

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/mac-lawson/tensorgo/tensor"
)

/*
ConvolutionalNeuralNetwork function

This function runs a simulated convolutional neural network.
*/
func ConvolutionalNeuralNetwork(layers uint64, activator uint8) int64 {
	return 0
}

/*
SimpleNeuralNetwork function

This function runs a simulated simple neural network.
It has three parts, the:
- input layer
- hidden layer
- output layer
*/
func SimpleNeuralNetwork(neurons uint64, activator uint8, tensors tensor.Gotensor_dtypefloat64, length uint64) (string, error) {
	r := rand.New(rand.NewSource(99))

	// generate weights
	wei := Weights{
		W: []float64{},
		B: 1.0,
	}
	for i := 0; i < int(length); i++ {
		wei.W = append(wei.W, r.Float64())
	}
	fmt.Println(wei.W)

	if len(wei.W) == 0 {
		return "", errors.New("value of 0 error")
	}

	// fmt.Println(wei)
	return "This function is currently under development", nil
}
