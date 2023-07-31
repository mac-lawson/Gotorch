package neuralnetwork

import (
	// "errors"

	// "github.com/mac-lawson/tensorgo/errorhandling"

	"errors"
	"fmt"

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

/*
	 TODO Take in a tensor
		(TEST) - Prints out the result of each epoch
		(FUTURE) - Should return an array containing the result of each epoch
*/

// FUTURE Neuron() function

func RealNeuron(input tensor.Gotensor_dtypefloat64, weights *Weights, activator int32) (*NeuronOutput, error) {
	// the total value of the weighted inputs to the neuron, this willbe added to the bias and run through the activiation function
	var weightedInputs = 0.0
	for i := 0; i < len(input.Data); i++ {
		for ii := 0; ii < len(input.Data[i]); ii++ {
			if len(input.Data[i]) > len(weights.W) {
				return &NeuronOutput{0}, errors.New(errorhandling.TensorNotMatching())
			} else {
				weightedInputs += float64(input.Data[i][ii] * weights.W[ii])

			}
		}
	}
	//TODO need to make sure the entire Y value is not rounded
	resultOfActiviationFunction, err := InnerNeuron(weightedInputs, 1.0, weights.B, uint8(activator))
	fmt.Println(resultOfActiviationFunction)
	if err != nil {
		return &NeuronOutput{0}, errors.New(errorhandling.ActiviationFunctionOptionNotProvided("n/a"))
	} else {
		return &NeuronOutput{float64(resultOfActiviationFunction)}, nil
	}
}

// TEST Neuron() function
// id: the ID of the neuron to ensure
func NeuronTest(input tensor.Gotensor_dtypefloat64, weights *Weights, activation uint8, id int64) error {
	// epoch counter
	var epoch uint8 = 0
	// verify that each tensor array has the same value of the weights
	for i := 0; i < len(input.Data); i++ {
		for ii := 0; ii < len(input.Data[i]); ii++ {
			if len(input.Data[i]) > len(weights.W) {
				return errors.New(errorhandling.TensorNotMatching())
			} else {
				yOutput, err := InnerNeuron(input.Data[i][ii], weights.W[ii], weights.B, activation)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Epoch", epoch, ":", yOutput)

				}
				epoch++
			}

		}
	}
	return nil
}
