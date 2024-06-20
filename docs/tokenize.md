## Gotorch Tokenizer
The tokenizer takes datasets that do not have a consistent datatype and converts them into
a standard integer or float datatype. 

### String to int64
```go
	data, err := datasets.FromCSVFile("datasets/samples/data.csv")
	if err != nil {
		panic(err)
	}

	tokenized_data, err2 := data.Tokenize()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(tokenized_data)
```
