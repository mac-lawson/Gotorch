package tokenizer

type TokenizedData_str struct {
	OrigData      [][]string
	TokenizedData [][]int64
	EncodingMap   map[string]int64
	TotalTokens   map[string]uint8
}
