# Gotorch
## Building simple neural networks in Golang

### Gotensors
Gotensors are basically the same as a numpy array. 

There are three different data types of gotensors:

- Gotensor_dtypefloat64 | float64
- Gotensor_dtypeint64 | int64
- Gotensor_dtypestring | string

### Printing formatted go tensors

- Readfloat64 | float64
- Readint64 | int64
- Readstring| string

Sample output:

```
5.3       6.6     7.8
4.4     6.6     7.9
3.5     2.6     7.8
7       6.6     7.7
8.5     4.6     7.6
9.3     6.6     7.4
5.7     1.6     7.8
```
#### sample main.go file
```go
package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/neuralnetwork"
	"github.com/mac-lawson/gotorch/tensor"
)

func main() {
	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{5.3, 6.6, 7.8},
			{4.4, 6.6, 7.9},
			{3.5, 2.6, 7.8},
			{7.0, 6.6, 7.7},
			{8.5, 4.6, 7.6},
			{9.3, 6.6, 7.4},
			{5.7, 1.6, 7.8},
		},
	}
	wei := neuralnetwork.Weights{
		W: []float64{4.4, 5.5, 6.6},
		B: 1.0,
	}

	y, err := neuralnetwork.RealNeuron(tn, &wei, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y.Y)
	}

}

```
