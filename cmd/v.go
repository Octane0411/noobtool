package cmd

import (
	"github.com/spf13/cobra"
	"noobtool/internel/valorant"
)

var vCmd = &cobra.Command{
	Use:   "v",
	Short: "获取当日瓦洛兰特皮肤商店",
	Run: func(cmd *cobra.Command, args []string) {
		valorant.GetWeapons()
	},
}
