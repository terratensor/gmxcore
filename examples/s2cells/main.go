package main

import (
	"fmt"

	"github.com/terratensor/gmxcore/pkg/s2cells"
)

func main() {
	level := 2
	cells := s2cells.GenerateCellsAtLevel(level)
	err := s2cells.SaveCellsToFile(cells, fmt.Sprintf("level-%d_cells.txt", level)) // Сохраняем ячейки в файл (например, level2_cells.txt)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Cells saved to level-%d_cells.txt\n", level)
	}
}
