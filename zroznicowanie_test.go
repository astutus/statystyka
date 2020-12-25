package statystyka

import "testing"

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
}
func TestRozstep(t *testing.T) {

}

func TestMinimumRozstepu(t *testing.T) {

}

func TestMaksimumRoztsepu(t *testing.T) {

}

func TestKwartylDolny(t *testing.T) {

}

func TestKwartylSrodkowy(t *testing.T) {

}

func TestKwartylGorny(t *testing.T) {

}