package statystyka

import (
	"log"
	"math"
	"sort"
)

// Rozstep podaje miare na jakiej przestrzeni rozproszone sa dane
func Rozstep(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[len(wejscie) - 1] - wejscie[0]
}

// MinimumRozstepu zwraca namniejsza wartosc w danych
func MinimumRozstepu(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[0]
}
// MaksimumRozstepu zwraca najwieksza wartosc w danych
func MaksimumRoztsepu(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	return wejscie[len(wejscie) - 1]
}

// KwartylSrodkowy zwraca wartosc dzielaca zbior danych na dwie czesci
func KwartylSrodkowy(wejscie []float64) float64 {
	return Mediana(wejscie)
}

func KwartylDolny(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	wielkosc := len(wejscie)
	if wielkosc < 3 {
		log.Fatalln("za malo elementow, minimalnie 3")
	}
	if wielkosc % 4 == 0 {
		return SredniaArytmetyczna(wejscie[wielkosc/4:wielkosc/4+2])
	}
	return  math.Ceil(float64(len(wejscie) / 4.0))

}

func KwartylGorny(wejscie []float64) float64 {
	wielkosc := len(wejscie)
	if wielkosc < 3 {
		log.Fatalln("za malo elementow, minimalnie 3")
	}
	if 3* wielkosc % 4 == 0 {
		return SredniaArytmetyczna(wejscie[3* wielkosc / 4: 3*wielkosc / 4 + 2])
	}
	return math.Ceil(float64(wejscie[3*wielkosc / 4]))
}