# Gotorch
Version 0.1.1

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/mac-lawson/gotorch/.github%2Fworkflows%2Fgo.yml)
![GitHub License](https://img.shields.io/github/license/mac-lawson/gotorch)

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/mac-lawson/gotorch)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/7558ccce51f34572a17710a3b2736fce)](https://app.codacy.com/gh/mac-lawson/Gotorch/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)


## Simulating simple neural networks in Golang

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
#### Getting started
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

## Features and Change log

TODO
- refactor functions that do not return a Return type
- refactor loops not optimized
