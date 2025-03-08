# Библиотека geoutils

Библиотека предоставляет утилиты для работы с географическими координатами.

## Установка

Для использования библиотеки добавьте её в ваш проект:

```bash
go get github.com/terratensor/gmxcore/pkg/geoutils
```

## Использование

### Функция Antipode

Функция Antipode вычисляет антипод для заданных координат (широта, долгота).

```go
import "github.com/terratensor/gmxcore/pkg/geoutils"

func main() {
    lat, lon := 45.0, 90.0
    antipodeLat, antipodeLon := geoutils.Antipode(lat, lon)
    fmt.Printf("Антипод: (%v, %v)\n", antipodeLat, antipodeLon)
}
```