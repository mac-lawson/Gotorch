package neuralnetwork

import (
	"errors"
	"math"

	"github.com/mac-lawson/tensorgo/errorhandling"
)

/*
* Activation Function
* Options:
* 1: Sigmoid
* 2: Tanh
* 3. reLu
 */
func Activation(option uint8, input float64) (float64, error) {
	switch option {
	default:
		return 0.0, errors.New(errorhandling.ActiviationFunctionOptionNotProvided("n/a"))
	case 1:
		return sigmoid(input), nil
	case 2:
		return tanh(input), nil
	case 3:
		return reLU(input), nil
	}

}

// Sigmoid function takes a float64 as input and returns the result of the sigmoid activation.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// Tanh function takes a float64 as input and returns the result of the hyperbolic tangent activation.
func tanh(x float64) float64 {
	return math.Tanh(x)
}

// ReLU function takes a float64 as input and returns the result of the Rectified Linear Unit activation.
func reLU(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}
