package kml

import (
	"encoding/xml"
	"os"
)

// KML представляет корневой элемент KML-документа.
type KML struct {
	XMLName  xml.Name `xml:"kml"`
	XMLNS    string   `xml:"xmlns,attr"` // Пространство имён KML
	Document Document `xml:"Document"`   // Документ KML
}

// Document представляет элемент <Document> в KML.
type Document struct {
	Name      string     `xml:"name,omitempty"`     // Название документа
	Styles    []Style    `xml:"Style,omitempty"`    // Стили документа
	StyleMaps []StyleMap `xml:"StyleMap,omitempty"` // Карты стилей
	Folders   []Folder   `xml:"Folder,omitempty"`   // Папки документа
}

// Folder представляет элемент <Folder> в KML.
type Folder struct {
	Name       string      `xml:"name"`                // Название папки
	Open       int         `xml:"open,omitempty"`      // Открыта ли папка (1 - да, 0 - нет)
	Placemarks []Placemark `xml:"Placemark,omitempty"` // Placemark'и в папке
	Folders    []Folder    `xml:"Folder,omitempty"`    // Вложенные папки
}

// Placemark представляет элемент <Placemark> в KML.
type Placemark struct {
	Name        string     `xml:"name"`                  // Название Placemark
	Description string     `xml:"description,omitempty"` // Описание Placemark
	StyleURL    string     `xml:"styleUrl,omitempty"`    // Ссылка на стиль
	LineString  LineString `xml:"LineString,omitempty"`  // Линия (LineString)
	Polygon     *Polygon   `xml:"Polygon,omitempty"`     // Полигон (Polygon)
}

// LineString представляет элемент <LineString> в KML.
type LineString struct {
	Tessellate  int    `xml:"tessellate"`  // Включить тесселяцию (1 - да, 0 - нет)
	Coordinates string `xml:"coordinates"` // Координаты линии
}

// Polygon представляет элемент <Polygon> в KML.
type Polygon struct {
	OuterBoundaryIs OuterBoundaryIs `xml:"outerBoundaryIs"` // Внешняя граница полигона
}

// OuterBoundaryIs представляет элемент <outerBoundaryIs> в KML.
type OuterBoundaryIs struct {
	LinearRing LinearRing `xml:"LinearRing"` // Линейное кольцо
}

// LinearRing представляет элемент <LinearRing> в KML.
type LinearRing struct {
	Coordinates string `xml:"coordinates"` // Координаты линейного кольца
}

// Style представляет элемент <Style> в KML.
type Style struct {
	ID        string    `xml:"id,attr"`   // Идентификатор стиля
	LineStyle LineStyle `xml:"LineStyle"` // Стиль линии
	PolyStyle PolyStyle `xml:"PolyStyle"` // Стиль заливки
}

// StyleMap представляет элемент <StyleMap> в KML.
type StyleMap struct {
	ID    string `xml:"id,attr"` // Идентификатор карты стилей
	Pairs []Pair `xml:"Pair"`    // Пары стилей
}

// Pair представляет элемент <Pair> в KML.
type Pair struct {
	Key      string `xml:"key"`      // Ключ (normal или highlight)
	StyleURL string `xml:"styleUrl"` // Ссылка на стиль
}

// LineStyle представляет элемент <LineStyle> в KML.
type LineStyle struct {
	Width float64 `xml:"width"` // Толщина линии
	Color string  `xml:"color"` // Цвет линии в формате KML (AABBGGRR)
}

// PolyStyle представляет элемент <PolyStyle> в KML.
type PolyStyle struct {
	Fill  int    `xml:"fill"`  // Включить заливку (1 - да, 0 - нет)
	Color string `xml:"color"` // Цвет заливки в формате KML (AABBGGRR)
}

// GenerateKML создаёт KML-файл на основе переданной структуры документа.
// filename - имя файла для сохранения.
// doc - структура документа KML.
func GenerateKML(filename string, doc Document) error {
	kml := KML{
		XMLNS: "http://www.opengis.net/kml/2.2",
		Document: Document{
			Name:      doc.Name,
			Styles:    doc.Styles,
			StyleMaps: doc.StyleMaps,
			Folders:   doc.Folders,
		},
	}

	// Создаём файл
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем XML в файл
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(kml); err != nil {
		return err
	}

	return nil
}

// NewFolder создаёт новую папку с указанным именем и состоянием (открыта/закрыта).
// name - название папки.
// open - состояние папки (1 - открыта, 0 - закрыта).
func NewFolder(name string, open int) Folder {
	return Folder{
		Name: name,
		Open: open,
	}
}

// AddPlacemark добавляет Placemark в папку.
// placemark - Placemark для добавления.
func (f *Folder) AddPlacemark(placemark Placemark) {
	f.Placemarks = append(f.Placemarks, placemark)
}

// AddFolder добавляет вложенную папку в текущую папку.
// folder - вложенная папка.
func (f *Folder) AddFolder(folder Folder) {
	f.Folders = append(f.Folders, folder)
}

// NewPlacemark создаёт новый Placemark с линией (LineString).
// name - название Placemark.
// description - описание Placemark.
// styleURL - ссылка на стиль.
// coordinates - координаты линии.
func NewPlacemark(name, description, styleURL string, coordinates string) Placemark {
	return Placemark{
		Name:        name,
		Description: description,
		StyleURL:    styleURL,
		LineString: LineString{
			Tessellate:  1,
			Coordinates: coordinates,
		},
	}
}

// NewPolygonPlacemark создаёт новый Placemark с полигоном (Polygon).
// name - название Placemark.
// description - описание Placemark.
// styleURL - ссылка на стиль.
// coordinates - координаты полигона.
func NewPolygonPlacemark(name, description, styleURL string, coordinates string) Placemark {
	return Placemark{
		Name:        name,
		Description: description,
		StyleURL:    styleURL,
		Polygon: &Polygon{
			OuterBoundaryIs: OuterBoundaryIs{
				LinearRing: LinearRing{
					Coordinates: coordinates,
				},
			},
		},
	}
}

// htmlColorToKML преобразует цвет из формата HTML (#RRGGBB) в формат KML (AABBGGRR).
// htmlColor - цвет в формате HTML.
func htmlColorToKML(htmlColor string) string {
	if len(htmlColor) != 7 || htmlColor[0] != '#' {
		return "ff000000" // Возвращаем чёрный цвет по умолчанию
	}

	red := htmlColor[1:3]
	green := htmlColor[3:5]
	blue := htmlColor[5:7]

	return "ff" + blue + green + red
}
