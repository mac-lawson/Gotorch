package main

import (
	"testing"

	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
)

func TestTensorAllDataTypes(t *testing.T) {
	tn := &tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{3.3, 4.5, 6.6},
			{4.4, 4.9, 1.3},
			{3.3, 4.5, 6.6},
			{3.3, 4.5, 6.6},
		},
	}
	if len(tn.Data) != 4 {
		t.Errorf("The length of the unit test tensor #1 should be 4")
	}

	tn1 := &tensor.Gotensor_dtypeint64{
		Data: [][]int64{
			{3, 4, 6},
			{4, 4, 1},
		},
	}
	if len(tn1.Data) != 2 {
		t.Errorf("The length of the unit test tensor #2 should be 2")
	}
}

func TestActivationFunctionSigmoid(t *testing.T) {
	result, err := neuralnetwork.Activation(1, 4.5)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if result != 0.9890130573694068 {
			t.Errorf("sigmoid(4.5) should return 0.9890130573694068")
		}
	}
}
