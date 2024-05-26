/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tfkhdyt/exchango/repo"
	"github.com/tfkhdyt/exchango/usecase"
)

var unknown bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show list of supported currencies",
	Long:  `Show list of supported currencies`,
	Run: func(cmd *cobra.Command, args []string) {
		currencies, err := repo.FindAllCurrencies()
		cobra.CheckErr(err)

		usecase.PrintCurrencies(currencies, unknown)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&unknown, "unknown", "u", false, "Show unknown currencies")
}
