package main

import (
	"fmt"

	"github.com/mac-lawson/tensorgo/neuralnetwork"
)

func main() {
	// tn := tensor.Gotensor_dtypefloat64{
	// 	Data: [][]float64{
	// 		{6.7, 4.5},
	// 		{4.5, 4.5},
	// 	},
	// }
	// wei := neuralnetwork.Weights{
	// 	W: []float64{4.5, 3.3},
	// 	B: 4.5,
	// }
	// resp, err := neuralnetwork.Neuron(tn, &wei, 1)
	// if err != nil {
	// 	fmt.Println(err)

	// } else {
	// 	fmt.Println(resp)
	// }

	y, err := neuralnetwork.InnerNeuron(3.4, 1.0, 3.0, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y)
	}
}
