package tokenizer

import (
	"github.com/mac-lawson/gotorch/cryptomath"
	"github.com/mac-lawson/gotorch/tensor"
)

func Tokenizer_str(data [][]string) *TokenizedData_str {
	m := make(map[string]int64)
	count := make(map[string]uint8)
	result := make([][]int64, len(data))

	for i, row := range data {
		result[i] = make([]int64, len(row))
		for j, str := range row {
			if _, exists := m[str]; !exists {
				m[str], _ = cryptomath.CryptoRandomInt64() // Generate a random int64
			}
			count[str] += 1
			result[i][j] = m[str]
		}
	}
	return &TokenizedData_str{
		OrigData:      data,
		TokenizedData: result,
		EncodingMap:   m,
		TotalTokens:   count,
	}
}

/*
*
* Tokenizer to Tensor Functions
*   Allows you to translate tokenized data sets to tensor data.
*
 */

func (td *TokenizedData_str) ToTensor() *tensor.Gotensor_dtypefloat64 {
	tensorarray := make([][]float64, len(td.TokenizedData))
	for i, row := range td.TokenizedData {
		tensorarray[i] = make([]float64, len(row))
		for j, val := range row {
			tensorarray[i][j] = float64(val)
		}
	}
	return tensor.GoTensorFloat64(tensorarray)
}
