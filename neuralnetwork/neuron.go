package neuralnetwork

import (
	"errors"

	"github.com/mac-lawson/tensorgo/errorhandling"
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

/* TODO Take in a tensor
*	- Each value of a tensor should be treated as an individual input
*   - Each should be processed by every neuron in the neural network
*	- Should return a neuron output (y)
*   - Supports data types strings, float64, and int64. LTS needs 32-bit support as well :)
*   - Runs the activation function!
 */
func Neuron(input tensor.Gotensor_dtypefloat64, weights *Weights, activator int32) (*NeuronOutput, error) {
	if len(input.Data[0]) != len(weights.W) {
		return &NeuronOutput{Y: 0}, errors.New(errorhandling.TensorNotMatching())
	} else {
		return &NeuronOutput{Y: 4.9}, nil
	}
}
