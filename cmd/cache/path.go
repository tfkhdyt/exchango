/*
Copyright © 2024 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cache

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tfkhdyt/exchango/usecase"
)

// cachePathCmd represents the cachePath command
var CachePathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print cache path",
	Long:  `Print cache path`,
	Run: func(cmd *cobra.Command, args []string) {
		myCacheDir, err := usecase.GetCachePath()
		cobra.CheckErr(err)

		fmt.Println(myCacheDir)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cachePathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cachePathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
