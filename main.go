package main

func main() {
	/*katana, err := newKatana(25, 25)
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
	kk, _ := loadKatanaFromFile("qqq-25x25.json")

	kk.init()
	kk.printKatana()
	fmt.Println(kk.checkColumns())
	*/
	//dd, _ := newData(25, 25)
	//saveDataToFile("dd25.json", dd)
	dd, _ := loadDataFromFile("dd25.json")
	dd.printRows()
	dd.printCols()

	ss, _ := newKatana1(dd)
	ss.printKatana()
}
