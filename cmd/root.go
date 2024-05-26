/*
Copyright © 2024 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/tfkhdyt/exchango/repo"
)

var (
	from string
	to   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "exchango [flags] [value]",
	Short:   "Currency Conversion Tool",
	Long:    `Currency Conversion Tool`,
	Args:    cobra.ExactArgs(1),
	Example: "exchango --from IDR --to USD 69420",
	Version: "0.0.1",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("value cannot be empty")
			os.Exit(1)
		}

		if from == "" || to == "" {
			fmt.Println("base and target currencies should be set")
			fmt.Println("example: exchango --from IDR --to USD 69420")
			os.Exit(1)
		}
		value, errFloat := strconv.ParseFloat(args[0], 64)
		if errFloat != nil {
			fmt.Printf("value should be an float64: %v\n", errFloat)
			os.Exit(1)
		}

		rate, err := repo.GetRate(from, to)
		cobra.CheckErr(err)

		fmt.Printf("%f\n", value*rate)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&from, "from", "f", "", "Base currency code")
	rootCmd.Flags().StringVarP(&to, "to", "t", "", "Target currency code")

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("failed to read user cache dir:", err.Error())
	}

	cachePath := path.Join(cacheDir, "exchango")
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		if err := os.MkdirAll(cachePath, os.ModePerm); err != nil {
			fmt.Println("failed to create user cache dir:", err.Error())
		}
	}
}
