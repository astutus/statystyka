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
//KwartylDolny to pierwszy kwartyl po podziale zbioru na cztery rowne czesci
func KwartylDolny(wejscie []float64) float64 {
	sort.Float64s(wejscie)
	wielkosc := len(wejscie)
	if wielkosc < 3 {
		log.Fatalln("za malo elementow, minimalnie 3")
	}
	if wielkosc % 4 == 0 {
		return SredniaArytmetyczna(wejscie[wielkosc/4-1:wielkosc/4+1])
	}
	pozycja := int(math.Ceil(float64(len(wejscie) / 4.0)))
	return wejscie[pozycja]

}

// KwartylGorny to trzeci kwartyl po podziale zbioru na cztery rowne czesci
func KwartylGorny(wejscie []float64) float64 {
	wielkosc := len(wejscie)
	if wielkosc < 3 {
		log.Fatalln("za malo elementow, minimalnie 3")
	}
	if 3* wielkosc % 4 == 0 {
		return SredniaArytmetyczna(wejscie[(3* wielkosc / 4) - 1: (3*wielkosc / 4) + 1 ])
	}
	return math.Ceil(float64(wejscie[3*wielkosc / 4]))
}
// Percentyl okresla procent (uporzadkowanych) danych, ktore odcina
func Percentyl(procent int, wejscie []float64 ) []float64 {
	dlugosc := len(wejscie)
	if dlugosc < 100 {
		log.Fatalln("zbyt malo danych")
	}
	sort.Float64s(wejscie)
	if (procent * dlugosc) % 100 == 0 {
		pozycja := procent * dlugosc / 100
		var result = make([]float64, procent)
		wartosc := SredniaArytmetyczna([]float64{wejscie[pozycja-1], wejscie[pozycja]})
		for _, w := range wejscie {
			if w < wartosc {
				result = append(result, w)
			}
		}
		result = append(result, wartosc)
		return result
	}
	pozycja := math.Ceil(float64(procent * dlugosc) / 100.0)
	return wejscie[0:int(pozycja)]

}