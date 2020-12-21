// Pakiet statystyka zajmuje sie obliczeniami statystycznymi
package statystyka

import "sort"

var (
	dane string = "surowe fakty"
	informacje string = "to dane ktorym nadano okreslone zanaczenie"
)

// SredniaArtytmetyczna liczby srednia arytmetyczna liczb zmiennoprzecinkowych
func SredniaArytmetyczna(liczby []float64) float64 {
	var result float64 = 0
	for _, v :=  range liczby {
		result += v
	}
	return result / float64(len(liczby))
}

// Dominanta podaje najczesciej wystepujaca liczbe z listy liczb zmiennoprzecinkowych
func Dominanta(liczby []float64) []float64 {
	var slownik map[float64]int
	slownik = make(map[float64]int)
	// zliczanie ilosci wystepowania kazdej liczby
	for _, v := range liczby {
		slownik[v]++
	}
	var maximum float64 = liczby[0]
	for _, v := range liczby {
		if slownik[v] > slownik[maximum] {
			maximum = v
		}
	}
	var result []float64
	for k, _ := range slownik {
		if k == maximum {
			result = append(result, k)
		}
	}
	return sort.Float64Slice(result)
}

// Mediana podaje wartosc srodkowa dla listy liczb zmiennoprzecinkowych
func Mediana(liczby []float64) float64 {
	sort.Float64s(liczby)
	dlugosc := len(liczby)
	// w przypadku nieparzystych podaje srodkowa wartosc
	if len(liczby) % 2 == 1 {
		return liczby[dlugosc/2]
	} else {
		// w przypadku parzystych oblicza srednia arytetyczna dwoch srodkowych wartosci
		var temp = []float64{liczby[dlugosc/2], liczby[dlugosc/2-1]}
		return SredniaArytmetyczna(temp)
	}
}
