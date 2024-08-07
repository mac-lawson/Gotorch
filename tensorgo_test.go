package main

import (
	"fmt"
	"testing"

	"github.com/mac-lawson/gotorch/neuralnetwork"
	"github.com/mac-lawson/gotorch/tensor"
)

/*
      _____                    _____                    _____                _____                    _____
     /\    \                  /\    \                  /\    \              /\    \                  /\    \
    /::\    \                /::\    \                /::\    \            /::\    \                /::\    \
    \:::\    \              /::::\    \              /::::\    \           \:::\    \              /::::\    \
     \:::\    \            /::::::\    \            /::::::\    \           \:::\    \            /::::::\    \
      \:::\    \          /:::/\:::\    \          /:::/\:::\    \           \:::\    \          /:::/\:::\    \
       \:::\    \        /:::/__\:::\    \        /:::/__\:::\    \           \:::\    \        /:::/__\:::\    \
       /::::\    \      /::::\   \:::\    \       \:::\   \:::\    \          /::::\    \       \:::\   \:::\    \
      /::::::\    \    /::::::\   \:::\    \    ___\:::\   \:::\    \        /::::::\    \    ___\:::\   \:::\    \
     /:::/\:::\    \  /:::/\:::\   \:::\    \  /\   \:::\   \:::\    \      /:::/\:::\    \  /\   \:::\   \:::\    \
    /:::/  \:::\____\/:::/__\:::\   \:::\____\/::\   \:::\   \:::\____\    /:::/  \:::\____\/::\   \:::\   \:::\____\
   /:::/    \::/    /\:::\   \:::\   \::/    /\:::\   \:::\   \::/    /   /:::/    \::/    /\:::\   \:::\   \::/    /
  /:::/    / \/____/  \:::\   \:::\   \/____/  \:::\   \:::\   \/____/   /:::/    / \/____/  \:::\   \:::\   \/____/
 /:::/    /            \:::\   \:::\    \       \:::\   \:::\    \      /:::/    /            \:::\   \:::\    \
/:::/    /              \:::\   \:::\____\       \:::\   \:::\____\    /:::/    /              \:::\   \:::\____\
\::/    /                \:::\   \::/    /        \:::\  /:::/    /    \::/    /                \:::\  /:::/    /
 \/____/                  \:::\   \/____/          \:::\/:::/    /      \/____/                  \:::\/:::/    /
                           \:::\    \               \::::::/    /                                 \::::::/    /
                            \:::\____\               \::::/    /                                   \::::/    /
                             \::/    /                \::/    /                                     \::/    /
                              \/____/                  \/____/                                       \/____/

*/

func TestTensorAllDataTypes(t *testing.T) {
	tn := tensor.Tensorfloat64{
		{3.3, 4.5, 6.6},
		{4.4, 4.9, 1.3},
		{3.3, 4.5, 6.6},
		{3.3, 4.5, 6.6},
	}
	if len(tn) != 4 {
		t.Errorf("The length of the unit test tensor #1 should be 4")
	}

	tn1 := tensor.Tensorfloat64{
		{3, 4, 6},
		{4, 4, 1},
	}
	if len(tn1) != 2 {
		t.Errorf("The length of the unit test tensor #2 should be 2")
	}
}

func TestActivationFunctionSigmoid(t *testing.T) {
	result, err := neuralnetwork.Activation(1, 4.5)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if result != 0.9890130573694068 {
			fmt.Println(result)
			t.Errorf("sigmoid(4.5) should return 0.9890130573694068")
		}
	}
}

func TestActivationFunctionreLu(t *testing.T) {
	result, err := neuralnetwork.Activation(3, 4.5)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if result != 4.5 {
			fmt.Println(result)
			t.Errorf("reLu(4.5) should return 4.5")
		}
	}
}

func TestActivationFunctionTanh(t *testing.T) {
	result, err := neuralnetwork.Activation(2, 4.5)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		if result != 0.9997532108480275 {
			fmt.Println(result)
			t.Errorf("tanh(4.5) should return 0.9997532108480275")
		}
	}
}
