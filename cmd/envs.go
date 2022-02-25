package cmd

import "github.com/spf13/cobra"

var envsCmd = &cobra.Command{
	Use:   "envs",
	Short: "管理自己的各种环境变量",
	Long:  "管理自己的各种环境变量",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
