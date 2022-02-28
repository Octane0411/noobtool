package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"noobtool/internel/tree"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "树状打印当前dir",
	Long:  "树状打印当前dir",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("路径不能为空")
			return
		}
		tree.PrintTree(args[0])
	},
}
