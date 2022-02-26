package cmd

import (
	"github.com/spf13/cobra"
	"noobtool/internel/tree"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "树状打印当前dir",
	Long:  "树状打印当前dir",
	Run: func(cmd *cobra.Command, args []string) {
		tree.PrintTree(args[0])
	},
}
