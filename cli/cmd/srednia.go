/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"encoding/csv"
	"strings"
)

var Plik string
var Dane []string

// sredniaCmd represents the srednia command
var sredniaCmd = &cobra.Command{
	Use:   "srednia",
	Short: "sluzy do obliczania roznych rodzajow srednich",
	Long: `sredia wylicza trzy rodzaje srednich: 
- artytmetyczna
- mediana (wartosc srodkowa)
- dominanta (najczesciej wystepujaca)`,
/*
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("srednia called")
	},
 */
}

func init() {
	sredniaCmd.PersistentFlags().StringVar(&Plik, "plik", "", "plik z danymi")
	sredniaCmd.PersistentFlags().StringSliceVar(&Dane, "dane", nil, "--dane 2.3,4.2,4.4")
	rootCmd.AddCommand(sredniaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sredniaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sredniaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
