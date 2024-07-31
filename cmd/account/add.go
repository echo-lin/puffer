package account

import (
	"fmt"
	"github.com/echo-lin/puffer/internal/services"
	"github.com/echo-lin/puffer/models"
	"github.com/spf13/cobra"
	"log"
)

func Add() *cobra.Command {
	var createAccountCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a account to system",
		Long:  `Add a account to system.`,
		Run: func(cmd *cobra.Command, args []string) {
			domain, _ := cmd.Flags().GetString("domain")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			description, _ := cmd.Flags().GetString("description")
			myAccount := models.Account{Domain: domain, Username: username, Password: password, Desc: description}
			//将新用户写入到 YAML 文件
			if err := appendUserToFile(myAccount); err != nil {
				log.Fatalf("%v", err)
			}

			fmt.Println("success")
		},
	}

	// 添加命令和标志参数
	createAccountCmd.Flags().StringP("domain", "d", "", "Which domain is the account for")
	createAccountCmd.Flags().StringP("username", "u", "", "Your username")
	createAccountCmd.Flags().StringP("password", "p", "", "Your password")
	createAccountCmd.Flags().StringP("description", "e", "", "description of a account")
	createAccountCmd.MarkFlagRequired("domain")
	createAccountCmd.MarkFlagRequired("username")
	createAccountCmd.MarkFlagRequired("password")

	return createAccountCmd
}

func appendUserToFile(account models.Account) error {
	return services.AddAccounts(account)
}
