package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Lines struct {
	Count  int
	Values [][]int
}

type Columns struct {
	Count  int
	Values [][]int
}

type Board struct {
	Values [][]bool
}

type Katana struct {
	Lines   Lines
	Columns Columns
	Board   Board
}

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

func printLines(katana *Katana) {
	fmt.Println()
	fmt.Printf("Lines (%d):\n", katana.Lines.Count)
	fmt.Println("-----------")
	n := 0
	for i := 0; i < katana.Lines.Count; i++ {
		fmt.Println(katana.Lines.Values[i])
		n++
		if n%5 == 0 {
			fmt.Println()
		}
	}
}

func printColumns(katana *Katana) {
	fmt.Println()
	fmt.Printf("Columns (%d):\n", katana.Columns.Count)
	fmt.Println("-----------")
	n := 0
	for i := 0; i < katana.Columns.Count; i++ {
		fmt.Println(katana.Columns.Values[i])
		n++
		if n%5 == 0 {
			fmt.Println()
		}
	}

}

func printBoard(katana *Katana) {
	fmt.Println()
	fmt.Printf("Board (%d x %d):\n", katana.Columns.Count, katana.Lines.Count)
	fmt.Println("----------------")

	fmt.Print("  _")
	n := 0
	for i := 0; i < katana.Columns.Count; i++ {
		n++
		if n%5 == 0 || n == 1 {
			fmt.Printf("%2d", n)
		} else {
			fmt.Print("__")
		}
	}
	fmt.Println()

	n = 0
	for i := 0; i < katana.Lines.Count; i++ {
		n++
		if n%5 == 0 || n == 1 {
			fmt.Printf("%2d|", n)
		} else {
			fmt.Print("  |")
		}

		for j := 0; j < katana.Columns.Count; j++ {

			if katana.Board.Values[i][j] {
				fmt.Print("#|")
			} else {
				fmt.Print("_|")
			}

		}
		fmt.Println()
	}
}

func printKatana(katana *Katana) {
	printColumns(katana)
	printLines(katana)
	printBoard(katana)
}
