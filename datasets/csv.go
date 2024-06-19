package datasets

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVdata struct {
	Data [][]string
}

func NewCSV(data [][]string) (*) {}
  // Create a new CSV object
} 

func FromCSVFile(path string) (*CSVdata, error) {
	file, err := os.Open("./samples/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records, nil
}
