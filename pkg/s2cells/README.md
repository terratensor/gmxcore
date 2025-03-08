# Библиотека для работы с ячейками S2

Библиотека предоставляет удобный интерфейс для работы с ячейками S2 на указанном уровне. Она позволяет вычислять все ячейки заданного уровня и сохранять их в файл.

## Установка

Для использования библиотеки необходимо установить её с помощью команды:

```bash
go get github.com/terratensor/gmxcore/pkg/s2cells
```

## Использование

### Импорт библиотеки

```go
import "github.com/terratensor/gmxcore/pkg/s2cells"
```

### Основные функции

1. **Вычисление количества ячеек на уровне**

```go
func CountCellsAtLevel(level int) int
```

- **Параметры**:
- `level` (int): Уровень ячеек S2.
- **Возвращает**:
- Количество ячеек на указанном уровне.

2. **Генерация всех ячеек на уровне**

```go
func GenerateCellsAtLevel(level int) []s2.CellID
```

- **Параметры**:
- `level` (int): Уровень ячеек S2.
- **Возвращает**:
- Список ячеек (`[]s2.CellID`) на указанном уровне.

3. **Сохранение ячеек в файл**

```go
func SaveCellsToFile(cells []s2.CellID, filename string) error
```

- **Параметры**:
- `cells` ([]s2.CellID): Список ячеек для сохранения.
- `filename` (string): Имя файла для сохранения.
- **Возвращает**:
- Ошибку, если что-то пошло не так.

### Пример использования

```go
package main

import (
    "fmt"
    "github.com/terratensor/gmxcore/pkg/s2cells"
)

func main() {
    // Указываем уровень ячеек
    level := 2

    // Вычисляем количество ячеек на уровне
    numCells := s2cells.CountCellsAtLevel(level)
    fmt.Printf("Количество ячеек на уровне %d: %d\n", level, numCells)

    // Генерируем ячейки на уровне
    cells := s2cells.GenerateCellsAtLevel(level)

    // Сохраняем ячейки в файл
    filename := fmt.Sprintf("level%d_cells.txt", level)
    err := s2cells.SaveCellsToFile(cells, filename)
    if err != nil {
        fmt.Println("Ошибка при сохранении ячеек в файл:", err)
        return
    }

    fmt.Printf("Ячейки уровня %d успешно сохранены в файл %s\n", level, filename)
}
```



### Пример вывода

Для уровня 2 программа создаст файл `level2_cells.txt` с содержимым:

```
29
2b
2d
2f
...
```
---
