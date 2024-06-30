package neuralnetwork

import (
	"errors"
	"fmt"

	"github.com/mac-lawson/gotorch/cryptomath"
	"github.com/mac-lawson/gotorch/tensor"
)

func BinaryClassification(neurons uint64, activator uint8, wei Weights, tensors tensor.Tensorfloat64, verbose bool) (*NeuronOutputArray, error) {
	// generate output so that we can average out the outputs of the function
	Y := NeuronOutputArray{
		Y: []float64{},
	}

	for i := 0; i < len(tensors); i++ {
		randomFloat, err := cryptomath.CryptoRandomFloat64()
		if err != nil {
			return &Y, err
		}
		wei.W = append(wei.W, randomFloat)
	}
	if len(wei.W) == 0 {
		return &Y, errors.New("value of 0 error")
	}

	// for every neuron
	for i := 0; i < int(neurons); i++ {
		// for every array of tensors
		for i := 0; i < len(tensors); i++ {
			// for every tensor inside of the array
			for tensor := 0; tensor < len(tensors[i]); tensor++ {
				// run it through the InnerNeuron function with the data and weights
				output, err := InnerNeuron(tensors[i][tensor], wei.W[tensor], wei.B, activator)
				if err != nil {
					// quit and raise an error if one is encountered during the running of the InnerNeuron function.
					// It returns all data received so far before the error to ensure integrity.
					return &Y, errors.New(err.Error())
				} else {
					if verbose {
						fmt.Println("Evolution Completed \t Set:", i, "Convolution", tensor)
						fmt.Println("Output:", output)
					}
					// add the result to the output array so that a value of the function can be return of the average of all of the results
					Y.Y = append(Y.Y, output)
				}
			}
		}
	}
	if verbose {
		fmt.Println("\033[32m", "Network Finished Running")
		fmt.Println("\033[34m", "\tInformation")
		fmt.Println("Total Tensor Arrays", len(tensors))
		fmt.Println("Total Weights", len(wei.W))
	}

	return &Y, nil
}
