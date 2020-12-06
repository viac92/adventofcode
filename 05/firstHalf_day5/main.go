// day 5

package main

import (
	"bufio"
	"fmt"
	"os"
)

const totalRow int = 128

func makeRow() []int {
	var rowSlice []int
	for i := 0; i < totalRow; i++ {
		rowSlice = append(rowSlice, i)
	} 
	return rowSlice
}

func main() {
	b := bufio.NewScanner(os.Stdin)
	rowSlice := makeRow()
	columnsSlice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	highestID := 0

	for b.Scan() {
		t := b.Text()
		rowCode := t[:7]

		for _,r := range rowCode {
			if r == 'B' {
				rowSlice = rowSlice[len(rowSlice) / 2:]
			} else {
				rowSlice = rowSlice[:len(rowSlice) / 2]
			}
		}

		columnsCode := t[7:]

		for _,r := range columnsCode {
			if r == 'R' {
				columnsSlice = columnsSlice[len(columnsSlice) / 2:]
			} else {
				columnsSlice = columnsSlice[:len(columnsSlice) / 2]
			}
		}

		id := rowSlice[0] * 8 + columnsSlice[0]
	

		if id > highestID {
			highestID = id
		}

		fmt.Printf("Row - %v ----> %v\nColumns - %v ----> %v\nID: %v\n\n", rowCode, rowSlice, columnsCode, columnsSlice, id)
		
		rowSlice = makeRow()
		columnsSlice = []int{0, 1, 2, 3, 4, 5, 6, 7}
	}
	
	fmt.Println(highestID)
}
