package geoutils

// Antipode вычисляет антипод для заданных координат (широта, долгота).
// Возвращает широту и долготу антипода.
func Antipode(lat, lon float64) (float64, float64) {
	antipodeLat := -lat
	antipodeLon := lon + 180
	if antipodeLon > 180 {
		antipodeLon -= 360
	}
	return antipodeLat, antipodeLon
}