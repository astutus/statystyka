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
	"os"
	"github.com/spf13/cobra"
	"github.com/astutus/statystyka"
)

// dominantaCmd represents the dominanta command
var dominantaCmd = &cobra.Command{
	Use:   "dominanta",
	Short: "dominanta wylicza najczesciej wystepujaca wartosc",
	Long: `dominanta wylicza najczesciej wystepujaca wartosc z podanego zbioru liczb`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dominanta called")
		if Plik != "" && Dane != nil {
			fmt.Println("UWAGA: musisz podac tylko jedna z opcji --plik lub --dane\n")
			sredniaCmd.Help()
			os.Exit(1)
		} else if Plik != "" {
			fmt.Println("mamy podany plik: ", Plik)
		} else if Dane != nil {
			fmt.Println("mamy podane dane: ", Dane)
			fmt.Println(statystyka.Dominanta(konwersja(Dane)))
		} else {
			fmt.Println("UWAGA: musisz podac jedna z opcji: --plik lub --dane\n")
			sredniaCmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	sredniaCmd.AddCommand(dominantaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dominantaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dominantaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

