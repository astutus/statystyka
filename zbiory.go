package statystyka

type zbiorStringow map[string]bool
type zbiorFloatow map[float64]bool
type zbiorIntow map[int]bool

// dla Stringow
func IntersectionZbiorStringow(a, b zbiorStringow) zbiorStringow {
	var wynik = zbiorStringow{}
	for k, _ := range a {
		if b[k] {
			wynik[k] = true
		}
	}
	return wynik
}
func UnionZbiorStringow(a, b zbiorStringow) zbiorStringow {
	var wynik = zbiorStringow{}
	for k, _ := range a {
		wynik[k] = true
	}
	for k, _ := range b {
		wynik[k] = true
	}
	return wynik
}
func DifferenceZbiorStringow(a, b zbiorStringow) zbiorStringow {
	var wynik = zbiorStringow{}
	for k, _ := range a {
		if ! b[k] {
			wynik[k] = true
		}
	}
	return wynik
}

// dla float'ow
func IntersectionZbiorFloatow(a, b zbiorFloatow) zbiorFloatow {
	var wynik = zbiorFloatow{}
	for k, _ := range a {
		if b[k] {
			wynik[k] = true
		}
	}
	return wynik
}
func UnionZbiorFloatow(a, b zbiorFloatow) zbiorFloatow {
	var wynik = zbiorFloatow{}
	for k, _ := range a {
		wynik[k] = true
	}
	for k, _ := range b {
		wynik[k] = true
	}
	return wynik
}
func DifferenceZbiorFloatow(a, b zbiorFloatow) zbiorFloatow {
	var wynik = zbiorFloatow{}
	for k, _ := range a {
		if ! b[k] {
			wynik[k] = true
		}
	}
	return wynik
}

// dla Int'ow
func IntersectionZbiorIntow(a, b zbiorIntow) zbiorIntow {
	var wynik = zbiorIntow{}
	for k, _ := range a {
		if b[k] {
			wynik[k] = true
		}
	}
	return wynik
}
func UnionZbiorIntow(a, b zbiorIntow) zbiorIntow {
	var wynik = zbiorIntow{}
	for k, _ := range a {
		wynik[k] = true
	}
	for k, _ := range b {
		wynik[k] = true
	}
	return wynik
}
func DifferenceZbiorIntow(a, b zbiorIntow) zbiorIntow {
	var wynik = zbiorIntow{}
	for k, _ := range a {
		if ! b[k] {
			wynik[k] = true
		}
	}
	return wynik
}

