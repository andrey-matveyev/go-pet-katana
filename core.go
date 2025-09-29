package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func newKatana(x, y int) (*Katana, error) {
	if x < 1 || y < 1 {
		return nil, errors.New("'x' or 'y' must been > 1")
	}

	item := Katana{}
	item.Columns.Count = x
	item.Lines.Count = y

	item.Board.Values = make([][]bool, item.Lines.Count)
	for i := 0; i < item.Lines.Count; i++ {
		item.Board.Values[i] = make([]bool, item.Columns.Count)
	}

	item.Columns.Values = make([][]int, item.Columns.Count)
	for i := 0; i < item.Columns.Count; i++ {
		item.Columns.Values[i] = make([]int, 0)
	}

	item.Lines.Values = make([][]int, item.Lines.Count)
	for i := 0; i < item.Lines.Count; i++ {
		item.Lines.Values[i] = make([]int, 0)
	}

	return &item, nil
}

func loadKatanaFromFile(fileName string) (*Katana, error) {
	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file %s: %w", fileName, err)
	}

	katana := Katana{}

	err = json.Unmarshal(jsonData, &katana)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &katana, nil
}

func saveKatanaToFile(fileName string, katana *Katana) error {
	jsonData, err := json.MarshalIndent(katana, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file %s: %w", fileName, err)
	}

	return nil
}
