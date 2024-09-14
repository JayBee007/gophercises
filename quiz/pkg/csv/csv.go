package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
