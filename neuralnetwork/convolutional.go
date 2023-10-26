package neuralnetwork

import (
	"errors"
	"fmt"

	"github.com/mac-lawson/gotorch/tensor"
)

/*
ConvolutionalNeuralNetwork function
* Activation Function
* Options:
* 1: Sigmoid
* 2: Tanh
* 3. reLu
This function runs a simulated convolutional neural network.
*/
func ConvolutionalNeuralNetwork(tensors tensor.Gotensor_dtypefloat64, weights Weights, layers uint64, activator uint8) (*ConvolutionalOutputArray, error) {
	Y := ConvolutionalOutputArray{
		Y: []NeuronOutputArray{},
	}
	// r := rand.New(rand.NewSource(99))
	for layer := 0; layer < int(layers); layer++ {
		fmt.Println("\033[34m", "Layer:", layer, "\x1b[0m")
		result, err := NeuralNetwork(3, activator, weights, tensors, false)
		if err != nil {
			return &Y, errors.New(err.Error())
		} else {
			Y.Y = append(Y.Y, *result)
		}
	}
	return &Y, nil
}
