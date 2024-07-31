package account

import (
	"github.com/echo-lin/puffer/internal/services"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func List() *cobra.Command {
	var createAccountCmd = &cobra.Command{
		Use:   "list",
		Short: "Show accounts list",
		Long:  `Show accounts list.`,
		Run: func(cmd *cobra.Command, args []string) {
			allAccounts, err := services.AllAccounts()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"归属", "用户名", "密码", "备注"})

			table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
				tablewriter.Colors{tablewriter.BgHiBlueColor, tablewriter.BgMagentaColor},
				tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
				tablewriter.Colors{tablewriter.BgHiMagentaColor, tablewriter.FgWhiteColor})

			table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
				tablewriter.Colors{tablewriter.Bold, tablewriter.FgMagentaColor})
			for _, row := range allAccounts {
				item := []string{row.Domain, row.Username, row.Password, row.Desc}
				table.Append(item)
			}
			table.Render()
		},
	}

	return createAccountCmd
}
