package statystyka

import "sort"

// Rozstep podaje miare na jakiej przestrzeni rozproszone sa dane
func Rozstep(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[len(wejscie) - 1] - wejscie[0]
}

func MinimumRozstepu(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[0]
}

func MaksimumRoztsepu(wejscie[]float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[len(wejscie) - 1]
}
