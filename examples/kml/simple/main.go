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
