package neuralnetwork

import (
	"errors"

	"github.com/mac-lawson/gotorch/tensor"
)

func Perceptron(inputData tensor.Tensorfloat64, weights Weights, activationFunction uint8) (*NeuronOutput, error) {
	preResult := NeuronOutputArray{
		Y: []float64{},
	}

	for _, inputVector := range inputData {
		var neuronOutputs []float64

		for i, inputValue := range inputVector {
			result, err := InnerNeuron(inputValue, weights.W[i], weights.B, activationFunction)
			if err != nil {
				return &NeuronOutput{Y: 0}, errors.New("Error encountered within neuron")
			}
			neuronOutputs = append(neuronOutputs, result)
		}

		preResult.Y = append(preResult.Y, neuronOutputs...)
	}

	return MeanOfNetworkResultPerceptron(preResult), nil
}
