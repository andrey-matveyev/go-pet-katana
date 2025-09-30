package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Data struct {
	RowValues [][]int
	ColValues [][]int
}

func (item *Data) printRows() {
	fmt.Println()
	fmt.Printf("Rows (%d):\n", len(item.RowValues))
	fmt.Println("----------")
	for i, values := range item.RowValues {
		fmt.Println(values)
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
}

func (item *Data) printCols() {
	fmt.Println()
	fmt.Printf("Cols (%d):\n", len(item.ColValues))
	fmt.Println("----------")
	for i, values := range item.ColValues {
		fmt.Println(values)
		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
}

func newData(lines, columns int) (*Data, error) {
	if lines < 1 || columns < 1 {
		return nil, errors.New("'lines' and 'columns' must be greater than 0")
	}

	item := Data{
		RowValues: make([][]int, lines),
		ColValues: make([][]int, columns),
	}
	for i := range item.RowValues {
		item.RowValues[i] = make([]int, 0)
	}
	for i := range item.ColValues {
		item.ColValues[i] = make([]int, 0)
	}

	return &item, nil
}

func loadDataFromFile(fileName string) (*Data, error) {
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file %s: %w", fileName, err)
	}

	data := Data{}

	if err = json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &data, nil
}

func saveDataToFile(fileName string, data *Data) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	if err = os.WriteFile(fileName, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", fileName, err)
	}

	return nil
}
