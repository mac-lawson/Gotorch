# Tensor

## Data Types
### Float64
```golang
package main

import (
	"github.com/mac-lawson/tensorgo/tensor"
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
}
```
### Int64
```golang
package main

import (
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {
	tn := tensor.Gotensor_dtypeint64{
		Data: [][]int64{
			{5, 6, 7},
			{4, 6, 7},
		},
	}
}
```
### String
```golang
package main

import (
	"github.com/mac-lawson/tensorgo/tensor"
)

func main() {
	tn := tensor.Gotensor_dtypestring{
		Data: [][]string{
			{"this", "is", "a"},
			{"string", "tensor"},
		},
	}
}

```