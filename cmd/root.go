/*
author:LY
Copyright © 2024 178382892@qq.com
*/
package cmd

import (
	"fmt"
	"github.com/echo-lin/puffer/cmd/account"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "puffer",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if os.Getuid() != 0 {
			fmt.Println("Error: This command requires root privileges")
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(account.Add())
	rootCmd.AddCommand(account.Delete())
	rootCmd.AddCommand(account.List())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/puffer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
