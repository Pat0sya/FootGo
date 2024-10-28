package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the football simulator",
	Long:  `This command allows you to configure settings for the football simulator.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Configuring the football simulator...WIP")
		// Логика для чтения или изменения настроек конфигурации
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
