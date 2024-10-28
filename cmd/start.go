package cmd

import (
	"Footballsim/pkg/match"
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new football match simulation",
	Long:  `This command initializes and starts a new football match simulation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting a new football match simulation...")
		match := match.NewMatch("FC Golang", "FC Python", 90)
		match.Start()

		fmt.Println("\nСтатистика команд:")
		match.Team1.DisplayTeamStats()
		match.Team2.DisplayTeamStats()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
