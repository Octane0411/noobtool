package cmd

import (
	"github.com/spf13/cobra"
	"noobtool/internel/valorant"
)

var vCmd = &cobra.Command{
	Use: "v",
	Run: func(cmd *cobra.Command, args []string) {
		valorant.GetWeapons()
	},
}
