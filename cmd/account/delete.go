package account

import (
	"fmt"
	"github.com/echo-lin/puffer/internal/services"
	"github.com/echo-lin/puffer/models"
	"github.com/spf13/cobra"
	"log"
)

func Delete() *cobra.Command {
	var deleteAccountCmd = &cobra.Command{
		Use:   "del",
		Short: "Delete a account",
		Long:  `Delete a account.`,
		Run: func(cmd *cobra.Command, args []string) {
			domain, _ := cmd.Flags().GetString("domain")
			username, _ := cmd.Flags().GetString("username")

			account := models.Account{Username: username, Domain: domain}
			if err := deleteAccount(account); err != nil {
				log.Fatalf("Error appending user to file: %v", err)
			}

			fmt.Println("success")
		},
	}

	// 添加命令和标志参数
	deleteAccountCmd.Flags().StringP("domain", "d", "", "Which platform is the account for")
	deleteAccountCmd.Flags().StringP("username", "u", "", "Your username")
	deleteAccountCmd.MarkFlagRequired("platform")
	deleteAccountCmd.MarkFlagRequired("username")

	return deleteAccountCmd
}

func deleteAccount(account models.Account) error {
	return services.DeleteAccount(account)
}
