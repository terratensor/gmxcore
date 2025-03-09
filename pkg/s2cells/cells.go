package s2cells

import (
	"fmt"
	"os"

	"github.com/golang/geo/s2"
)

// CountCellsAtLevel вычисляет количество ячеек на указанном уровне
func CountCellsAtLevel(level int) int {
	if level == 0 {
		return 6 // На уровне 0 всегда 6 ячеек (по одной на каждую грань)
	}
	return 6 * (1 << (2 * uint(level))) // 6 * 4^level
}

// GenerateCellsAtLevel генерирует все ячейки на указанном уровне
func GenerateCellsAtLevel(level int) []s2.CellID {
	var cellIDs []s2.CellID

	// Находим первую ячейку на уровне Level
	cellID := s2.CellIDFromFace(0).ChildBeginAtLevel(level)

	// Итерируемся по всем ячейкам на уровне Level
	for {
		// Переходим к следующей ячейке
		cellID = cellID.Next()
		// Добавляем ячейку в список
		cellIDs = append(cellIDs, cellID)
		// Если мы вернулись к начальной ячейке, завершаем цикл
		if cellID == s2.CellIDFromFace(0).ChildBeginAtLevel(level) {
			break
		}

	}

	return cellIDs
}

// SaveCellsToFile сохраняет ячейки в файл
func SaveCellsToFile(cells []s2.CellID, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %w", err)
	}
	defer file.Close()

	for _, cellID := range cells {
		_, err := file.WriteString(fmt.Sprintf("%s\n", cellID.ToToken()))
		if err != nil {
			return fmt.Errorf("ошибка при записи в файл: %w", err)
		}
	}

	return nil
}
