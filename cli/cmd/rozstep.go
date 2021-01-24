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
	"fmt"

	"github.com/spf13/cobra"
)

// rozstepCmd represents the rozstep command
var rozstepCmd = &cobra.Command{
	Use:   "rozstep",
	Short: "rozstep to miara tego na jakiej przestrzeni rozproszone sa dane",
	Long: `rozstep to miara tego na jakiej przestrzeni rozproszone sa dane
wyznacza sie go jako roznice miedzy najwieksza, a najmniejsza wartoscia cechy w zbiorze danych`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rozstep called")
	},
}

func init() {
	rozstepCmd.PersistentFlags().StringVar(&Plik, "plik", "", "plik z danymi")
	rozstepCmd.PersistentFlags().StringSliceVar(&Dane, "dane", nil, "--dane 2.3,4.2,4.4")
	rootCmd.AddCommand(rozstepCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rozstepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rozstepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
