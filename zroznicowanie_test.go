package statystyka

import (
	"testing"
)

type proba struct {
	wejscie         []float64
	rozstep         float64
	minimumRozstepu float64
	maximumRoztepu  float64
	kwartylDolny float64
	kwartylSrodkowy float64
	kwartylGorny float64
}

var proby = []proba{
	proba{
		[]float64{8.5, 2.1, 9.0, 1.8, 20.0},
		18.2,
		1.8,
		20.0,
		2.1,
		8.5,
		9.0,
	},
	proba {
		[]float64{10, 30, 3, 3, 13, 6, 11, 7, 7, 10, 10},
		27,
		3,
		30,
		6,
		10,
			11,
	},
	proba {
		[]float64{4, 17, 7, 14, 18, 12, 3, 16, 10, 4, 4, 11},
		15,
		3,
		18,
		4,
		10.5,
		15,
	},
}
func TestRozstep(t *testing.T) {
	for _, p := range proby {
		if Rozstep(p.wejscie) != p.rozstep {
			t.Error("zle obliczony rozstep")
		}
	}
}

func TestMinimumRozstepu(t *testing.T) {
	for _, p := range proby {
		if MinimumRozstepu(p.wejscie) != p.minimumRozstepu {
			t.Error("zle obliczone minimum rozstepu")
		}
	}
}

func TestMaksimumRoztsepu(t *testing.T) {
	for _, p := range proby {
		if MaksimumRoztsepu(p.wejscie) != p.maximumRoztepu {
			t.Error("zle obliczone maksimum rozstepu")
		}
	}
}

func TestKwartylDolny(t *testing.T) {
	for _, p := range proby {
		if KwartylDolny(p.wejscie) != p.kwartylDolny {
			t.Error("zle obliczony kwartyl dolny")
		}
	}
}

func TestKwartylSrodkowy(t *testing.T) {
	for _, p := range proby {
		if KwartylSrodkowy(p.wejscie) != p.kwartylSrodkowy {
			t.Error("zle obliczony kwartyl srodkowy")
		}
	}
}

func TestKwartylGorny(t *testing.T) {
	for _, p := range proby {
		if KwartylGorny(p.wejscie) != p.kwartylGorny {
			t.Error("zle obliczony kwartyl gorny")
		}
	}
}

func TestPercentyl(t *testing.T) {
	testSlice := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	testAnswer := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if ! CompareSlices(Percentyl(10, testSlice), testAnswer) {
		t.Error("zle obliczony percentyl")
	}

}