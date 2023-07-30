package neuralnetwork

import (
	// "errors"

	// "github.com/mac-lawson/tensorgo/errorhandling"

	"fmt"

	"github.com/mac-lawson/tensorgo/tensor"
)

type NeuronOutput struct {
	Y float64
}

/* TODO Generate a weight for the tensor going into the function
* - array of weight value (w) needs to be associated with the tensor (i.e tensor[1]== w[1])
* - b value is assigned as a seperate input for every neuron
 */
type Weights struct {
	W []float64
	B float64
	// wMin float64
}

/*
	 TODO Take in a tensor
		(CURRENT) - Prints out the result of each epoch
		(FUTURE) - Should return an array containing the result of each epoch
*/

// FUTURE Neuron() function

// func Neuron(input tensor.Gotensor_dtypefloat64, weights *Weights, activator int32) (*NeuronOutput, error) {
// 	if len(input.Data[0]) != len(weights.W) {
// 		return &NeuronOutput{Y: 0.0}, errors.New(errorhandling.TensorNotMatching())
// 	} else {
// 		return &NeuronOutput{Y: 0.0}, nil
// 	}
// }

// CURRENT Neuron() function

func Neuron(input tensor.Gotensor_dtypefloat64, weights *Weights, activation uint8) error {
	// verify that each tensor array has the same value of the weights
	for i := 0; i < len(input.Data); i++ {
		for ii := 0; ii < len(input.Data[i]); ii++ {
			// ended work here, still need to iterate through each epoch
			fmt.Println(input.Data[i][ii])
		}
	}
	return nil
}
