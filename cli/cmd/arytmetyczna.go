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

// arytmetycznaCmd represents the arytmetyczna command
var arytmetycznaCmd = &cobra.Command{
	Use:   "arytmetyczna",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Plik != "" && Dane != nil {
			fmt.Println("UWAGA: musisz podac tylko jedna z opcji --plik lub --dane\n")
			sredniaCmd.Help()
			os.Exit(1)
		} else if Plik != "" {
			fmt.Println("mamy podany plik: ", Plik)
		} else if Dane != nil {
			fmt.Println("mamy podane dane: ", Dane)
			fmt.Println(statystyka.SredniaArytmetyczna(konwersja(Dane)))
		} else {
			fmt.Println("UWAGA: musisz podac jedna z opcji: --plik lub --dane\n")
			sredniaCmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	sredniaCmd.AddCommand(arytmetycznaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// arytmetycznaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// arytmetycznaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
