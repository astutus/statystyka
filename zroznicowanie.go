package statystyka

import (
	"log"
	"sort"
)

// Rozstep podaje miare na jakiej przestrzeni rozproszone sa dane
func Rozstep(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[len(wejscie) - 1] - wejscie[0]
}

func MinimumRozstepu(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[0]
}

func MaksimumRoztsepu(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[len(wejscie) - 1]
}

func KwartylSrodkowy(wejscie []float64) float64 {
	return Mediana(wejscie)
}

func KwartylDolny(wejscie []float64) float64 {
	if len(wejscie) < 3 {
		log.Fatalln("za malo elementow, minimalnie 3")
	}
}

func KwartylGorny(wejscie []float64) float64 {
	if len(wejscie) < 3 {
		log.Fatalln("za malo elementow, minimalnie 3")
	}
}