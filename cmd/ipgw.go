package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"noobtool/internel/ipgw"
)

var ipgwCmd = &cobra.Command{
	Use:   "ipgw",
	Short: "ipgw",
	Long:  "ipgw",
	Run: func(cmd *cobra.Command, args []string) {
		h := ipgw.NewIpgwHandler()
		h.Login()
		fmt.Println("登录成功")
	},
}
