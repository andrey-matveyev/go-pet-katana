package main

import (
	"fmt"
)

func main() {
	katana, err := newKatana(3, 4)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	katana.Columns.Values[0] = append(katana.Columns.Values[0], 1)
	katana.Columns.Values[0] = append(katana.Columns.Values[0], 2)
	katana.Columns.Values[0] = append(katana.Columns.Values[0], 3)

	katana.Columns.Values[1] = append(katana.Columns.Values[1], 4)
	katana.Columns.Values[1] = append(katana.Columns.Values[1], 5)
	katana.Columns.Values[1] = append(katana.Columns.Values[1], 6)
	katana.Columns.Values[1] = append(katana.Columns.Values[1], 7)

	saveKatanaToFile("qqq.json", katana)
	kk, _ := loadKatanaFromFile("qqq.json")

	//ss := append(katana.Columns.Values[0], 1)
	//ss = append(katana.Columns.Values[0], 2)
	//ss = append(katana.Columns.Values[0], 4)
	//fmt.Println(ss)

	//fmt.Printf("Lines: %b\\%b", 4, 5)
	fmt.Println(katana.Columns.Values)
	fmt.Println(kk.Columns.Values)

}
