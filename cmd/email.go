package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"noobtool/internel/email"
)

var toers, ccers, subject, body string

func init() {
	emailCmd.Flags().StringVarP(&toers, "toers", "t", "", "请输入接受者，多个用逗号分割，默认为自己")
	emailCmd.Flags().StringVarP(&ccers, "ccers", "c", "", "请输入抄送者，多个用逗号分割")
	emailCmd.Flags().StringVarP(&subject, "subject", "s", "", "请输入邮件标题")
	emailCmd.Flags().StringVarP(&body, "body", "b", "", "请输入邮件内容")
}

var emailCmd = &cobra.Command{
	Use: "email",
	Short: "管理自己的邮件",
	Long:  "...",
	Run: func(cmd *cobra.Command, args []string) {
		err := email.SendEmail(toers, ccers, subject, body)
		if err != nil {
			log.Fatalf("sendEmail err: %v", err)
		}
		fmt.Println("发送成功！")
	},
}



