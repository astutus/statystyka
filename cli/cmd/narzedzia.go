package cmd

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func pobieranieDanych(fileName string) []float64 {
	var dane []float64
	plik, err := os.OpenFile(fileName, os.O_RDONLY, 0400)
	if err != nil {
		log.Fatalln("nie da sie otworzyc pliku")
	}
	reader := csv.NewReader(plik)
	pobrane, err := reader.Read()
	if err != nil {
		log.Fatalln("blad odczytu pliku")
	}
	for _, v := range pobrane {
		liczba, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			log.Fatalln("blad, moze byc tylko liczba zmiennoprzecinkowa")
		}
		dane = append(dane, liczba)
	}
	return dane
}

func konwersja(slice1 []string) []float64 {
	var floatSlice []float64
	for _, v := range slice1 {
		liczba, err := strconv.ParseFloat(v, 64)
		if err != nil {
			os.Exit(2)
		}
		floatSlice = append(floatSlice, liczba)
	}
	return floatSlice
}

