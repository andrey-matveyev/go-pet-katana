package main

import "fmt"

type katana struct {
	data           *Data
	emptyRowValues [][]int  // rows, cols
	emptyRowCounts []int    // total number of possible free cells in a rows
	board          [][]bool // rows, cols
}

func (item *katana) printBoard() {
	fmt.Println()
	fmt.Printf("Board (%d x %d):\n", len(item.data.ColValues), len(item.data.ColValues))
	fmt.Println("----------------")

	fmt.Print("  ")
	for i := range item.board[0] {
		if (i+1)%5 == 0 || i == 0 {
			fmt.Printf("%2d", i+1)
		} else {
			fmt.Print(" _")
		}
	}
	fmt.Println()

	for i, values := range item.board {
		if (i+1)%5 == 0 || i == 0 {
			fmt.Printf("%2d|", i+1)
		} else {
			fmt.Print("  |")
		}
		for _, value := range values {
			if value {
				fmt.Print("#|")
			} else {
				fmt.Print("_|")
			}
		}
		fmt.Println()
	}
}

func (item *katana) printKatana() {
	item.data.printRows()
	item.data.printCols()
	item.printBoard()
}

func newKatana1(inpData *Data) (*katana, error) {
	item := katana{
		data:           inpData,
		emptyRowValues: make([][]int, len(inpData.RowValues)),
		emptyRowCounts: make([]int, len(inpData.RowValues)),
		board:          make([][]bool, len(inpData.RowValues)),
	}
	for i := range item.emptyRowValues {
		item.emptyRowValues[i] = make([]int, 0)
	}
	for i := range item.board {
		item.board[i] = make([]bool, len(inpData.ColValues))
	}

	return &item, nil
}
