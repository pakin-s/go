package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("profile.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	// If there are no records, exit
	if len(records) == 0 {
		fmt.Println("No records found in CSV file")
		return
	}

	// Assuming the first row contains the headers
	headers := records[0]

	// Print each record with its corresponding header
	for _, record := range records[1:] {
		// Pair each header with its corresponding value
		for i, value := range record {
			fmt.Printf("%s: %s\n", headers[i], value)
		}
		fmt.Println() // Add a newline between records
	}
}
