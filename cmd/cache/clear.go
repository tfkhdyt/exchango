/*
Copyright © 2024 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cache

import (
	"github.com/spf13/cobra"

	"github.com/tfkhdyt/exchango/usecase"
)

var clearAllCache bool

var force bool

// cacheClearCmd represents the cacheClear command
var CacheClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear unused cache",
	Long:  `Clear unused cache`,
	Run: func(cmd *cobra.Command, args []string) {
		err := usecase.ClearCache(clearAllCache, force)
		cobra.CheckErr(err)
	},
}

func init() {
	CacheClearCmd.Flags().
		BoolVarP(&clearAllCache, "all", "a", false, "Clear all cache")
	CacheClearCmd.Flags().
		BoolVarP(&force, "force", "f", false, "Force clear cache")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheClearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheClearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
