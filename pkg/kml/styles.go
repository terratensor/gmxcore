package kml

// StyleConfig представляет конфигурацию для настройки стиля.
type StyleConfig struct {
	LineWidth float64 // Толщина линии
	LineColor string  // Цвет линии в формате HTML (#RRGGBB)
	FillColor string  // Цвет заливки в формате HTML (#RRGGBB)
	Fill      int     // Включить заливку (1 - да, 0 - нет)
}

// UpdateStyle обновляет стиль с указанным ID.
// styles - список стилей.
// id - идентификатор стиля для обновления.
// config - конфигурация стиля.
func UpdateStyle(styles []Style, id string, config StyleConfig) []Style {
	for i, style := range styles {
		if style.ID == id {
			styles[i].LineStyle.Width = config.LineWidth
			styles[i].LineStyle.Color = htmlColorToKML(config.LineColor)
			styles[i].PolyStyle.Fill = config.Fill
			if config.FillColor != "" {
				styles[i].PolyStyle.Color = htmlColorToKML(config.FillColor)
			}
			break
		}
	}
	return styles
}

// UpdateStyleMap обновляет StyleMap с указанным ID.
// styleMaps - список карт стилей.
// id - идентификатор карты стилей.
// normalStyleURL - стиль для обычного состояния.
// highlightStyleURL - стиль для состояния при наведении.
func UpdateStyleMap(styleMaps []StyleMap, id, normalStyleURL, highlightStyleURL string) []StyleMap {
	for i, styleMap := range styleMaps {
		if styleMap.ID == id {
			for j, pair := range styleMap.Pairs {
				if pair.Key == "normal" {
					styleMaps[i].Pairs[j].StyleURL = normalStyleURL
				} else if pair.Key == "highlight" {
					styleMaps[i].Pairs[j].StyleURL = highlightStyleURL
				}
			}
			break
		}
	}
	return styleMaps
}

// DefaultStyles возвращает предустановленные стили.
func DefaultStyles() []Style {
	return []Style{
		{
			ID: "geohash0",
			LineStyle: LineStyle{
				Width: 0.1,
				Color: htmlColorToKML("#00ff00"), // Зелёный цвет
			},
			PolyStyle: PolyStyle{
				Fill: 0,
			},
		},
		{
			ID: "geohash1",
			LineStyle: LineStyle{
				Width: 0.1,
				Color: htmlColorToKML("#00ff00"), // Зелёный цвет
			},
			PolyStyle: PolyStyle{
				Fill: 0,
			},
		},
	}
}

// DefaultStyleMaps возвращает предустановленные карты стилей.
func DefaultStyleMaps() []StyleMap {
	return []StyleMap{
		{
			ID: "geohash",
			Pairs: []Pair{
				{
					Key:      "normal",
					StyleURL: "#geohash0",
				},
				{
					Key:      "highlight",
					StyleURL: "#geohash1",
				},
			},
		},
	}
}

// DefaultDocumentConfig возвращает предустановленную конфигурацию документа.
// name - название документа.
func DefaultDocumentConfig(name string) Document {
	return Document{
		Name:      name,
		Styles:    DefaultStyles(),
		StyleMaps: DefaultStyleMaps(),
	}
}
