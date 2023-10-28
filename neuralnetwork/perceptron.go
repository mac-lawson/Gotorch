package neuralnetwork

import (
	"errors"

	"github.com/mac-lawson/gotorch/tensor"
)

func Perceptron(tensor tensor.Gotensor_dtypefloat64, weights Weights, activation int16) (*NeuronOutput, error) {
	preresult := NeuronOutputArray{
		Y: []float64{},
	}
	for index, _ := range tensor.Data {
		for weightindex, value := range tensor.Data[index] {
			res, err := InnerNeuron(value, weights.W[weightindex], weights.B, uint8(activation))
			if err != nil {
				return &NeuronOutput{Y: 0}, errors.New("Error inside neuron")
			} else {
				preresult.Y = append(preresult.Y, res)
			}
		}
	}
	return MeanOfNetworkResultPerceptron(preresult), nil

}

func MultiLayerPerceptron() {

}
