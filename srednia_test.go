package statystyka

import (
	"fmt"
	"testing"
)

type averages struct {
	wejscie []float64
	sredniaArytmetyczna float64
	dominanta []float64
	mediana float64
}

var testy = []averages{
	{
		[]float64 {2.0, 30.2, 79.8, 30.2, 79.8, 10.0, 20.0, 2.0},
		31.75,
		[]float64 {2.0, 30.2, 79.8},
		29.25,
	},
	{
			[]float64 {0, 0.4, 0.6, 0.8, 1.4, 0.4},
			0.6,
			[]float64 {0.4},
			0.5,
	},
	{
				[]float64 {34.0,  43.0,  81.0,  106.0,  106.0,  116.0},
				81.0,
				[]float64 {106.0},
				93.5,
	},
}

func TestSredniaArytmetyczna(t *testing.T) {
	for _, average := range testy {
		result := SredniaArytmetyczna(average.wejscie)
		fmt.Println("wyliczona srednia arytmetyczna", result)
		if result != average.sredniaArytmetyczna {
			t.Error("zle wylicza srednia arytmetyczna")
		}

	}
}

func TestMediana(t *testing.T) {
	for _, average := range testy {
		result := Mediana(average.wejscie)
		if result != average.mediana {
			t.Error("zle wylicza mediane")
		}
	}

}

func TestDominanta(t *testing.T) {
	for _, average := range testy {
		result := Dominanta(average.wejscie)
		if ! CompareSlices(result, average.dominanta) {
			t.Error("zle wylicza najczesciej wystepujace wartosci")
		}
	}
}

func CompareSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
	}
	return true
}

