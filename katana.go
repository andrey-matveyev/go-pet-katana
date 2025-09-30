package main

import (
	"fmt"
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
	Values [][]bool // lines, columns
}

type Katana struct {
	Lines   Lines
	Columns Columns
	Board   Board
}

func (item *Katana) printLines() {
	fmt.Println()
	fmt.Printf("Lines (%d):\n", item.Lines.Count)
	fmt.Println("-----------")
	n := 0
	for i := 0; i < item.Lines.Count; i++ {
		fmt.Println(item.Lines.Values[i])
		n++
		if n%5 == 0 {
			fmt.Println()
		}
	}
}

func (item *Katana) printColumns() {
	fmt.Println()
	fmt.Printf("Columns (%d):\n", item.Columns.Count)
	fmt.Println("-----------")
	n := 0
	for i := 0; i < item.Columns.Count; i++ {
		fmt.Println(item.Columns.Values[i])
		n++
		if n%5 == 0 {
			fmt.Println()
		}
	}
}

func (item *Katana) printBoard() {
	fmt.Println()
	fmt.Printf("Board (%d x %d):\n", item.Columns.Count, item.Lines.Count)
	fmt.Println("----------------")

	fmt.Print("  _")
	n := 0
	for i := 0; i < item.Columns.Count; i++ {
		n++
		if n%5 == 0 || n == 1 {
			fmt.Printf("%2d", n)
		} else {
			fmt.Print("__")
		}
	}
	fmt.Println()

	n = 0
	for i := 0; i < item.Lines.Count; i++ {
		n++
		if n%5 == 0 || n == 1 {
			fmt.Printf("%2d|", n)
		} else {
			fmt.Print("  |")
		}

		for j := 0; j < item.Columns.Count; j++ {
			if item.Board.Values[i][j] {
				fmt.Print("#|")
			} else {
				fmt.Print("_|")
			}
		}
		fmt.Println()
	}
}

func (item *Katana) printKatana() {
	item.printColumns()
	item.printLines()
	item.printBoard()
}

func (item *Katana) init() {
	for i := 0; i < item.Columns.Count; i++ {
		for j := 0; j < item.Lines.Count; j++ {
			item.Board.Values[i][j] = false
		}
	}

	n := 0
	for i := 0; i < item.Lines.Count; i++ {
		n = 0
		for _, value := range item.Lines.Values[i] {
			for j := 0; j < value; j++ {
				item.Board.Values[i][n] = true
				n++
			}
			n++
		}
	}
}

func (item *Katana) checkColumns() bool {
	for i := 0; i < item.Columns.Count; i++ {
		n := 0
		number := 0

		for _, value := range item.Board.Values[i] {
			if value {
				number++
			} else {
				if number != 0 && number != item.Columns.Values[i][n] {
					return false
				}
				n++
				number = 0
			}
		}
		if number != 0 && number != item.Columns.Values[i][n] {
			return false
		}
	}
	return true
}
