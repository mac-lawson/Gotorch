package neuralnetwork

import "math/rand"

func MeanOfNetworkResultPerceptron(result NeuronOutputArray) *NeuronOutput {
	finresult := NeuronOutput{
		Y: 0,
	}

	sum := 0.0
	numvalues := 0
	for index, value := range result.Y {
		sum += value
		numvalues += index
	}
	final := sum / float64(numvalues)
	finresult.Y = final

	return &finresult
}
// TODO random weights
func RandomWeights(xsize int32) *Weights {
	weightarray := make([]float64, xsize)
	for i := range weightarray {
		weightarray[i] = rand.Float64()

	}
	b := rand.Float64()

	return &Weights{
		W: weightarray, 
		B: b,
		

	}

}
