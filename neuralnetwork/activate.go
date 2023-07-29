package neuralnetwork

import "math"

func Activation(option int32, input float64) float64 {
	switch option {
	default:
		return 0.0
	case 1:
		return sigmoid(input)
	case 2:
		return tanh(input)
	case 3:
		return reLU(input)
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
