/*
Copyright © 2024 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/tfkhdyt/exchango/cmd/cache"
)

// CacheCmd represents the cache command
var CacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage cache",
	Long:  `Manage cache`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("cache called")
	// },
}

func init() {
	CacheCmd.AddCommand(cache.CachePathCmd, cache.CacheClearCmd)
	rootCmd.AddCommand(CacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
