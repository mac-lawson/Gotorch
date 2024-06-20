package tokenizer

import (
	"math/rand"
	"time"
)

func Tokenizer_str(data [][]string) *TokenizedData_str {
	m := make(map[string]int64)
	count := make(map[string]uint8)
	rand.Seed(time.Now().UnixNano()) // Seeding the random number generator
	result := make([][]int64, len(data))

	for i, row := range data {
		result[i] = make([]int64, len(row))
		for j, str := range row {
			if _, exists := m[str]; !exists {
				m[str] = rand.Int63() // Generate a random int64
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
