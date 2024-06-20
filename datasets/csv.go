package datasets

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/mac-lawson/gotorch/tokenizer"
)

func (c *CSVdata) Pop(index uint8) *CSVdata {
	new_data := append(c.Data[:index], c.Data[index+1:]...)
	data, _, err := NewCSV(new_data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (c *CSVdata) Tokenize() (*tokenizer.TokenizedData_str, error) {
	if reflect.TypeOf(c.Data[0][0]) == reflect.TypeOf("") {
		return tokenizer.Tokenizer_str(c.Data), nil
	} else {
		return nil, errors.New("DATA IS NOT TYPE OF STRING")
	}
}

func (c *CSVdata) Add(data [][]string) [][]string {
	return append(c.Data, data...)
}

func NewCSV(data [][]string) (*CSVdata, *Dataset, error) {
	csvdata := &CSVdata{
		Data: data,
	}
	dataset := &Dataset{
		Type:        0, // CSV
		Description: "Local CSV file data",
		Source:      "Local file",
	}
	return csvdata, dataset, nil
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
	data, _, err := NewCSV(records)
	return data, err
}
