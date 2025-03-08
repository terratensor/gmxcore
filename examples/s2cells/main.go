package main

import (
	"fmt"

	"github.com/terratensor/gmxcore/pkg/s2cells"
)

func main() {
	level := 2
	cells := s2cells.GenerateCellsAtLevel(level)
	err := s2cells.SaveCellsToFile(cells, "level2_cells.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cells saved to level2_cells.txt")
	}
}
