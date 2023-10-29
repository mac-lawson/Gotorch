package neuralnetwork

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/mac-lawson/gotorch/tensor"
)

/*
SimpleNeuralNetwork function

This function runs a simulated simple neural network.
It has three parts, the:
- input layer
- hidden layer
- output layer
* Activation Function
* Options:
* 1: Sigmoid
* 2: Tanh
* 3. reLu

Verbose mode (bool):
If you would like to see a more verbose output of the result, enable verbose mode.
*/
func BinaryClassification(neurons uint64, activator uint8, wei Weights, tensors tensor.Gotensor_dtypefloat64, verbose bool) (*NeuronOutputArray, error) {
	r := rand.New(rand.NewSource(99))
	// generate output so that we can average out the outputs of the function
	Y := NeuronOutputArray{
		Y: []float64{},
	}

	for i := 0; i < len(tensors.Data); i++ {
		wei.W = append(wei.W, r.Float64())
	}
	if len(wei.W) == 0 {
		return &Y, errors.New("value of 0 error")
	}

	// for every neyron
	for i := 0; i < int(neurons); i++ {
		// for every array of tensors
		for i := 0; i < len(tensors.Data); i++ {
			// for every tensor inside of the array
			for tensor := 0; tensor < len(tensors.Data[i]); tensor++ {
				// run it through the InnerNeuron function with the data and weights
				// TODO change this so that you can chose the activiation function (1 is default)
				output, err := InnerNeuron(tensors.Data[i][tensor], wei.W[tensor], wei.B, activator)
				if err != nil {
					// quit and raise an error if one is encountered during the running of the InnerNeuron function.
					// It returns all data recived so far before the error to ensure integrity.
					return &Y, errors.New(err.Error())
				} else {
					// TODO change this so that there is a verbose mode and a regular output mode
					fmt.Println("Evolution Completed \t Set:", i, "Convolution", tensor)
					fmt.Println("Output:", output)
					// add the result to the output array so that a value of the function can be return of the average of all of the results
					Y.Y = append(Y.Y, output)

				}
			}
		}
	}
	if verbose {
		fmt.Println("\033[32m", "Network Finished Running")
		fmt.Println("\033[34m", "\tInformation")
		fmt.Println("Total Tensor Arrays", len(tensors.Data))
		fmt.Println("Total Weights", len(wei.W))
	}

	return &Y, nil
}
