package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "football-simulator",
	Short: "Football Simulator CLI Application",
	Long:  `A Command Line Interface (CLI) application for simulating football matches.`,
}

// Execute запускает корневую команду, вызывается в main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
