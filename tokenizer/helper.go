package tokenizer

import (
	"fmt"
)

func (td *TokenizedData_str) Helper() {
	for i := range td.EncodingMap {
		fmt.Print(i, " is assigned to token ")
		fmt.Print(td.EncodingMap[i], " and is repeated ")
		fmt.Println(td.TotalTokens[i], "times")
	}
}

func (td *TokenizedData_str) GetTokenCount(i string) uint8 {
	return td.TotalTokens[i]
}

func (td *TokenizedData_str) MaxValue() uint8 {
	maxValue := uint8(0)
	for _, v := range td.TotalTokens {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
