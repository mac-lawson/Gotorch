package datasets

import (
	"encoding/csv"
	"log"
	"os"
)

func (c *CSVdata) Pop(index uint8) *CSVdata {
	new_data := append(c.Data[:index], c.Data[index+1:]...)
	data, err := NewCSV(new_data)
	return
}

func (c *CSVdata) Add(data [][]string) [][]string {
	return append(c.Data, data...)
}

func NewCSV(data [][]string) (*CSVdata, error) {
	return &CSVdata{
		Data: data,
	}, nil
}

func FromCSVFile(path string) (*CSVdata, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return NewCSV(records)
}
