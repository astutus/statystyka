/*
Copyright © 2021 Olimp Bockowski

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
	"fmt"
	"github.com/astutus/statystyka"
	"os"
	"github.com/spf13/cobra"
)

// medianaCmd represents the mediana command
var medianaCmd = &cobra.Command{
	Use:   "mediana",
	Short: "mediana wylicza srodkowa wartosc",
	Long: `mediana wylicza srodkowa wartosc ze zbioru danych`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mediana called")
		if Plik != "" && Dane != nil {
			fmt.Println("UWAGA: musisz podac tylko jedna z opcji --plik lub --dane\n")
			sredniaCmd.Help()
			os.Exit(1)
		} else if Plik != "" {
			fmt.Println("mamy podany plik: ", Plik)
		} else if Dane != nil {
			fmt.Println("mamy podane dane: ", Dane)
			fmt.Println(statystyka.Mediana(konwersja(Dane)))
		} else {
			fmt.Println("UWAGA: musisz podac jedna z opcji: --plik lub --dane\n")
			sredniaCmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	sredniaCmd.AddCommand(medianaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// medianaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// medianaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
