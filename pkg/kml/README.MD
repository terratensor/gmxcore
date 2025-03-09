# Библиотека для генерации KML-файлов

Эта библиотека предоставляет простой способ генерации KML-файлов на языке Go. Она поддерживает:
- Создание KML-документов с папками, Placemark'ами и стилями.
- Настройку стилей (цвета, толщину линий, заливку).
- Генерацию вложенных структур папок.

## Установка

```bash
go get github.com/terratensor/gmxcore/pkg/kml
```

## Примеры

### 1. Простой пример
Создание базового KML-файла с одним Placemark.

```go
package main

import (
	"fmt"
	"log"

	"github.com/golang/geo/s2"
	"github.com/terratensor/gmxcore/pkg/kml"
)

func main() {
	// Создаём документ с предустановленными стилями
	doc := kml.DefaultDocumentConfig("Простой KML")

	// Создаём корневую папку
	rootFolder := kml.NewFolder("Корневая папка", 1)

	// Добавляем Placemark
	cellID := s2.CellIDFromToken("1234")
	cell := s2.CellFromCellID(cellID)
	coordinates := getCellCoordinates(cell)

	placemark := kml.NewPlacemark(
		cellID.ToToken(),
		"Пример Placemark",
		"#geohash",
		coordinates,
	)
	rootFolder.AddPlacemark(placemark)

	// Добавляем корневую папку в документ
	doc.Folders = append(doc.Folders, rootFolder)

	// Генерируем KML-файл
	if err := kml.GenerateKML("simple.kml", doc); err != nil {
		log.Fatalf("Ошибка при генерации KML: %v", err)
	}
}

func getCellCoordinates(cell s2.Cell) string {
	var coordinates string
	for i := 0; i < 4; i++ {
		vertex := s2.LatLngFromPoint(cell.Vertex(i))
		coordinates += fmt.Sprintf("%f,%f,0 ", vertex.Lng.Degrees(), vertex.Lat.Degrees())
	}
	// Замыкаем полигон
	firstVertex := s2.LatLngFromPoint(cell.Vertex(0))
	coordinates += fmt.Sprintf("%f,%f,0", firstVertex.Lng.Degrees(), firstVertex.Lat.Degrees())
	return coordinates
}
```

### 2. Пример с вложенными папками
Создание KML-файла с вложенными папками.

```go
package main

import (
	"fmt"
	"log"

	"github.com/golang/geo/s2"
	"github.com/terratensor/gmxcore/pkg/kml"
)

func main() {
	// Создаём документ с предустановленными стилями
	doc := kml.DefaultDocumentConfig("KML с вложенными папками")

	// Создаём корневую папку
	rootFolder := kml.NewFolder("Корневая папка", 1)

	// Добавляем Placemark в корневую папку
	cellID := s2.CellIDFromToken("1234")
	cell := s2.CellFromCellID(cellID)
	coordinates := getCellCoordinates(cell)

	placemark := kml.NewPlacemark(
		cellID.ToToken(),
		"Пример Placemark",
		"#geohash",
		coordinates,
	)
	rootFolder.AddPlacemark(placemark)

	// Создаём вложенную папку
	subFolder := kml.NewFolder("Вложенная папка", 1)
	subFolder.AddPlacemark(placemark)

	// Добавляем вложенную папку в корневую
	rootFolder.AddFolder(subFolder)

	// Добавляем корневую папку в документ
	doc.Folders = append(doc.Folders, rootFolder)

	// Генерируем KML-файл
	if err := kml.GenerateKML("nested_folders.kml", doc); err != nil {
		log.Fatalf("Ошибка при генерации KML: %v", err)
	}
}

func getCellCoordinates(cell s2.Cell) string {
	var coordinates string
	for i := 0; i < 4; i++ {
		vertex := s2.LatLngFromPoint(cell.Vertex(i))
		coordinates += fmt.Sprintf("%f,%f,0 ", vertex.Lng.Degrees(), vertex.Lat.Degrees())
	}
	// Замыкаем полигон
	firstVertex := s2.LatLngFromPoint(cell.Vertex(0))
	coordinates += fmt.Sprintf("%f,%f,0", firstVertex.Lng.Degrees(), firstVertex.Lat.Degrees())
	return coordinates
}
```

### 3. Пример с настройкой стилей
Создание KML-файла с кастомными стилями.

```go
package main

import (
    "log"

    "github.com/terratensor/gmxcore/pkg/kml"
    "github.com/golang/geo/s2"
)

func main() {
// Создаём документ с предустановленными стилями
doc := kml.DefaultDocumentConfig("KML с кастомными стилями")

// Настраиваем стили
styleConfig := kml.StyleConfig{
    LineWidth: 0.5, // Толщина линии
    LineColor: "#ff0000", // Красный цвет
    FillColor: "#00ff00", // Зелёный цвет заливки
    Fill: 1, // Включаем заливку
}
doc.Styles = kml.UpdateStyle(doc.Styles, "geohash0", styleConfig)

// Настраиваем StyleMap
doc.StyleMaps = kml.UpdateStyleMap(doc.StyleMaps, "geohash", "#geohash0", "#geohash1")

// Создаём корневую папку
rootFolder := kml.NewFolder("Корневая папка", 1)

// Добавляем Placemark
cellID := s2.CellIDFromToken("1234")
cell := s2.CellFromCellID(cellID)
coordinates := getCellCoordinates(cell)

placemark := kml.NewPlacemark(
    cellID.ToToken(),
    "Пример Placemark",
    "#geohash",
    coordinates,
)
rootFolder.AddPlacemark(placemark)

// Добавляем корневую папку в документ
doc.Folders = append(doc.Folders, rootFolder)

// Генерируем KML-файл
if err := kml.GenerateKML("custom_styles.kml", doc); err != nil {
        log.Fatalf("Ошибка при генерации KML: %v", err)
    }
}

func getCellCoordinates(cell s2.Cell) string {
    var coordinates string
    for i := 0; i < 4; i++ {
        vertex := s2.LatLngFromPoint(cell.Vertex(i))
        coordinates += fmt.Sprintf("%f,%f,0 ", vertex.Lng.Degrees(), vertex.Lat.Degrees())
    }
    // Замыкаем полигон
    firstVertex := s2.LatLngFromPoint(cell.Vertex(0))
    coordinates += fmt.Sprintf("%f,%f,0", firstVertex.Lng.Degrees(), firstVertex.Lat.Degrees())
    return coordinates
}
```

## Справочник API

### Функции

- **`DefaultDocumentConfig(name string) Document`**: Создаёт документ с предустановленными стилями.
- **`GenerateKML(filename string, doc Document) error`**: Генерирует KML-файл.
- **`NewFolder(name string, open int) Folder`**: Создаёт новую папку.
- **`NewPlacemark(name, description, styleURL, coordinates string) Placemark`**: Создаёт новый Placemark.
- **`UpdateStyle(styles []Style, id string, config StyleConfig) []Style`**: Обновляет стиль.
- **`UpdateStyleMap(styleMaps []StyleMap, id, normalStyleURL, highlightStyleURL string) []StyleMap`**: Обновляет карту стилей.

### Типы

- **`Document`**: Представляет KML-документ.
- **`Folder`**: Представляет папку в KML-документе.
- **`Placemark`**: Представляет Placemark в KML-документе.
- **`StyleConfig`**: Конфигурация для настройки стилей.
